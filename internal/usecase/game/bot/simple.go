package bot

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/Eight-Stones/ecs-tank-engine/v2/common"
	"github.com/Eight-Stones/ecs-tank-engine/v2/components"
	"github.com/Eight-Stones/ecs-tank-engine/v2/pkg/utils"

	commonL "go-micro.dev/v4/logger"

	"go-micro-service-template/entity"
)

type Gamer interface {
	Rotate(id string, direction uint) int
	Move(id string, direction uint) int
	Shoot(id string) int
	Vision(id string) (int, []entity.Cell)
	Radar(id string) (int, []entity.Cell)
	Player(id string) *entity.Tank
	IsDead(id string) bool
}

type SimpleBot struct {
	id     string
	cancel func()
	dir    string
	game   Gamer
	logger commonL.Logger
}

func NewSimpleBot(id string, game Gamer, logger commonL.Logger) *SimpleBot {
	return &SimpleBot{
		id:     id,
		dir:    "right",
		game:   game,
		logger: logger,
	}
}

func (sb *SimpleBot) Run(ctx context.Context) {
	ctx, cancel := context.WithCancel(ctx)
	ticker := time.NewTicker(time.Millisecond * 100)

	sb.cancel = cancel

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				if sb.game.IsDead(sb.id) {
					return
				}
				sb.do()
			}
		}
	}()
}

type Meta struct {
	X         int
	Y         int
	Direction string
}

func defineDirectionString(in uint) string {
	switch in {
	case uint(components.Up):
		return "up"
	case uint(components.Down):
		return "down"
	case uint(components.Left):
		return "left"
	case uint(components.Right):
		return "right"
	}

	return ""
}

func defineDirectionUint(in string) uint {
	switch in {
	case "left":
		return uint(components.Left)
	case "right":
		return uint(components.Right)
	case "up":
		return uint(components.Up)
	case "down":
		return uint(components.Down)
	}
	return 0
}

func defineEnemyDirectionStrike(pos, ePos int, line string) (dir string, diff int) {
	if pos > ePos {
		diff = pos - ePos
		if line == "x" {
			dir = "left"
		} else {
			dir = "down"
		}
	} else {
		diff = ePos - pos
		if line == "x" {
			dir = "right"
		} else {
			dir = "up"
		}
	}
	return
}

func defineEnemyDirection(x, y int, m *Meta) (dir string, diff int) {
	var (
		// разница
		xd = -1
		yd = -1
		// направление к врагу по оси
		xdir = ""
		ydir = ""
	)

	switch {
	case m.X == x: // если одна и та же строка
		ydir, yd = defineEnemyDirectionStrike(m.Y, y, "y")
	case m.Y == y: // если одна и та же колонка
		xdir, xd = defineEnemyDirectionStrike(m.X, x, "x")
	default: // если враг не находится на строке/колонке
		ydir, yd = defineEnemyDirectionStrike(m.Y, y, "y")
		xdir, xd = defineEnemyDirectionStrike(m.X, x, "x")
	}

	// вычисляет ближайщую разницу
	switch {
	case xd != -1 && yd == -1:
		dir = xdir
		diff = xd
	case xd == -1 && yd != -1:
		dir = ydir
		diff = yd
	case xd < yd:
		dir = xdir
		diff = xd
	case xd >= yd:
		dir = ydir
		diff = yd
	}

	return dir, diff
}

func (sb *SimpleBot) meta() *Meta {
	meta := sb.game.Player(sb.id)
	if meta == nil {
		sb.cancel()
		return nil
	}

	return &Meta{
		X:         meta.X,
		Y:         meta.Y,
		Direction: defineDirectionString(meta.Direction),
	}
}

func (sb *SimpleBot) view() []entity.Cell {
	result, view := sb.game.Radar(sb.id)
	if !utils.CheckBitMask(result, common.OkRadar) {
		// запросить обзор, в случае недоступности радара
		result, view = sb.game.Vision(sb.id)
		// если запрос неудачен, выйти
		if !utils.CheckBitMask(result, common.OkVision) {
			return nil
		}
	}
	return view
}

type enemy struct {
	X           int
	Y           int
	ToDirection string
	eqrow       bool
	eqcolumn    bool
	diff        int
	typ         string
}

func (e *enemy) isDanger(in *enemy) bool {
	first := 0
	second := 0

	if e.eqrow || e.eqcolumn {
		first++
	}

	if in.eqrow || in.eqcolumn {
		second++
	}

	return first > second && e.diff < in.diff
}

func (sb *SimpleBot) checkView(m *Meta) []enemy {
	view := sb.view()
	if view == nil {
		return nil
	}

	var ens []enemy
	for _, v := range view {
		// Проверить, что этот объект не он сам
		if v.X == m.X && v.Y == m.Y {
			continue
		}

		dir, diff := defineEnemyDirection(v.X, v.Y, m)

		en := enemy{
			X:           v.X,
			Y:           v.Y,
			ToDirection: dir,
			eqrow:       m.X == v.X, // та же строка
			eqcolumn:    m.Y == v.Y, // та же колонка
			diff:        diff,       // чем меньше, тем ближе враг
			typ:         v.Type.String(),
		}

		ens = append(ens, en)
	}

	return ens
}

var reverse = map[string]string{
	"left":  "right",
	"right": "left",
	"up":    "down",
	"down":  "up",
}

var cross = map[string]string{
	"left":  "down",
	"right": "up",
	"up":    "left",
	"down":  "right",
}

func (sb *SimpleBot) enemyAction(m *Meta, in []enemy) {
	sb.logger.Log(commonL.WarnLevel, "ACTION BY FINDING ENEMY "+sb.id[:5])
	main := in[0]
	result := 0
	for idx := 1; idx < len(in); idx++ {
		if !main.isDanger(&in[idx]) {
			main = in[idx]
		}
	}

	if main.typ == "bullet" && (main.eqcolumn || main.eqrow) {
		sb.logger.Log(commonL.WarnLevel, "BULLET TO CLOSE, TRY TO DODGE "+sb.id[:5])
		result = sb.game.Move(sb.id, defineDirectionUint(cross[main.ToDirection]))
		if !utils.CheckBitMask(result, common.OkStep) {
			sb.logger.Log(commonL.WarnLevel, "BULLET TO CLOSE, DODGE FAILED "+sb.id[:5])
			return
		}
		return
	}

	if (main.eqcolumn || main.eqrow) && main.diff <= 12 {
		sb.logger.Log(commonL.WarnLevel, "ENEMY IN LINE, TRY TO CROSS "+sb.id[:5])
		result = sb.game.Rotate(sb.id, defineDirectionUint(main.ToDirection))
		if !utils.CheckBitMask(result, common.OkRotate) {
			sb.logger.Log(commonL.WarnLevel, "ENEMY IN LINE, FAILED ROTATE "+sb.id[:5])
			return
		}

		sb.logger.Log(commonL.WarnLevel, "ENEMY IN CROSS, TRY TO SHOOT "+sb.id[:5]+" dir:"+main.ToDirection+" x:"+fmt.Sprintf("%d", m.X)+" y:"+fmt.Sprintf("%d", m.Y))
		result = sb.game.Shoot(sb.id)
		if !utils.CheckBitMask(result, common.OkShot) {
			sb.logger.Log(commonL.WarnLevel, "ENEMY IN CROSS, FAILED SHOOT "+sb.id[:5])
			return
		}

	}

	if main.diff < 2 {
		sb.logger.Log(commonL.WarnLevel, "ENEMY TO CLOSE, MOVE BACKWARD "+sb.id[:5]+" dir:"+main.ToDirection+" x:"+fmt.Sprintf("%d", m.X)+" y:"+fmt.Sprintf("%d", m.Y))
		result = sb.game.Move(sb.id, defineDirectionUint(reverse[main.ToDirection]))
		if !utils.CheckBitMask(result, common.OkStep) {
			sb.logger.Log(commonL.WarnLevel, "ENEMY TO CLOSE, FAILED MOVE "+sb.id[:5])
			return
		}
		return
	}

	if main.diff >= 8 {
		sb.logger.Log(commonL.WarnLevel, "ENEMY TO FAR, MOVE FORWARD "+sb.id[:5]+" dir:"+main.ToDirection+" x:"+fmt.Sprintf("%d", m.X)+" y:"+fmt.Sprintf("%d", m.Y))
		result = sb.game.Move(sb.id, defineDirectionUint(main.ToDirection))
		if !utils.CheckBitMask(result, common.OkStep) {
			sb.logger.Log(commonL.WarnLevel, "ENEMY TO FAR, FAILED MOVE "+sb.id[:5])
			return
		}
	}

}

func NextDirection(lastDir string, x, y int) string {
	// Инициализация генератора случайных чисел
	rand.Seed(time.Now().UnixNano())

	// Карта зеркальных направлений
	mirror := map[string]string{
		"left":  "right",
		"right": "left",
		"up":    "down",
		"down":  "up",
	}

	// Все возможные направления
	allDirections := []string{"left", "right", "up", "down"}

	// Фильтрация направлений: исключаем зеркальное к предыдущему
	allowedDirections := make([]string, 0)
	for _, dir := range allDirections {
		if dir != mirror[lastDir] {
			allowedDirections = append(allowedDirections, dir)
		}
	}

	// Проверка доступности направлений с учетом границ сетки
	possibleDirections := make([]string, 0)
	for _, dir := range allowedDirections {
		switch dir {
		case "left":
			if x > 0 {
				possibleDirections = append(possibleDirections, dir)
			}
		case "right":
			if x < 14 {
				possibleDirections = append(possibleDirections, dir)
			}
		case "up":
			if y < 14 {
				possibleDirections = append(possibleDirections, dir)
			}
		case "down":
			if y > 0 {
				possibleDirections = append(possibleDirections, dir)
			}
		}
	}

	// Если нет возможных направлений, возвращаем текущее состояние
	if len(possibleDirections) == 0 {
		return lastDir
	}

	// Случайный выбор из возможных направлений
	newDir := possibleDirections[rand.Intn(len(possibleDirections))]

	// Обновление координат

	return newDir
}

func (sb *SimpleBot) freeAction(m *Meta) {
	sb.logger.Log(commonL.WarnLevel, "FREE ACTION "+sb.id[:5])

	if sb.dir == "" {
		sb.dir = "right"
	}

	sb.dir = NextDirection(sb.dir, m.X, m.Y)

	sb.game.Move(sb.id, defineDirectionUint(sb.dir))

}

func (sb *SimpleBot) do() {
	// запрашиваем мету
	m := sb.meta()
	// запрашиваем обзор и врагов
	enemies := sb.checkView(m)
	// если есть враги, действуем в соотвествтии
	if len(enemies) > 0 {
		sb.enemyAction(m, enemies)
		return
	}
	// если врагов нет, то выбираем маршрут движения
	sb.freeAction(m)
}

func (sb *SimpleBot) Stop() {
	sb.cancel()
}

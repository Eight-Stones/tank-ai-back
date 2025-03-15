package game

import (
	"context"
	"time"

	engine "github.com/Eight-Stones/ecs-tank-engine/v2"
	"github.com/Eight-Stones/ecs-tank-engine/v2/components"
	"github.com/Eight-Stones/ecs-tank-engine/v2/config"
	"github.com/Eight-Stones/ecs-tank-engine/v2/systems"
	"github.com/google/uuid"

	"go-micro-service-template/entity"
)

type GameStatus uint

func (s GameStatus) String() string {
	switch s {
	case CreateGameStatus:
		return "Created"
	case ProcessGameStatus:
		return "Processing"
	case EndGameStatus:
		return "Ended"
	}
	return "Unknown"
}

const (
	CreateGameStatus GameStatus = iota
	ProcessGameStatus
	EndGameStatus
)

type Game struct {
	id     string
	status GameStatus
	engine *engine.Field
}

func NewGame() *Game {
	return &Game{
		id:     uuid.New().String(),
		status: CreateGameStatus,
		engine: engine.New(config.WithTankVision(7), config.WithTankRadar(12), config.WithJobsAutoMoverDuration(time.Millisecond*100)),
	}
}

func (g *Game) ID() string {
	return g.id
}

func (g *Game) Status() GameStatus {
	return g.status
}

func (g *Game) Start(ctx context.Context) {
	g.engine.Start(ctx)
	g.status = ProcessGameStatus
}

func (g *Game) Stop() {
	g.engine.Stop()
	g.status = EndGameStatus
}

func (g *Game) AddTank() (string, error) {
	return g.engine.AddTank()
}

func (g *Game) Rotate(id string, direction uint) int {
	return g.engine.Rotate(id, components.Direction(direction))
}

func (g *Game) Move(id string, direction uint) int {
	return g.engine.Move(id, components.Direction(direction))
}

func (g *Game) Shoot(id string) int {
	return g.engine.Shoot(id)
}

func (g *Game) Vision(id string) (int, []entity.Cell) {
	result, view := g.engine.Vision(id)
	out := make([]entity.Cell, 0, len(view))
	for idx := range view {
		for jdx := range view[idx] {
			if view[idx][jdx].ObjectType == 0 {
				continue
			}
			out = append(out, entity.Cell{
				X:         view[idx][jdx].X,
				Y:         view[idx][jdx].Y,
				Direction: entity.Direction(view[idx][jdx].Direction),
				Type:      entity.ObjectType(view[idx][jdx].ObjectType),
			})
		}
	}
	return result, out
}

func (g *Game) Radar(id string) (int, []entity.Cell) {
	result, view := g.engine.Radar(id)
	out := make([]entity.Cell, 0, len(view))
	for idx := range view {
		for jdx := range view[idx] {
			if view[idx][jdx].ObjectType == 0 {
				continue
			}
			out = append(out, entity.Cell{
				X:    view[idx][jdx].X,
				Y:    view[idx][jdx].Y,
				Type: entity.ObjectType(view[idx][jdx].ObjectType),
			})
		}
	}
	return result, out
}

func (g *Game) Player(id string) *entity.Tank {
	tanks := g.Tanks()
	for _, tank := range tanks {
		if tank.ID == id {
			return &tank
		}
	}
	return nil
}

func (g *Game) IsDead(id string) bool {
	for _, dead := range g.engine.DeadObjects {
		if dead.GetInfo().Id == id {
			return true
		}
	}
	return false
}

func (g *Game) Players() []*entity.Player {
	tanks := g.Tanks()
	out := make([]*entity.Player, 0, len(tanks))
	for _, tank := range tanks {
		out = append(out, &entity.Player{
			ID:   tank.ID,
			Game: g.id,
			Tank: tank,
		})
	}
	return out
}

// Tanks collect all alive tanks in engine.
func (g *Game) Tanks() []entity.Tank {
	var tanks []entity.Tank
	for _, obj := range g.engine.Objects {
		if obj.GetInfo().Type != components.TypeTankId {
			continue
		}
		health := obj.(systems.HealthSystem)
		coordinates := obj.(systems.PositionSystem)
		ammo := obj.(systems.ShootingSystem)
		tanks = append(tanks, entity.Tank{
			Object: entity.Object{
				ID:        obj.GetInfo().Id,
				Direction: uint(coordinates.GetPosition().Direction),
				HP:        health.GetHealth().HitPoints,
				IsAlive:   health.GetHealth().HitPoints > 0,
				Type:      entity.ObjectType(obj.GetInfo().Type),
				Coordinates: entity.Coordinates{
					X: coordinates.GetPosition().X,
					Y: coordinates.GetPosition().Y,
				},
			},
			Ammo: ammo.GetShooting().Ammo,
		})
	}
	return tanks
}

func (g *Game) View() [][]entity.Cell {
	rows := make([][]entity.Cell, 15)
	for idx := range rows {
		rows[idx] = make([]entity.Cell, 15)
		for jdx := range rows[idx] {
			rows[idx][jdx] = entity.Cell{
				X: idx,
				Y: jdx,
			}
		}
	}

	for idx := range g.engine.Objects {
		info := g.engine.Objects[idx].GetInfo()
		position := g.engine.Objects[idx].(systems.PositionSystem).GetPosition()
		typ := uint(components.TypeTankId)
		if info.Type != components.TypeTankId {
			typ = uint(components.TypeBulletId)
		}
		rows[position.X][position.Y] = entity.Cell{
			X:         position.X,
			Y:         position.Y,
			Direction: entity.Direction(position.Direction),
			Type:      entity.ObjectType(typ),
		}
	}
	return rows
}

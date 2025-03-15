package convert

import (
	"go-micro-service-template/entity"
	"go-micro-service-template/internal/controller/rest/dto"
)

func PlayerToDTO(in *entity.Player) *dto.Player {
	return &dto.Player{
		ID:        in.ID,
		Name:      in.Name,
		HitPoints: in.Tank.HP,
		Ammo:      in.Tank.Ammo,
		IsAlive:   in.Tank.IsAlive,
		Direction: entity.Direction(in.Tank.Direction).String(),
		X:         in.Tank.X,
		Y:         in.Tank.Y,
	}
}

func PlayersToDTO(in []*entity.Player) []*dto.Player {
	out := make([]*dto.Player, 0, len(in))
	for i := range in {
		out = append(out, PlayerToDTO(in[i]))
	}
	return out
}

func GameToDTO(in *entity.Game) *dto.Game {
	return &dto.Game{
		ID:      in.ID,
		Status:  in.Status,
		Players: PlayersToDTO(in.Player),
	}
}

func GamesToDTO(in []*entity.Game) []*dto.Game {
	out := make([]*dto.Game, 0, len(in))
	for i := range in {
		out = append(out, GameToDTO(in[i]))
	}
	return out
}

func ColumnToDTO(in []entity.Cell) dto.GameColumn {
	out := make(dto.GameColumn, 0, len(in))
	for idx := range in {
		out = append(out, dto.GameCell{
			X:         in[idx].X,
			Y:         in[idx].Y,
			Direction: in[idx].Direction.String(),
			Type:      in[idx].Type.String(),
		})
	}
	return out
}

func LineToDTO(in [][]entity.Cell) dto.GameField {
	lines := make(dto.GameField, 0, len(in))
	for idx := range in {
		lines = append(lines, ColumnToDTO(in[idx]))
	}
	return lines
}

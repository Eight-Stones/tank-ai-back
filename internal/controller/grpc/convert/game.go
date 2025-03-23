package convert

import (
	"go-micro-service-template/entity"
	gamev1 "go-micro-service-template/internal/controller/proto/game/v1"
)

func EntityToGame(in *entity.Game) *gamev1.Game {
	return &gamev1.Game{
		Id: in.ID,
	}
}

func EntityToGames(in []*entity.Game) []*gamev1.Game {
	out := make([]*gamev1.Game, 0, len(in))
	for _, v := range in {
		out = append(out, EntityToGame(v))
	}
	return out
}

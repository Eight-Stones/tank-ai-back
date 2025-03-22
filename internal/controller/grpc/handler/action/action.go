package action

import (
	"go-micro-service-template/entity"
)

type Noder interface {
	Game(id string) (*entity.Game, error)
}

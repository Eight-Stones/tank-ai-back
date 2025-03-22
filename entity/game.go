package entity

import (
	"time"
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
	ID     string
	Start  time.Time
	Status GameStatus
	Player []*Player
}

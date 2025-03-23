package game

import (
	"go-micro-service-template/entity"
)

func (n *Node) gameByPlayer(id string) *Game {
	player, ok := n.players[id]
	if !ok {
		return nil
	}

	g, ok := n.games[player.Game]
	if !ok {
		return nil
	}

	return g
}

func (n *Node) Move(id string, direction uint32) int {
	g := n.gameByPlayer(id)
	if g == nil {
		return 0
	}

	return g.Move(id, uint(direction))
}

func (n *Node) Rotate(id string, direction uint32) int {
	g := n.gameByPlayer(id)
	if g == nil {
		return 0
	}

	return g.Rotate(id, uint(direction))
}

func (n *Node) Shoot(id string) int {
	g := n.gameByPlayer(id)
	if g == nil {
		return 0
	}

	return g.Shoot(id)
}

func (n *Node) Vision(id string) (int, []entity.Cell) {
	g := n.gameByPlayer(id)
	if g == nil {
		return 0, nil

	}

	return g.Vision(id)
}

func (n *Node) Radar(id string) (int, []entity.Cell) {
	g := n.gameByPlayer(id)
	if g == nil {
		return 0, nil
	}

	return g.Radar(id)
}

func (n *Node) Info(id string) *entity.Tank {
	g := n.gameByPlayer(id)
	if g == nil {
		return nil
	}

	return g.Player(id)
}

package game

import (
	"context"
	"fmt"

	commonL "go-micro.dev/v4/logger"

	"go-micro-service-template/entity"
	"go-micro-service-template/internal/usecase/game/bot"
	er "go-micro-service-template/pkg/error"
)

type Node struct {
	// games list games aggregates by ID.
	games map[string]*Game
	// players list players by their ID
	players map[string]*entity.Player
	// bots list of ai group by game.
	bots   map[string][]entity.Bot
	logger commonL.Logger
}

func NewNode(logger commonL.Logger) *Node {
	return &Node{
		games:   make(map[string]*Game),
		players: make(map[string]*entity.Player),
		bots:    make(map[string][]entity.Bot),
		logger:  logger,
	}
}

// Game return game.
func (n *Node) Game(id string) (*entity.Game, error) {
	g := n.games[id]
	if g == nil {
		return nil, er.InternalType.New("game not found")
	}

	players := g.Players()
	for _, player := range players {
		if v, ok := n.players[player.ID]; ok {
			player.Name = v.Name
			continue
		}
	}

	return &entity.Game{
		ID:     g.id,
		Status: g.status.String(),
		Player: players,
	}, nil
}

func (n *Node) Games() []*entity.Game {
	out := make([]*entity.Game, 0, len(n.games))
	for _, game := range n.games {
		out = append(out, &entity.Game{
			ID:     game.id,
			Status: game.status.String(),
			Player: game.Players(),
		})
	}
	return out
}

// CreateGame creates new game.
func (n *Node) CreateGame() string {
	game := NewGame()
	n.games[game.ID()] = game
	return game.ID()
}

func (n *Node) StartGame(ctx context.Context, id string) {
	g, ok := n.games[id]
	if !ok {
		return
	}
	g.Start(context.Background())

	if bots := n.bots[id]; len(bots) > 0 {
		for _, b := range bots {
			b.Run(ctx)
		}
	}
}

func (n *Node) StopGame(id string) {
	g, ok := n.games[id]
	if !ok {
		return
	}
	bots := n.bots[id]
	for _, b := range bots {
		b.Stop()
	}

	g.Stop()
}

func (n *Node) AddPlayer(id, alias string) (string, error) {
	g, ok := n.games[id]
	if !ok {
		return "", fmt.Errorf("game not found")
	}
	playerID, err := g.AddTank()
	if err != nil {
		return "", fmt.Errorf("failed to add player: %w", err)
	}
	n.players[playerID] = &entity.Player{
		ID:   playerID,
		Game: id,
		Name: alias,
	}
	return playerID, nil
}

func (n *Node) AddBot(id string) (string, error) {
	g, ok := n.games[id]
	if !ok {
		return "", fmt.Errorf("game not found")
	}
	playerID, err := g.AddTank()
	if err != nil {
		return "", fmt.Errorf("failed to add bot: %w", err)
	}

	n.players[playerID] = &entity.Player{
		ID:   playerID,
		Game: id,
		Name: "bot_" + playerID[:5],
	}

	n.bots[id] = append(n.bots[id], bot.NewSimpleBot(playerID, g, n.logger))

	return playerID, nil
}

func (n *Node) View(id string) [][]entity.Cell {
	g, ok := n.games[id]
	if !ok {
		return [][]entity.Cell{}
	}
	return g.View()
}

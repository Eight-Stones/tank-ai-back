package game

import (
	"context"
	"time"

	"go-micro-service-template/entity"
	"go-micro-service-template/internal/controller/rest/convert"
	"go-micro-service-template/internal/controller/rest/dto"
	er "go-micro-service-template/pkg/error"
)

type Noder interface {
	CreateGame() string
	StartGame(ctx context.Context, id string) time.Time
	StopGame(id string)

	StartTime(id string) (bool, time.Time)

	Game(id string) (*entity.Game, error)
	Games() []*entity.Game
	AddBot(id string) (string, error)
	Field(id string) [][]entity.Cell
}

// Node ...
type Node struct {
	noder Noder
}

// New build new rest controller for Node.
func New(node Noder) *Node {
	return &Node{
		noder: node,
	}
}

func (n *Node) GetGames(
	_ context.Context,
	_ *dto.GetNodeGamesRequest,
	resp *dto.GetNodeGamesResponse,
) error {
	games := n.noder.Games()
	if len(games) == 0 {
		return nil // todo
	}
	resp.Payload = convert.GamesToDTO(games)
	return nil
}

func (n *Node) GetGameParamId(
	_ context.Context,
	req *dto.GetNodeGameRequest,
	resp *dto.GetNodeGameResponse,
) error {
	game, err := n.noder.Game(req.ID)
	if err != nil {
		return err
	}
	resp.Payload = convert.GameToDTO(game)
	return nil
}

func (n *Node) GetGameViewParamId(
	_ context.Context,
	req *dto.GetNodeViewRequest,
	resp *dto.GetNodeViewResponse,
) error {
	view := n.noder.Field(req.ID)
	if view == nil {
		return nil // todo
	}
	resp.Payload = convert.LineToDTO(view)
	return nil
}

// PostGame create game.
func (n *Node) PostGame(
	_ context.Context,
	_ *dto.PostCreateGameRequest,
	resp *dto.PostCreateGameResponse,
) error {
	id := n.noder.CreateGame()
	resp.Payload = &dto.Game{ID: id}
	return nil
}

// PutGameParamId start/end game.
func (n *Node) PutGameParamId(
	ctx context.Context,
	req *dto.PutRunGameRequest,
	resp *dto.PutRunGameResponse,
) error {
	if req.Action == "run" {
		resp.Start = n.noder.StartGame(ctx, req.ID)
	}

	if req.Action == "stop" {
		n.noder.StopGame(req.ID)
	}

	return nil
}

// PostBotParamId create bot in selected game.
func (n *Node) PostBotParamId(
	_ context.Context,
	req *dto.PostBotRequest,
	_ *dto.PostBotResponse,
) error {
	_, err := n.noder.AddBot(req.ID)
	if err != nil {
		return er.Wrap(err, "n.noder.AddBot")
	}

	return nil
}

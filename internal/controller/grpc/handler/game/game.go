package game

import (
	"context"

	"go-micro.dev/v4/server"

	"go-micro-service-template/entity"
	"go-micro-service-template/internal/controller/grpc/convert"
	gamev1 "go-micro-service-template/internal/controller/proto/game/v1"
	er "go-micro-service-template/pkg/error"
)

type Gamer interface {
	Game(id string) (*entity.Game, error)
	Games() []*entity.Game
	AddPlayer(id, alias string) (string, error)
}

var _ gamev1.GameServiceHandler = &Server{}

type Server struct {
	noder Gamer
}

func New(n Gamer) *Server {
	return &Server{
		noder: n,
	}
}

// Register register handlers on grpc server.
func (s *Server) Register(server server.Server) error {
	err := gamev1.RegisterGameServiceHandler(server, s)
	if err != nil {
		return er.Wrap(err, "register grpc handler games error")
	}
	return nil
}

func (s *Server) Games(_ context.Context, _ *gamev1.GamesReq, resp *gamev1.GamesResp) error {
	games := s.noder.Games()
	resp.Games = convert.EntityToGames(games)
	return nil
}

func (s *Server) Join(_ context.Context, req *gamev1.JoinReq, resp *gamev1.JoinResp) error {
	id, err := s.noder.AddPlayer(req.GameId, req.Alias)
	if err != nil {
		return er.Wrap(err, "add player error")
	}
	resp.PlayerId = id
	return nil
}

func (s *Server) Ready(_ context.Context, req *gamev1.ReadyReq, resp *gamev1.ReadyResp) error {
	g, err := s.noder.Game(req.GameId)
	if err != nil {
		return er.Wrap(err, "get game error")
	}
	if g.Status == entity.ProcessGameStatus {
		resp.IsStarted = true
		resp.StartTimeUtc = g.Start.Unix()
		return nil
	}

	resp.IsStarted = false
	resp.StartTimeUtc = 0

	return nil
}

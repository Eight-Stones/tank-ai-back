package action

import (
	"context"

	"go-micro.dev/v4/server"

	"go-micro-service-template/entity"
	actionv1 "go-micro-service-template/internal/controller/proto/action/v1"
	er "go-micro-service-template/pkg/error"
)

type Actioner interface {
	Move(playerID string, direction uint32) int
	Rotate(id string, direction uint32) int
	Shoot(id string) int
	Vision(id string) (int, []entity.Cell)
	Radar(id string) (int, []entity.Cell)
	Info(id string) *entity.Tank
}

var _ actionv1.ActionServiceHandler = &Server{}

type Server struct {
	action Actioner
}

func New(n Actioner) *Server {
	return &Server{
		action: n,
	}
}

func (s *Server) Info(_ context.Context, req *actionv1.InfoReq, resp *actionv1.InfoResp) error {
	t := s.action.Info(req.Id)
	if t == nil {
		return er.InternalType.New("info not found")
	}
	resp.Info = &actionv1.Info{
		Id:   t.ID,
		X:    int32(t.Coordinates.X),
		Y:    int32(t.Coordinates.Y),
		Hp:   int32(t.HP),
		Ammo: int32(t.Ammo),
	}
	return nil
}

func (s *Server) Rotate(_ context.Context, req *actionv1.RotateReq, resp *actionv1.RotateResp) error {
	resp.Code = actionv1.ActionType(s.action.Rotate(req.Id, req.Direction))
	return nil
}

func (s *Server) Move(_ context.Context, req *actionv1.MoveReq, resp *actionv1.MoveResp) error {
	resp.Code = actionv1.ActionType(s.action.Move(req.Id, req.Direction))
	return nil
}

func (s *Server) Shoot(ctx context.Context, req *actionv1.ShootReq, resp *actionv1.ShootResp) error {
	resp.Code = actionv1.ActionType(s.action.Shoot(req.Id))
	return nil
}

func (s *Server) Vision(ctx context.Context, req *actionv1.VisionReq, resp *actionv1.VisionResp) error {
	code, view := s.action.Vision(req.Id)
	resp.Code = actionv1.ActionType(code)
	out := make([]*actionv1.Cell, 0, len(view))
	for _, v := range view {
		out = append(out, &actionv1.Cell{
			X:    int64(v.X),
			Y:    int64(v.Y),
			Type: actionv1.ObjectType(v.Type),
		})
	}
	resp.Cells = out
	return nil
}

func (s *Server) Radar(ctx context.Context, req *actionv1.RadarReq, resp *actionv1.RadarResp) error {
	code, view := s.action.Radar(req.Id)
	resp.Code = actionv1.ActionType(code)
	out := make([]*actionv1.Cell, 0, len(view))
	for _, v := range view {
		out = append(out, &actionv1.Cell{
			X:    int64(v.X),
			Y:    int64(v.Y),
			Type: actionv1.ObjectType(v.Type),
		})
	}
	resp.Cells = out
	return nil
}

// Register register handlers on grpc server.
func (s *Server) Register(server server.Server) error {
	err := actionv1.RegisterActionServiceHandler(server, s)
	if err != nil {
		return er.Wrap(err, "register grpc handler actions error")
	}
	return nil
}

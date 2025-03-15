package game_controller

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"go-micro-service-template/internal/app/game_controller/config"
	"go-micro-service-template/internal/controller/grpc"
	"go-micro-service-template/internal/controller/rest"
	restgame "go-micro-service-template/internal/controller/rest/handler/game"
	"go-micro-service-template/internal/controller/rest/handler/probe"
	"go-micro-service-template/internal/usecase/game"
	er "go-micro-service-template/pkg/error"
	"go-micro-service-template/pkg/micro/loggerm"
)

func Run(configPath string) error {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	// --------------------------------------------
	// -------------------config-------------------
	// --------------------------------------------

	cfg, err := config.New(configPath)
	if err != nil {
		return er.Wrap(err, "failed to initialize config")
	}

	// --------------------------------------------
	// -------------------logger-------------------
	// --------------------------------------------

	l, err := loggerm.New(
		loggerm.WithLevelString(cfg.Observability.Logger.Level),
		loggerm.WithInitialFields(cfg.Observability.Logger.Keys),
	)
	if err != nil {
		return er.Wrap(err, "failed to initialize logger")
	}

	// --------------------------------------------
	// -------------------db poll------------------
	// --------------------------------------------

	// --------------------------------------------
	// -------------------gateway------------------
	// --------------------------------------------

	// --------------------------------------------
	// -------------------logic--------------------
	// --------------------------------------------

	node := game.NewNode(l)

	// --------------------------------------------
	// -----------------handler--------------------
	// --------------------------------------------

	probeHandler := probe.New()
	gameHandler := restgame.New(node)

	// --------------------------------------------
	// ----------------rest server-----------------
	// --------------------------------------------

	rc := rest.New(
		rest.WithName("tank-ai-rest"),
		rest.WithHost(cfg.Controller.RestTank.Host),
		rest.WithPort(cfg.Controller.RestTank.Port),
		rest.WithLogger(loggerm.Sugar(l)),
		rest.WithHandler(probeHandler),
		rest.WithHandler(gameHandler),
	)

	// --------------------------------------------
	// ----------------grpc server-----------------
	// --------------------------------------------

	gs, err := grpc.New(
		grpc.WithName("tank-ai-grpc"),
		grpc.WithHost(cfg.Controller.GrpcTank.Host),
		grpc.WithPort(cfg.Controller.GrpcTank.Port),
		grpc.WithLogger(loggerm.Sugar(l)),
	)
	if err != nil {
		return er.Wrap(err, "failed to initialize grpc server")
	}

	// --------------------------------------------
	// -------------start http server--------------
	// --------------------------------------------

	rc.Start()

	// --------------------------------------------
	// -------------start grpc server--------------
	// --------------------------------------------

	gs.Start()

	<-ctx.Done()

	return nil
}

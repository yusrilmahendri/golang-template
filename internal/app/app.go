// Package app configures and runs application.
package app

import (
	"database/sql"
v1 "go-clean-template/internal/controller/http/v1"
	"go-clean-template/config"        // ini alias dari repository kamu
	"go-clean-template/internal/repo" // ini alias dari repository kamuumsi ini path handler kamu
	"go-clean-template/internal/usecase"
	"go-clean-template/pkg/httpserver"
)

// run create objects via constructors used mysql
func Run(cfg *config.Config, db *sql.DB) {
	// Sekarang kamu bisa inject db ke repository, contoh:
	pegawaiRepo := repo.NewPegawaiRepository(db)
	pegawaiUC := usecase.NewPegawaiUsecase(pegawaiRepo)
	server := httpserver.New(
	httpserver.Port(cfg.HTTP.Port),
	httpserver.Prefork(cfg.HTTP.UsePreforkMode),
	func(r *gin.Engine) {
		v1.NewRouter(r, pegawaiUC)
	},
	server.Start()
}

// Run creates objects via constructors.
// func Run(cfg *config.Config) {
// 	l := logger.New(cfg.Log.Level)

// 	// Repository
// 	pg, err := postgres.New(cfg.PG.URL, postgres.MaxPoolSize(cfg.PG.PoolMax))
// 	if err != nil {
// 		l.Fatal(fmt.Errorf("app - Run - postgres.New: %w", err))
// 	}
// 	defer pg.Close()

// 	// Use-Case
// 	translationUseCase := translation.New(
// 		persistent.New(pg),
// 		webapi.New(),
// 	)

// 	// RabbitMQ RPC Server
// 	rmqRouter := amqprpc.NewRouter(translationUseCase, l)

// 	rmqServer, err := server.New(cfg.RMQ.URL, cfg.RMQ.ServerExchange, rmqRouter, l)
// 	if err != nil {
// 		l.Fatal(fmt.Errorf("app - Run - rmqServer - server.New: %w", err))
// 	}

// 	// gRPC Server
// 	grpcServer := grpcserver.New(grpcserver.Port(cfg.GRPC.Port))
// 	grpc.NewRouter(grpcServer.App, translationUseCase, l)

// 	// HTTP Server
// 	httpServer := httpserver.New(httpserver.Port(cfg.HTTP.Port), httpserver.Prefork(cfg.HTTP.UsePreforkMode))
// 	http.NewRouter(httpServer.App, cfg, translationUseCase, l)

// 	// Start servers
// 	rmqServer.Start()
// 	grpcServer.Start()
// 	httpServer.Start()

// 	// Waiting signal
// 	interrupt := make(chan os.Signal, 1)
// 	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

// 	select {
// 	case s := <-interrupt:
// 		l.Info("app - Run - signal: %s", s.String())
// 	case err = <-httpServer.Notify():
// 		l.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
// 	case err = <-grpcServer.Notify():
// 		l.Error(fmt.Errorf("app - Run - grpcServer.Notify: %w", err))
// 	case err = <-rmqServer.Notify():
// 		l.Error(fmt.Errorf("app - Run - rmqServer.Notify: %w", err))
// 	}

// 	// Shutdown
// 	err = httpServer.Shutdown()
// 	if err != nil {
// 		l.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
// 	}

// 	err = grpcServer.Shutdown()
// 	if err != nil {
// 		l.Error(fmt.Errorf("app - Run - grpcServer.Shutdown: %w", err))
// 	}

// 	err = rmqServer.Shutdown()
// 	if err != nil {
// 		l.Error(fmt.Errorf("app - Run - rmqServer.Shutdown: %w", err))
// 	}
// }

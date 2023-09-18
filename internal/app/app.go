package app

import (
	"context"
	"fmt"
	"github.com/go-pg/pg"
	"github.com/rs/zerolog"
	"main/api"
	"main/api/middleware"
	"main/internal/config"
	"main/internal/handler"
	"main/internal/service"
	"main/internal/storage"
	"net/http"
	"time"
)

type appServer struct {
	config config.Config
	srv    *http.Server
	db     *pg.DB
	logger *zerolog.Logger
}

func NewServer(config config.Config, logger *zerolog.Logger) *appServer {
	return &appServer{
		config: config,
		logger: logger,
	}
}

func (server *appServer) Serve(ctx context.Context) error {
	server.logger.Info().Msg("Starting server")

	a := server.config.GetDBString()

	dbPool := pg.Connect(&a)
	defer func(dbPool *pg.DB) {
		err := dbPool.Close()
		if err != nil {

		}
	}(dbPool)
	server.db = dbPool
	orderStorage := storage.NewStorage(dbPool)

	orders, err := orderStorage.MigrateStorage()
	if err != nil {
		return err
	}

	orderService := service.NewOrder(orderStorage, orders)
	orderHandler := handler.NewHandler(orderService)
	routes := api.ConfigureRoutes(orderHandler)
	routes.Use(middleware.LogRequest)

	server.srv = &http.Server{
		Addr:    "0.0.0.0:" + server.config.Port,
		Handler: routes,
	}

	server.logger.Info().Msg("Server started.")

	err = server.srv.ListenAndServe()

	if err != nil && err != http.ErrServerClosed {
		server.logger.Err(err).Msg("Failure while serving")
		return err
	}

	return nil
}

func (server *appServer) Shutdown() error {
	server.logger.Info().Msg("server stopped")

	ctxShutDown, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	server.db.Close()

	defer func() {
		cancel()
	}()

	var err error

	if err = server.srv.Shutdown(ctxShutDown); err != nil {
		server.logger.Err(err)

		err = fmt.Errorf("server shutdown failed %w. ", err)

		return err
	}
	server.logger.Info().Msg("Shutdown!")

	return nil
}

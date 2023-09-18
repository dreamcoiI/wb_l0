package main

import (
	"context"
	"fmt"
	"github.com/rs/zerolog"
	"main/internal/app"
	config2 "main/internal/config"
	"os"
	"os/signal"
)

func main() {
	logger := new(zerolog.Logger)

	envFilePath := "config.env"
	config := config2.LoadConfigFromEnv(envFilePath)

	ctx, cancel := context.WithCancel(context.Background())

	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt)

	server := app.NewServer(config, logger)

	go func() {
		oscall := <-c

		logger.Info().Msg(fmt.Sprintf("system call:%+v", oscall))

		if err := server.Shutdown(); err != nil {
			logger.Err(err)
		}

		cancel()
	}()
	if err := server.Serve(ctx); err != nil {
		logger.Err(err)
	}
}

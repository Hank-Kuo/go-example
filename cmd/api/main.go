package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Hank-Kuo/go-example/config"
	"github.com/Hank-Kuo/go-example/internal/server"
	"github.com/Hank-Kuo/go-example/pkg/database"
	"github.com/Hank-Kuo/go-example/pkg/logger"
	"github.com/Hank-Kuo/go-example/pkg/tracer"
)

func main() {
	log.Println("Starting chat-app server")
	cfg, err := config.GetConf()

	if err != nil {
		panic(fmt.Errorf("load config: %v", err))
	}

	apiLogger := logger.NewApiLogger(cfg)
	apiLogger.InitLogger()

	// init database
	db, err := database.ConnectDB(&cfg.Database)
	if err != nil {
		panic(fmt.Errorf("load database: %v", err))
	}
	defer db.Close()

	traceProvider, err := tracer.NewJaeger(cfg)
	if err != nil {
		apiLogger.Fatal("Cannot create tracer", err)
	} else {
		apiLogger.Info("Jaeger connected")
	}
	defer func() {
		if err := traceProvider.Shutdown(context.Background()); err != nil {
			apiLogger.Error("Cannot shutdown tracer", err)
		}
	}()

	// init server
	srv := server.NewServer(cfg, db, apiLogger)
	srv.Run()
}

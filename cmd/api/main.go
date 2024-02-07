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

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func main() {
	log.Println("Starting go-example server")
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

	// db migrate
	driver, err := postgres.WithInstance(db.DB, &postgres.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		"file://./migrations/postgres",
		cfg.Database.Db, driver)
	if err != nil {
		panic(fmt.Errorf("migrate database: %v", err))
	}
	m.Up()

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

package server

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"

	"github.com/Hank-Kuo/go-example/config"
	"github.com/Hank-Kuo/go-example/pkg/logger"
)

type Server struct {
	engine *gin.Engine
	cfg    *config.Config
	db     *sqlx.DB
	logger logger.Logger
}

func NewServer(cfg *config.Config, db *sqlx.DB, logger logger.Logger) *Server {
	return &Server{
		engine: nil,
		cfg:    cfg,
		db:     db,
		logger: logger,
	}
}

func (s *Server) Run() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	httpServer := s.newHttpServer()

	// graceful shutdown
	<-ctx.Done()
	s.logger.Info("Shutdown Server ...")

	if err := httpServer.Shutdown(ctx); err != nil {
		s.logger.Fatal(err)
	}
}

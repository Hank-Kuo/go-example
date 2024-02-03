package server

import (
	"net/http"
	"time"

	taskDelivery "go-example/internal/api/delivery/task"
	taskRepository "go-example/internal/api/repository/task"
	taskService "go-example/internal/api/service/task"
	http_middleware "go-example/internal/middleware/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) registerHttpHanders(engine *gin.Engine) {

	api := engine.Group("/api")
	taskRepo := taskRepository.NewRepo(s.db)
	taskSrv := taskService.NewService(s.cfg, taskRepo, s.logger)
	taskDelivery.NewHttpHandler(api, taskSrv, s.logger)

}

func (s *Server) newHttpServer() *http.Server {
	if s.cfg.Server.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	engine := gin.Default()
	http_middleware.NewGlobalMiddlewares(engine)

	s.registerHttpHanders(engine)

	httpServer := &http.Server{
		Addr:           ":" + s.cfg.Server.Port,
		Handler:        engine,
		ReadTimeout:    time.Second * time.Duration(s.cfg.Server.ReadTimeout),
		WriteTimeout:   time.Second * time.Duration(s.cfg.Server.WriteTimeout),
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			s.logger.Fatalf("Error http ListenAndServe: %s", err)
		}
	}()

	return httpServer
}

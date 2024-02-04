package server

import (
	"net/http"
	"time"

	taskDelivery "github.com/Hank-Kuo/go-example/internal/api/delivery/task"
	taskRepository "github.com/Hank-Kuo/go-example/internal/api/repository/task"
	taskService "github.com/Hank-Kuo/go-example/internal/api/service/task"
	http_middleware "github.com/Hank-Kuo/go-example/internal/middleware/http"

	_ "github.com/Hank-Kuo/go-example/docs"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
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
	if s.cfg.Server.Debug {
		engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
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

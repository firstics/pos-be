package api

import (
	"github.com/gin-gonic/gin"

	"github.com/pos-be/pkg/api/handler"
	"github.com/pos-be/pkg/api/middleware"
	"github.com/pos-be/pkg/config"
	"github.com/pos-be/pkg/driver"
)

type ServerHTTP struct {
	engine *gin.Engine
}

type Middlewares struct {
	ErrorHandler *middleware.ErrorHandler
}

type Handlers struct {
	ExampleHandler *handler.ExampleHandler
}

func NewServerHTTP(middlewares *Middlewares, handlers Handlers, log driver.Logger, cfg config.Config) *ServerHTTP {
	engine := gin.New()
	log.Info("Server is started")

	engine.Use(gin.Recovery())

	// Use logger from Gin
	if cfg.API.Log.Tracing {
		engine.Use(gin.Logger())
	}

	// Use error handler
	engine.Use(middlewares.ErrorHandler.Handler())

	engine.GET("healthcheck", func(c *gin.Context) {
		c.String(200, "OK")
	})

	engine.GET("/example", handlers.ExampleHandler.GetText)

	return &ServerHTTP{engine: engine}
}

func (sh *ServerHTTP) Start() {
	sh.engine.Run(":8081")
}

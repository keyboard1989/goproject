package application

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/keyboard1989/goproject/handlers"
)

type Service struct {
	app *App
}

func NewService(app *App) *Service {
	return &Service{
		app: app,
	}
}

func (s *Service) Run() {
	gin.SetMode(gin.DebugMode)

	if err := s.Engine().Run(s.app.config.Server.Addr + ":" + s.app.config.Server.Port); err != nil {
		fmt.Println("start error")
	}
}

func (s *Service) Engine() *gin.Engine {
	engine := gin.New()

	apiHandler := handlers.NewApiHandler(s.app.httpClient, s.app.statis)
	cmdHandler := handlers.NewCmdHandler(s.app.statis)
	debugHandler := handlers.NewDebugHandler()
	statusHandler := handlers.NewStatusHandler(s.app.statis)

	engine.GET("/", defaultHandler)
	engine.POST("/api", apiHandler.Api())
	engine.GET("/cmd", cmdHandler.Cmd())
	engine.GET("/debug", debugHandler.Debug())
	engine.GET("/status", statusHandler.Status())

	engine.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "not found url \n")
	})
	return engine
}

func defaultHandler(c *gin.Context) {
	c.String(http.StatusOK, "goproject\n")
}

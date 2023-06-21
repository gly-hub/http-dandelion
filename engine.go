package http_engine

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
)

type Engine struct {
	app interface{}
	Router
}

func New(app EngineType) *Engine {
	var engine = &Engine{}
	switch app {
	case Fiber:
		engine.app = fiber.New(fiber.Config{
			DisableStartupMessage: true,
		})
		// 初始化一个空分组作为初始路由
		engine.route = engine.app.(*fiber.App).Group("")
	case Gin:
		gin.SetMode(gin.ReleaseMode)
		engine.app = gin.New()
		// 初始化一个空分组作为初始路由
		engine.route = engine.app.(*gin.Engine).Group("")
	default:
		panic(errors.New("unknown engine"))
	}

	return engine
}

func (e *Engine) Server(addr, port string) error {
	switch e.app.(type) {
	case *fiber.App:
		return e.app.(*fiber.App).Listen(fmt.Sprintf("%s:%s", addr, port))
	case *gin.Engine:
		return e.app.(*gin.Engine).Run(fmt.Sprintf("%s:%s", addr, port))
	}
	return errors.New("unknown engine")
}

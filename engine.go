package http_engine

import (
	"errors"
	"github.com/gly-hub/http-dandelion/core"
	"github.com/gly-hub/http-dandelion/fiber"
	"github.com/gly-hub/http-dandelion/gin"
)

type EngineType string

const (
	Gin      EngineType = "gin"
	Fiber    EngineType = "fiber"
	FastHttp EngineType = "fasthttp"
)

type Engine struct {
	core.IRouter
}

func New(app EngineType) *Engine {
	var engine = &Engine{}
	switch app {
	case Fiber:
		engine.IRouter = &fiber.Router{}
		engine.Init()
	case Gin:
		engine.IRouter = &gin.Router{}
		engine.Init()
	default:
		panic(errors.New("unknown engine"))
	}

	return engine
}

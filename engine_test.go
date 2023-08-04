package http_engine

import (
	"fmt"
	"github.com/gly-hub/http-dandelion/core"
	"testing"
)

func route(engine *Engine) {
	apiGroup := engine.Group("/api")
	apiGroup2 := apiGroup.Group("/t")
	{
		apiGroup2.Get("/test", func(ctx core.IContext) error {
			return nil
		})
	}

	apiGroup3 := apiGroup.Group("/tt")
	{
		apiGroup3.Get("/test", func(ctx core.IContext) error {
			fmt.Println(1232)
			return nil
		})
	}
}

func TestNew(t *testing.T) {
	engine := New("fiber")
	route(engine)
	engine.Server("", "6523")
}

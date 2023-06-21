package http_engine

import (
	"fmt"
	"testing"
)

func route(engine *Engine) {
	apiGroup := engine.Group("/api")
	apiGroup2 := apiGroup.Group("/t")
	{
		apiGroup2.Get("/test", func(ctx Context) error {
			//fmt.Println(ctx.Get("key"))
			return nil
		})
	}

	apiGroup3 := apiGroup.Group("/tt")
	{
		apiGroup3.Get("/test", func(ctx Context) error {
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

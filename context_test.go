package http_engine

import (
	"fmt"
	"testing"
)

func TestContext_Header(t *testing.T) {
	engine := New("fiber")
	apiGroup := engine.Group("/api")
	apiGroup2 := apiGroup.Group("/t")
	{
		apiGroup2.Get("/test", func(ctx Context) error {
			fmt.Println(ctx.Header().Request().Value("id"))
			ctx.Header().Response().Set("id", "修改了")
			return nil
		})
	}
	engine.Server("", "6523")
}

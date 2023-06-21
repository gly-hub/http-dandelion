package http_engine

import (
	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
)

type Handler func(ctx Context) error

func ginHandler(handler ...Handler) (hand []gin.HandlerFunc) {
	for _, h := range handler {
		hand = append(hand, func(c *gin.Context) {
			_ = h(Context{ctx: c})
		})
	}
	return
}

func fiberHandler(handler ...Handler) (hand []fiber.Handler) {
	for _, h := range handler {
		hand = append(hand, func(c *fiber.Ctx) error {
			return h(Context{ctx: c})
		})
	}
	return
}

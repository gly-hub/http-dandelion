package fiber

import (
	"fmt"
	"github.com/gly-hub/http-dandelion/core"
	"github.com/gofiber/fiber/v2"
	"strings"
)

type Router struct {
	app   *fiber.App
	route fiber.Router
}

func (r *Router) Init() {
	r.app = fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})
	// 初始化一个空分组作为初始路由
	r.route = r.app.Group("")
}

func (r *Router) Group(path string, handlers ...core.Handler) core.IRouter {
	path = r.convertPath(path)
	var router = &Router{}
	router.route = r.route.Group(path, r.fiberHandler(handlers...)...)
	return router
}

func (r *Router) Get(path string, handlers ...core.Handler) core.IRouter {
	path = r.convertPath(path)
	var router = &Router{}
	router.route = r.route.(fiber.Router).Get(path, r.fiberHandler(handlers...)...)
	return router
}

func (r *Router) Post(path string, handlers ...core.Handler) core.IRouter {
	path = r.convertPath(path)
	var router = &Router{}
	router.route = r.route.(fiber.Router).Post(path, r.fiberHandler(handlers...)...)
	return router
}

func (r *Router) Put(path string, handlers ...core.Handler) core.IRouter {
	path = r.convertPath(path)
	var router = &Router{}
	router.route = r.route.(fiber.Router).Put(path, r.fiberHandler(handlers...)...)
	return router
}

func (r *Router) Delete(path string, handlers ...core.Handler) core.IRouter {
	path = r.convertPath(path)
	var router = &Router{}
	router.route = r.route.(fiber.Router).Delete(path, r.fiberHandler(handlers...)...)
	return router
}

func (r *Router) Options(path string, handlers ...core.Handler) core.IRouter {
	path = r.convertPath(path)
	var router = &Router{}
	router.route = r.route.(fiber.Router).Options(path, r.fiberHandler(handlers...)...)
	return router
}

func (r *Router) Patch(path string, handlers ...core.Handler) core.IRouter {
	path = r.convertPath(path)
	var router = &Router{}
	router.route = r.route.(fiber.Router).Patch(path, r.fiberHandler(handlers...)...)
	return router
}

func (r *Router) Head(path string, handlers ...core.Handler) core.IRouter {
	path = r.convertPath(path)
	var router = &Router{}
	router.route = r.route.(fiber.Router).Head(path, r.fiberHandler(handlers...)...)
	return router
}

func (r *Router) Use(handler ...core.Handler) core.IRouter {
	var h []interface{}
	for _, hh := range r.fiberHandler(handler...) {
		h = append(h, hh)
	}
	r.route.(fiber.Router).Use(h...)
	return r
}

func (r *Router) convertPath(path string) string {
	routeList := strings.Split(path, "/")
	for i, v := range routeList {
		if strings.HasPrefix(v, "{{") && strings.HasSuffix(v, "}}") {
			routeList[i] = ":" + strings.TrimSuffix(strings.TrimPrefix(v, "{{"), "}}")
		}
	}
	return strings.Join(routeList, "/")
}

func (r *Router) fiberHandler(handler ...core.Handler) (hand []fiber.Handler) {
	for _, h := range handler {
		hand = append(hand, func(c *fiber.Ctx) error {
			return h(&Context{Ctx: c})
		})
	}
	return
}

func (r *Router) Server(addr, port string) error {
	return r.app.Listen(fmt.Sprintf("%s:%s", addr, port))
}

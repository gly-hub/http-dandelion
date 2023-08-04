package gin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gly-hub/http-dandelion/core"
	"strings"
)

type Router struct {
	app   *gin.Engine
	route any
}

func (r *Router) Init() {
	gin.SetMode(gin.ReleaseMode)
	r.app = gin.New()
	// 初始化一个空分组作为初始路由
	r.route = r.app.Group("")
}

func (r *Router) Group(path string, handlers ...core.Handler) core.IRouter {
	path = r.convertPath(path)
	var router = &Router{}
	router.route = r.route.(*gin.RouterGroup).Group(path, r.ginHandler(handlers...)...)
	return router
}

func (r *Router) Get(path string, handlers ...core.Handler) core.IRouter {
	path = r.convertPath(path)
	var router = &Router{}
	router.route = r.route.(*gin.RouterGroup).GET(path, r.ginHandler(handlers...)...)
	return router
}

func (r *Router) Post(path string, handlers ...core.Handler) core.IRouter {
	path = r.convertPath(path)
	var router = &Router{}
	router.route = r.route.(*gin.RouterGroup).POST(path, r.ginHandler(handlers...)...)
	return router
}

func (r *Router) Put(path string, handlers ...core.Handler) core.IRouter {
	path = r.convertPath(path)
	var router = &Router{}
	router.route = r.route.(*gin.RouterGroup).PUT(path, r.ginHandler(handlers...)...)
	return router
}

func (r *Router) Delete(path string, handlers ...core.Handler) core.IRouter {
	path = r.convertPath(path)
	var router = &Router{}
	router.route = r.route.(*gin.RouterGroup).DELETE(path, r.ginHandler(handlers...)...)
	return router
}

func (r *Router) Options(path string, handlers ...core.Handler) core.IRouter {
	path = r.convertPath(path)
	var router = &Router{}
	router.route = r.route.(*gin.RouterGroup).OPTIONS(path, r.ginHandler(handlers...)...)
	return router
}

func (r *Router) Patch(path string, handlers ...core.Handler) core.IRouter {
	path = r.convertPath(path)
	var router = &Router{}
	router.route = r.route.(*gin.RouterGroup).PATCH(path, r.ginHandler(handlers...)...)
	return router
}

func (r *Router) Head(path string, handlers ...core.Handler) core.IRouter {
	path = r.convertPath(path)
	var router = &Router{}
	router.route = r.route.(*gin.RouterGroup).HEAD(path, r.ginHandler(handlers...)...)
	return router
}

func (r *Router) Use(handler ...core.Handler) core.IRouter {
	r.route.(*gin.RouterGroup).Use(r.ginHandler(handler...)...)
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

func (r *Router) ginHandler(handler ...core.Handler) (hand []gin.HandlerFunc) {
	for _, h := range handler {
		hand = append(hand, func(c *gin.Context) {
			_ = h(&Context{Ctx: c})
		})
	}
	return
}

func (r *Router) Server(addr, port string) error {
	return r.app.Run(fmt.Sprintf("%s:%s", addr, port))
}

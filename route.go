package http_engine

import (
	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
	"strings"
)

type Router struct {
	route interface{}
}

func (r *Router) Group(path string, handlers ...Handler) *Router {
	path = r.convertPath(path)
	var router = &Router{}
	switch r.route.(type) {
	case fiber.Router:
		router.route = r.route.(fiber.Router).Group(path, fiberHandler(handlers...)...)
	case *gin.RouterGroup:
		router.route = r.route.(*gin.RouterGroup).Group(path, ginHandler(handlers...)...)
	}
	return router
}

func (r *Router) Get(path string, handlers ...Handler) *Router {
	path = r.convertPath(path)
	var router = &Router{}
	switch r.route.(type) {
	case fiber.Router:
		router.route = r.route.(fiber.Router).Get(path, fiberHandler(handlers...)...)
	case *gin.RouterGroup:
		router.route = r.route.(*gin.RouterGroup).GET(path, ginHandler(handlers...)...)
	}
	return router
}

func (r *Router) Post(path string, handlers ...Handler) *Router {
	path = r.convertPath(path)
	var router = &Router{}
	switch r.route.(type) {
	case fiber.Router:
		router.route = r.route.(fiber.Router).Post(path, fiberHandler(handlers...)...)
	case *gin.RouterGroup:
		router.route = r.route.(*gin.RouterGroup).POST(path, ginHandler(handlers...)...)
	}
	return router
}

func (r *Router) Put(path string, handlers ...Handler) *Router {
	path = r.convertPath(path)
	var router = &Router{}
	switch r.route.(type) {
	case fiber.Router:
		router.route = r.route.(fiber.Router).Put(path, fiberHandler(handlers...)...)
	case *gin.RouterGroup:
		router.route = r.route.(*gin.RouterGroup).PUT(path, ginHandler(handlers...)...)
	}
	return router
}

func (r *Router) Delete(path string, handlers ...Handler) *Router {
	path = r.convertPath(path)
	var router = &Router{}
	switch r.route.(type) {
	case fiber.Router:
		router.route = r.route.(fiber.Router).Delete(path, fiberHandler(handlers...)...)
	case *gin.RouterGroup:
		router.route = r.route.(*gin.RouterGroup).DELETE(path, ginHandler(handlers...)...)
	}
	return router
}

func (r *Router) Options(path string, handlers ...Handler) *Router {
	path = r.convertPath(path)
	var router = &Router{}
	switch r.route.(type) {
	case fiber.Router:
		router.route = r.route.(fiber.Router).Options(path, fiberHandler(handlers...)...)
	case *gin.RouterGroup:
		router.route = r.route.(*gin.RouterGroup).OPTIONS(path, ginHandler(handlers...)...)
	}
	return router
}

func (r *Router) Patch(path string, handlers ...Handler) *Router {
	path = r.convertPath(path)
	var router = &Router{}
	switch r.route.(type) {
	case fiber.Router:
		router.route = r.route.(fiber.Router).Patch(path, fiberHandler(handlers...)...)
	case *gin.RouterGroup:
		router.route = r.route.(*gin.RouterGroup).PATCH(path, ginHandler(handlers...)...)
	}
	return router
}

func (r *Router) Head(path string, handlers ...Handler) *Router {
	path = r.convertPath(path)
	var router = &Router{}
	switch r.route.(type) {
	case fiber.Router:
		router.route = r.route.(fiber.Router).Head(path, fiberHandler(handlers...)...)
	case *gin.RouterGroup:
		router.route = r.route.(*gin.RouterGroup).HEAD(path, ginHandler(handlers...)...)
	}
	return router
}

func (r *Router) Use(handler ...Handler) *Router {
	switch r.route.(type) {
	case fiber.Router:
		var h []interface{}
		for _, hh := range fiberHandler(handler...) {
			h = append(h, hh)
		}
		r.route.(fiber.Router).Use(h...)
	case *gin.RouterGroup:
		r.route.(*gin.RouterGroup).Use(ginHandler(handler...)...)
	}
	return r
}

func (r *Router) convertPath(path string) string {
	routeList := strings.Split(path, "/")
	switch r.route.(type) {
	case fiber.Router:
		for i, v := range routeList {
			if strings.HasPrefix(v, "{{") && strings.HasSuffix(v, "}}") {
				routeList[i] = ":" + strings.TrimSuffix(strings.TrimPrefix(v, "{{"), "}}")
			}
		}
		return strings.Join(routeList, "/")
	case *gin.RouterGroup:
		for i, v := range routeList {
			if strings.HasPrefix(v, "{{") && strings.HasSuffix(v, "}}") {
				routeList[i] = ":" + strings.TrimSuffix(strings.TrimPrefix(v, "{{"), "}}")
			}
		}
		return strings.Join(routeList, "/")
	}
	return path
}

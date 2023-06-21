package http_engine

import (
	"github.com/gin-gonic/gin"
	"github.com/gly-hub/http-dandelion/header"
	"github.com/gly-hub/http-dandelion/params"
	"github.com/gly-hub/http-dandelion/query"
	"github.com/gofiber/fiber/v2"
	"io/ioutil"
	"strings"
)

type Context struct {
	ctx interface{}
}

// Header 提供头部相关操作方法
func (c *Context) Header() header.IHeader {
	switch c.ctx.(type) {
	case *fiber.Ctx:
		return &header.FiberHeader{Ctx: c.ctx.(*fiber.Ctx)}
	case *gin.Context:
		return &header.GinHeader{Ctx: c.ctx.(*gin.Context)}
	}
	return nil
}

// RemoteIP 解析来自请求的IP。RemoteAddr，规范化并返回IP(不带端口)。
func (c *Context) RemoteIP() string {
	var remoteAddr string
	switch c.ctx.(type) {
	case *fiber.Ctx:
		remoteAddr = c.ctx.(*fiber.Ctx).IP()
	case *gin.Context:
		remoteAddr = c.ctx.(*gin.Context).ClientIP()
	}
	return remoteAddr
}

// ContentType 返回请求的Content-Type报头
func (c *Context) ContentType() string {
	switch c.ctx.(type) {
	case *fiber.Ctx:
		return filterFlags(c.ctx.(*fiber.Ctx).Get("Content-Type"))
	case *gin.Context:
		return filterFlags(c.ctx.(*gin.Context).GetHeader("Content-Type"))
	}
	return ""
}

// IsWebsocket 如果请求头表明客户端正在发起websocket握手，则返回true
func (c *Context) IsWebsocket() bool {
	var connection, upgrade string
	switch c.ctx.(type) {
	case *fiber.Ctx:
		connection = c.ctx.(*fiber.Ctx).Get("Connection")
		upgrade = c.ctx.(*fiber.Ctx).Get("Upgrade")
	case *gin.Context:
		connection = c.ctx.(*gin.Context).GetHeader("Connection")
		upgrade = c.ctx.(*gin.Context).GetHeader("Upgrade")
	}
	if strings.Contains(strings.ToLower(connection), "upgrade") &&
		strings.EqualFold(upgrade, "websocket") {
		return true
	}
	return false
}

// URLParam 获取路由参数
func (c *Context) URLParam() params.IParams {
	switch c.ctx.(type) {
	case *fiber.Ctx:
		return &params.FiberParams{Ctx: c.ctx.(*fiber.Ctx)}
	case *gin.Context:
		return &params.GinParams{Ctx: c.ctx.(*gin.Context)}
	}
	return nil
}

// URLQuery 获取字符串参数
func (c *Context) URLQuery() query.IQuery {
	switch c.ctx.(type) {
	case *fiber.Ctx:
		return &query.FiberQuery{Ctx: c.ctx.(*fiber.Ctx)}
	case *gin.Context:
		return &query.GinQuery{Ctx: c.ctx.(*gin.Context)}
	}
	return nil
}

// Body 获取body数据
func (c *Context) Body() []byte {
	switch c.ctx.(type) {
	case *fiber.Ctx:
		return c.ctx.(*fiber.Ctx).Body()
	case *gin.Context:
		body, _ := ioutil.ReadAll(c.ctx.(*gin.Context).Request.Body)
		return body
	}
	return nil
}

// ReadJSON 充body获取json
func (c *Context) ReadJSON(outPtr interface{}) error {
	switch c.ctx.(type) {
	case *fiber.Ctx:
		return c.ctx.(*fiber.Ctx).BodyParser(outPtr)
	case *gin.Context:
		return c.ctx.(*gin.Context).ShouldBindJSON(outPtr)
	}
	return nil
}

// Json 相应
func (c *Context) Json(code int, data interface{}) error {
	switch c.ctx.(type) {
	case *fiber.Ctx:
		return c.ctx.(*fiber.Ctx).Status(code).JSON(data)
	case *gin.Context:
		c.ctx.(*gin.Context).JSON(code, data)
		return nil
	}
	return nil
}

// Next 执行当前路由中的下一个方法
func (c *Context) Next() error {
	switch c.ctx.(type) {
	case *fiber.Ctx:
		return c.ctx.(*fiber.Ctx).Next()
	case *gin.Context:
		c.ctx.(*gin.Context).Next()
		return nil
	}
	return nil
}

// Status 设置响应状态码
func (c *Context) Status(status int) *Context {
	switch c.ctx.(type) {
	case *fiber.Ctx:
		c.ctx = c.ctx.(*fiber.Ctx).Status(status)
		return c
	case *gin.Context:
		c.ctx.(*gin.Context).Status(status)
		return c
	}
	return c
}

// Abort 结束执行当前路由中的方法
func (c *Context) Abort() error {
	switch c.ctx.(type) {
	case *fiber.Ctx:
		return fiber.NewError(fiber.StatusOK, "")
	case *gin.Context:
		c.ctx.(*gin.Context).Abort()
		return nil
	}
	return nil
}

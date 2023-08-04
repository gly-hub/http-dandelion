package fiber

import (
	"github.com/gly-hub/http-dandelion/core"
	"github.com/gofiber/fiber/v2"
	"strings"
)

type Context struct {
	Ctx *fiber.Ctx
}

// Header 提供头部相关操作方法
func (c *Context) Header() core.IHeader {
	return &Header{Ctx: c.Ctx}
}

// RemoteIP 解析来自请求的IP。RemoteAddr，规范化并返回IP(不带端口)。
func (c *Context) RemoteIP() string {
	return c.Ctx.IP()
}

// ContentType 返回请求的Content-Type报头
func (c *Context) ContentType() string {
	return c.filterFlags(c.Ctx.Get("Content-Type"))
}

func (c *Context) filterFlags(content string) string {
	for i, char := range content {
		if char == ' ' || char == ';' {
			return content[:i]
		}
	}
	return content
}

// IsWebsocket 如果请求头表明客户端正在发起websocket握手，则返回true
func (c *Context) IsWebsocket() bool {
	var connection, upgrade string
	connection = c.Ctx.Get("Connection")
	upgrade = c.Ctx.Get("Upgrade")
	if strings.Contains(strings.ToLower(connection), "upgrade") &&
		strings.EqualFold(upgrade, "websocket") {
		return true
	}
	return false
}

// URLParam 获取路由参数
func (c *Context) URLParam() core.IParams {
	return &Params{Ctx: c.Ctx}
}

// URLQuery 获取字符串参数
func (c *Context) URLQuery() core.IQuery {
	return &Query{Ctx: c.Ctx}
}

// Body 获取body数据
func (c *Context) Body() []byte {
	return c.Ctx.Body()
}

// ReadJSON 从body获取json
func (c *Context) ReadJSON(outPtr interface{}) error {
	return c.Ctx.BodyParser(outPtr)
}

// Json 响应
func (c *Context) Json(code int, data interface{}) error {
	return c.Ctx.Status(code).JSON(data)
}

// Next 执行当前路由中的下一个方法
func (c *Context) Next() error {
	return c.Ctx.Next()
}

// Status 设置响应状态码
func (c *Context) Status(status int) core.IContext {
	c.Ctx = c.Ctx.Status(status)
	return c
}

// Abort 结束执行当前路由中的方法
func (c *Context) Abort() error {
	return fiber.NewError(fiber.StatusOK, "")
}

package gin

import (
	"github.com/gin-gonic/gin"
	"github.com/gly-hub/http-dandelion/core"
	"io/ioutil"
	"strings"
)

type Context struct {
	Ctx *gin.Context
}

// Header 提供头部相关操作方法
func (c *Context) Header() core.IHeader {
	return &Header{Ctx: c.Ctx}
}

// RemoteIP 解析来自请求的IP。RemoteAddr，规范化并返回IP(不带端口)。
func (c *Context) RemoteIP() string {
	return c.Ctx.ClientIP()
}

// ContentType 返回请求的Content-Type报头
func (c *Context) ContentType() string {
	return c.filterFlags(c.Ctx.GetHeader("Content-Type"))
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
	connection = c.Ctx.GetHeader("Connection")
	upgrade = c.Ctx.GetHeader("Upgrade")
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
	body, _ := ioutil.ReadAll(c.Ctx.Request.Body)
	return body
}

// ReadJSON 充body获取json
func (c *Context) ReadJSON(outPtr interface{}) error {
	return c.Ctx.ShouldBindJSON(outPtr)
}

// Json 相应
func (c *Context) Json(code int, data interface{}) error {
	c.Ctx.JSON(code, data)
	return nil
}

// Next 执行当前路由中的下一个方法
func (c *Context) Next() error {
	c.Ctx.Next()
	return nil
}

// Status 设置响应状态码
func (c *Context) Status(status int) core.IContext {
	c.Ctx.Status(status)
	return c
}

// Abort 结束执行当前路由中的方法
func (c *Context) Abort() error {
	c.Ctx.Abort()
	return nil
}

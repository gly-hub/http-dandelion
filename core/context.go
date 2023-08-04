package core

type IContext interface {
	// Header 提供头部相关操作方法
	Header() IHeader
	// RemoteIP 解析来自请求的IP。RemoteAddr，规范化并返回IP(不带端口)。
	RemoteIP() string
	// ContentType 返回请求的Content-Type报头
	ContentType() string
	// IsWebsocket 如果请求头表明客户端正在发起websocket握手，则返回true
	IsWebsocket() bool
	// URLParam 获取路由参数
	URLParam() IParams
	// URLQuery 获取字符串参数
	URLQuery() IQuery
	// Body 获取body数据
	Body() []byte
	// ReadJSON 从body获取json
	ReadJSON(outPtr interface{}) error
	// Json 响应
	Json(code int, data interface{}) error
	// Next 执行当前路由中的下一个方法
	Next() error
	// Status 设置响应状态码
	Status(statusCode int) IContext
	// Abort 结束执行当前路由中的方法
	Abort() error
}

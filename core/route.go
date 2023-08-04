package core

type IRouter interface {
	Init()
	Group(path string, handlers ...Handler) IRouter
	Get(path string, handlers ...Handler) IRouter
	Post(path string, handlers ...Handler) IRouter
	Put(path string, handlers ...Handler) IRouter
	Delete(path string, handlers ...Handler) IRouter
	Options(path string, handlers ...Handler) IRouter
	Patch(path string, handlers ...Handler) IRouter
	Head(path string, handlers ...Handler) IRouter
	Use(handler ...Handler) IRouter
	Server(addr, port string) error
}

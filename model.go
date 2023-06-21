package http_engine

type EngineType string

const (
	Gin      EngineType = "gin"
	Fiber    EngineType = "fiber"
	FastHttp EngineType = "fasthttp"
)

package core

type IHeader interface {
	Request() IRequestHeader
	Response() IResponseHeader
}

type IRequestHeader interface {
	Set(key string, value string)
	SetInt(key string, value int)
	SetInt32(key string, value int32)
	SetInt64(key string, value int64)
	SetBool(key string, value bool)
	Int(key string) (int, error)
	IntDefault(key string, def int) int
	Int32(key string) (int32, error)
	Int32Default(key string, def int32) int32
	Int64(key string) (int32, error)
	Int64Default(key string, def int64) int64
	Value(key string) string
	ValueDefault(key string, def string) string
	Bool(key string) (bool, error)
	BoolDefault(key string, def bool) bool
}

type IResponseHeader interface {
	Set(key string, value string)
	SetInt(key string, value int)
	SetInt64(key string, value int64)
	SetInt32(key string, value int32)
	SetBool(key string, value bool)
}

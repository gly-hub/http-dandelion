package params

type IParams interface {
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
	Parser(outObj interface{}) error
}

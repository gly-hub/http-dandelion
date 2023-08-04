package gin

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"reflect"
	"strconv"
)

type Params struct {
	Ctx *gin.Context
}

func (p Params) get(key string) string {
	return p.Ctx.Param(key)
}

func (p Params) Int(key string) (int, error) {
	vStr := p.get(key)
	return strconv.Atoi(vStr)
}

func (p Params) IntDefault(key string, def int) int {
	vStr := p.get(key)
	value, err := strconv.Atoi(vStr)
	if err != nil {
		return def
	}
	return value
}

func (p Params) Int32(key string) (int32, error) {
	vStr := p.get(key)
	value, err := strconv.ParseInt(vStr, 10, 32)
	if err != nil {
		return 0, err
	}
	return int32(value), nil
}

func (p Params) Int32Default(key string, def int32) int32 {
	vStr := p.get(key)
	value, err := strconv.ParseInt(vStr, 10, 32)
	if err != nil {
		return def
	}
	return int32(value)
}

func (p Params) Int64(key string) (int32, error) {
	vStr := p.get(key)
	value, err := strconv.ParseInt(vStr, 10, 32)
	if err != nil {
		return 0, err
	}
	return int32(value), nil
}

func (p Params) Int64Default(key string, def int64) int64 {
	vStr := p.get(key)
	value, err := strconv.ParseInt(vStr, 10, 64)
	if err != nil {
		return def
	}
	return value
}

func (p Params) Value(key string) string {
	return p.get(key)
}

func (p Params) ValueDefault(key string, def string) string {
	if p.get(key) == "" {
		return def
	}
	return p.get(key)
}

func (p Params) Bool(key string) (bool, error) {
	vStr := p.get(key)
	return strconv.ParseBool(vStr)
}

func (p Params) BoolDefault(key string, def bool) bool {
	vStr := p.get(key)
	value, err := strconv.ParseBool(vStr)
	if err != nil {
		return def
	}
	return value
}

func (p Params) Parser(outObj interface{}) error {
	rv := reflect.ValueOf(outObj)
	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}
	t := rv.Type()
	for i := 0; i < rv.NumField(); i++ {
		tag := t.Field(i).Tag

		// 获取 json tag 和 form tag
		jsonTag := tag.Get("json")
		var data any
		switch rv.Field(i).Type() {
		case reflect.TypeOf(string("")):
			data = p.get(jsonTag)
		case reflect.TypeOf(0):
			data, _ = p.Int(jsonTag)
		case reflect.TypeOf(int32(0)):
			data, _ = p.Int32(jsonTag)
		case reflect.TypeOf(int64(0)):
			data, _ = p.Int64(jsonTag)
		case reflect.TypeOf(true):
			data, _ = p.Bool(jsonTag)
		default:
			return errors.New(fmt.Sprintf("the field [%s] type [%s] is not supported",
				t.Field(i).Name, t.Field(i).Type))
		}

		rv.Field(i).Set(reflect.ValueOf(data))
	}
	return nil
}

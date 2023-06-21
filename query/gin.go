package query

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

type GinQuery struct {
	Ctx *gin.Context
}

func (p GinQuery) get(key string) string {
	return p.Ctx.Query(key)
}

func (p GinQuery) Int(key string) (int, error) {
	vStr := p.get(key)
	return strconv.Atoi(vStr)
}

func (p GinQuery) IntDefault(key string, def int) int {
	vStr := p.get(key)
	value, err := strconv.Atoi(vStr)
	if err != nil {
		return def
	}
	return value
}

func (p GinQuery) Int32(key string) (int32, error) {
	vStr := p.get(key)
	value, err := strconv.ParseInt(vStr, 10, 32)
	if err != nil {
		return 0, err
	}
	return int32(value), nil
}

func (p GinQuery) Int32Default(key string, def int32) int32 {
	vStr := p.get(key)
	value, err := strconv.ParseInt(vStr, 10, 32)
	if err != nil {
		return def
	}
	return int32(value)
}

func (p GinQuery) Int64(key string) (int32, error) {
	vStr := p.get(key)
	value, err := strconv.ParseInt(vStr, 10, 32)
	if err != nil {
		return 0, err
	}
	return int32(value), nil
}

func (p GinQuery) Int64Default(key string, def int64) int64 {
	vStr := p.get(key)
	value, err := strconv.ParseInt(vStr, 10, 64)
	if err != nil {
		return def
	}
	return value
}

func (p GinQuery) Value(key string) string {
	return p.get(key)
}

func (p GinQuery) ValueDefault(key string, def string) string {
	if p.get(key) == "" {
		return def
	}
	return p.get(key)
}

func (p GinQuery) Bool(key string) (bool, error) {
	vStr := p.get(key)
	return strconv.ParseBool(vStr)
}

func (p GinQuery) BoolDefault(key string, def bool) bool {
	vStr := p.get(key)
	value, err := strconv.ParseBool(vStr)
	if err != nil {
		return def
	}
	return value
}

func (p GinQuery) Parser(outObj interface{}) error {
	return p.Ctx.BindQuery(outObj)
}

package query

import (
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type FiberQuery struct {
	Ctx *fiber.Ctx
}

func (p FiberQuery) get(key string) string {
	return p.Ctx.Query(key)
}

func (p FiberQuery) Int(key string) (int, error) {
	vStr := p.get(key)
	return strconv.Atoi(vStr)
}

func (p FiberQuery) IntDefault(key string, def int) int {
	vStr := p.get(key)
	value, err := strconv.Atoi(vStr)
	if err != nil {
		return def
	}
	return value
}

func (p FiberQuery) Int32(key string) (int32, error) {
	vStr := p.get(key)
	value, err := strconv.ParseInt(vStr, 10, 32)
	if err != nil {
		return 0, err
	}
	return int32(value), nil
}

func (p FiberQuery) Int32Default(key string, def int32) int32 {
	vStr := p.get(key)
	value, err := strconv.ParseInt(vStr, 10, 32)
	if err != nil {
		return def
	}
	return int32(value)
}

func (p FiberQuery) Int64(key string) (int32, error) {
	vStr := p.get(key)
	value, err := strconv.ParseInt(vStr, 10, 32)
	if err != nil {
		return 0, err
	}
	return int32(value), nil
}

func (p FiberQuery) Int64Default(key string, def int64) int64 {
	vStr := p.get(key)
	value, err := strconv.ParseInt(vStr, 10, 64)
	if err != nil {
		return def
	}
	return value
}

func (p FiberQuery) Value(key string) string {
	return p.get(key)
}

func (p FiberQuery) ValueDefault(key string, def string) string {
	if p.get(key) == "" {
		return def
	}
	return p.get(key)
}

func (p FiberQuery) Bool(key string) (bool, error) {
	vStr := p.get(key)
	return strconv.ParseBool(vStr)
}

func (p FiberQuery) BoolDefault(key string, def bool) bool {
	vStr := p.get(key)
	value, err := strconv.ParseBool(vStr)
	if err != nil {
		return def
	}
	return value
}

func (p FiberQuery) Parser(outObj interface{}) error {
	return p.Ctx.QueryParser(outObj)
}

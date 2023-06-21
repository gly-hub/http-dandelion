package header

import (
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type FiberHeader struct {
	Ctx *fiber.Ctx
}

func (p FiberHeader) Request() IRequestHeader {
	return FiberRequestHeader{p.Ctx}
}

func (p FiberHeader) Response() IResponseHeader {
	return FiberResponseHeader{p.Ctx}
}

type FiberRequestHeader struct {
	Ctx *fiber.Ctx
}

func (p FiberRequestHeader) get(key string) string {
	return p.Ctx.Get(key)
}

func (p FiberRequestHeader) Set(key string, value string) {
	p.Ctx.Set(key, value)
}

func (p FiberRequestHeader) SetInt(key string, value int) {
	vStr := strconv.Itoa(value)
	p.Ctx.Set(key, vStr)
}

func (p FiberRequestHeader) SetInt32(key string, value int32) {
	vStr := strconv.FormatInt(int64(value), 10)
	p.Ctx.Set(key, vStr)
}

func (p FiberRequestHeader) SetInt64(key string, value int64) {
	vStr := strconv.FormatInt(value, 10)
	p.Ctx.Set(key, vStr)
}

func (p FiberRequestHeader) SetBool(key string, value bool) {
	vStr := strconv.FormatBool(value)
	p.Ctx.Set(key, vStr)
}

func (p FiberRequestHeader) Int(key string) (int, error) {
	vStr := p.get(key)
	return strconv.Atoi(vStr)
}

func (p FiberRequestHeader) IntDefault(key string, def int) int {
	vStr := p.get(key)
	value, err := strconv.Atoi(vStr)
	if err != nil {
		return def
	}
	return value
}

func (p FiberRequestHeader) Int32(key string) (int32, error) {
	vStr := p.get(key)
	value, err := strconv.ParseInt(vStr, 10, 32)
	if err != nil {
		return 0, err
	}
	return int32(value), nil
}

func (p FiberRequestHeader) Int32Default(key string, def int32) int32 {
	vStr := p.get(key)
	value, err := strconv.ParseInt(vStr, 10, 32)
	if err != nil {
		return def
	}
	return int32(value)
}

func (p FiberRequestHeader) Int64(key string) (int32, error) {
	vStr := p.get(key)
	value, err := strconv.ParseInt(vStr, 10, 32)
	if err != nil {
		return 0, err
	}
	return int32(value), nil
}

func (p FiberRequestHeader) Int64Default(key string, def int64) int64 {
	vStr := p.get(key)
	value, err := strconv.ParseInt(vStr, 10, 64)
	if err != nil {
		return def
	}
	return value
}

func (p FiberRequestHeader) Value(key string) string {
	return p.get(key)
}

func (p FiberRequestHeader) ValueDefault(key string, def string) string {
	if p.get(key) == "" {
		return def
	}
	return p.get(key)
}

func (p FiberRequestHeader) Bool(key string) (bool, error) {
	vStr := p.get(key)
	return strconv.ParseBool(vStr)
}

func (p FiberRequestHeader) BoolDefault(key string, def bool) bool {
	vStr := p.get(key)
	value, err := strconv.ParseBool(vStr)
	if err != nil {
		return def
	}
	return value
}

type FiberResponseHeader struct {
	Ctx *fiber.Ctx
}

func (f FiberResponseHeader) Set(key string, value string) {
	f.Ctx.Response().Header.Set(key, value)
}

func (f FiberResponseHeader) SetInt(key string, value int) {
	vStr := strconv.Itoa(value)
	f.Ctx.Response().Header.Set(key, vStr)
}

func (f FiberResponseHeader) SetInt32(key string, value int32) {
	vStr := strconv.FormatInt(int64(value), 10)
	f.Ctx.Response().Header.Set(key, vStr)
}

func (f FiberResponseHeader) SetInt64(key string, value int64) {
	vStr := strconv.FormatInt(value, 10)
	f.Ctx.Response().Header.Set(key, vStr)
}

func (f FiberResponseHeader) SetBool(key string, value bool) {
	vStr := strconv.FormatBool(value)
	f.Ctx.Response().Header.Set(key, vStr)
}

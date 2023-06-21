package header

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

type GinHeader struct {
	Ctx *gin.Context
}

func (p *GinHeader) Request() IRequestHeader {
	return GinRequestHeader{p.Ctx}
}

func (p *GinHeader) Response() IResponseHeader {
	return GinResponseHeader{p.Ctx}
}

type GinRequestHeader struct {
	Ctx *gin.Context
}

func (p GinRequestHeader) get(key string) string {
	return p.Ctx.GetHeader(key)
}

func (p GinRequestHeader) Set(key string, value string) {
	p.Ctx.Request.Header.Set(key, value)
}

func (p GinRequestHeader) SetInt(key string, value int) {
	vStr := strconv.Itoa(value)
	p.Ctx.Request.Header.Set(key, vStr)
}

func (p GinRequestHeader) SetInt32(key string, value int32) {
	vStr := strconv.FormatInt(int64(value), 10)
	p.Ctx.Request.Header.Set(key, vStr)
}

func (p GinRequestHeader) SetInt64(key string, value int64) {
	vStr := strconv.FormatInt(value, 10)
	p.Ctx.Request.Header.Set(key, vStr)
}

func (p GinRequestHeader) SetBool(key string, value bool) {
	vStr := strconv.FormatBool(value)
	p.Ctx.Request.Header.Set(key, vStr)
}

func (p GinRequestHeader) Int(key string) (int, error) {
	vStr := p.get(key)
	return strconv.Atoi(vStr)
}

func (p GinRequestHeader) IntDefault(key string, def int) int {
	vStr := p.get(key)
	value, err := strconv.Atoi(vStr)
	if err != nil {
		return def
	}
	return value
}

func (p GinRequestHeader) Int32(key string) (int32, error) {
	vStr := p.get(key)
	value, err := strconv.ParseInt(vStr, 10, 32)
	if err != nil {
		return 0, err
	}
	return int32(value), nil
}

func (p GinRequestHeader) Int32Default(key string, def int32) int32 {
	vStr := p.get(key)
	value, err := strconv.ParseInt(vStr, 10, 32)
	if err != nil {
		return def
	}
	return int32(value)
}

func (p GinRequestHeader) Int64(key string) (int32, error) {
	vStr := p.get(key)
	value, err := strconv.ParseInt(vStr, 10, 32)
	if err != nil {
		return 0, err
	}
	return int32(value), nil
}

func (p GinRequestHeader) Int64Default(key string, def int64) int64 {
	vStr := p.get(key)
	value, err := strconv.ParseInt(vStr, 10, 64)
	if err != nil {
		return def
	}
	return value
}

func (p GinRequestHeader) Value(key string) string {
	return p.get(key)
}

func (p GinRequestHeader) ValueDefault(key string, def string) string {
	if p.get(key) == "" {
		return def
	}
	return p.get(key)
}

func (p GinRequestHeader) Bool(key string) (bool, error) {
	vStr := p.get(key)
	return strconv.ParseBool(vStr)
}

func (p GinRequestHeader) BoolDefault(key string, def bool) bool {
	vStr := p.get(key)
	value, err := strconv.ParseBool(vStr)
	if err != nil {
		return def
	}
	return value
}

type GinResponseHeader struct {
	Ctx *gin.Context
}

func (f GinResponseHeader) Set(key string, value string) {
	f.Ctx.Set(key, value)
}

func (f GinResponseHeader) SetInt(key string, value int) {
	vStr := strconv.Itoa(value)
	f.Ctx.Set(key, vStr)
}

func (f GinResponseHeader) SetInt32(key string, value int32) {
	vStr := strconv.FormatInt(int64(value), 10)
	f.Ctx.Set(key, vStr)
}

func (f GinResponseHeader) SetInt64(key string, value int64) {
	vStr := strconv.FormatInt(value, 10)
	f.Ctx.Set(key, vStr)
}

func (f GinResponseHeader) SetBool(key string, value bool) {
	vStr := strconv.FormatBool(value)
	f.Ctx.Set(key, vStr)
}

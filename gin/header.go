package gin

import (
	"github.com/gin-gonic/gin"
	"github.com/gly-hub/http-dandelion/core"
	"strconv"
)

type Header struct {
	Ctx *gin.Context
}

func (p *Header) Request() core.IRequestHeader {
	return RequestHeader{p.Ctx}
}

func (p *Header) Response() core.IResponseHeader {
	return ResponseHeader{p.Ctx}
}

type RequestHeader struct {
	Ctx *gin.Context
}

func (p RequestHeader) get(key string) string {
	return p.Ctx.GetHeader(key)
}

func (p RequestHeader) Set(key string, value string) {
	p.Ctx.Request.Header.Set(key, value)
}

func (p RequestHeader) SetInt(key string, value int) {
	vStr := strconv.Itoa(value)
	p.Ctx.Request.Header.Set(key, vStr)
}

func (p RequestHeader) SetInt32(key string, value int32) {
	vStr := strconv.FormatInt(int64(value), 10)
	p.Ctx.Request.Header.Set(key, vStr)
}

func (p RequestHeader) SetInt64(key string, value int64) {
	vStr := strconv.FormatInt(value, 10)
	p.Ctx.Request.Header.Set(key, vStr)
}

func (p RequestHeader) SetBool(key string, value bool) {
	vStr := strconv.FormatBool(value)
	p.Ctx.Request.Header.Set(key, vStr)
}

func (p RequestHeader) Int(key string) (int, error) {
	vStr := p.get(key)
	return strconv.Atoi(vStr)
}

func (p RequestHeader) IntDefault(key string, def int) int {
	vStr := p.get(key)
	value, err := strconv.Atoi(vStr)
	if err != nil {
		return def
	}
	return value
}

func (p RequestHeader) Int32(key string) (int32, error) {
	vStr := p.get(key)
	value, err := strconv.ParseInt(vStr, 10, 32)
	if err != nil {
		return 0, err
	}
	return int32(value), nil
}

func (p RequestHeader) Int32Default(key string, def int32) int32 {
	vStr := p.get(key)
	value, err := strconv.ParseInt(vStr, 10, 32)
	if err != nil {
		return def
	}
	return int32(value)
}

func (p RequestHeader) Int64(key string) (int32, error) {
	vStr := p.get(key)
	value, err := strconv.ParseInt(vStr, 10, 32)
	if err != nil {
		return 0, err
	}
	return int32(value), nil
}

func (p RequestHeader) Int64Default(key string, def int64) int64 {
	vStr := p.get(key)
	value, err := strconv.ParseInt(vStr, 10, 64)
	if err != nil {
		return def
	}
	return value
}

func (p RequestHeader) Value(key string) string {
	return p.get(key)
}

func (p RequestHeader) ValueDefault(key string, def string) string {
	if p.get(key) == "" {
		return def
	}
	return p.get(key)
}

func (p RequestHeader) Bool(key string) (bool, error) {
	vStr := p.get(key)
	return strconv.ParseBool(vStr)
}

func (p RequestHeader) BoolDefault(key string, def bool) bool {
	vStr := p.get(key)
	value, err := strconv.ParseBool(vStr)
	if err != nil {
		return def
	}
	return value
}

type ResponseHeader struct {
	Ctx *gin.Context
}

func (f ResponseHeader) Set(key string, value string) {
	f.Ctx.Set(key, value)
}

func (f ResponseHeader) SetInt(key string, value int) {
	vStr := strconv.Itoa(value)
	f.Ctx.Set(key, vStr)
}

func (f ResponseHeader) SetInt32(key string, value int32) {
	vStr := strconv.FormatInt(int64(value), 10)
	f.Ctx.Set(key, vStr)
}

func (f ResponseHeader) SetInt64(key string, value int64) {
	vStr := strconv.FormatInt(value, 10)
	f.Ctx.Set(key, vStr)
}

func (f ResponseHeader) SetBool(key string, value bool) {
	vStr := strconv.FormatBool(value)
	f.Ctx.Set(key, vStr)
}

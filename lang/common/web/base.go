package web

import (
	"encoding/json"
	"lang/common/utils/structutils"
	"github.com/valyala/fasthttp"
	"strings"
)

const (
	ContentTypeName = "Content-Type"
	ContentTypeJson = "application/json;charset=utf-8"
	GetMethod = "GET"
	PostMethod = "POST"
	AllMethod = "ALL"
)

type Response struct {
	Code int `json:"code"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
}

func NewResponse(data interface{}) *Response  {
	return &Response{0,"success",data}
}

func WriteJson(ctx *fasthttp.RequestCtx,data interface{})  {
	ctx.Response.Header.Add(ContentTypeName,ContentTypeJson)
	json,error := json.Marshal(data)
	if (error != nil){
		panic(error)
	}
	ctx.Write(json)
}

//处理请求
type handler struct {
	method string //AllMethod表示处理所有请求
	handlerFuc (func (ctx *fasthttp.RequestCtx))
}

var handlerMap = make(map[string](*handler))

func Post(pattern string,h func (ctx *fasthttp.RequestCtx))  {
	hh,ok := handlerMap[pattern]
	if (ok && (hh.method == PostMethod || hh.method == AllMethod)){
		panic("uri " + pattern + " is exist ")
	}
	handlerMap[pattern] = &(handler{PostMethod,h})
}

func Get(pattern string,h func (ctx *fasthttp.RequestCtx))  {
	hh,ok := handlerMap[pattern]
	if (ok && (hh.method == GetMethod || hh.method == AllMethod)){
		panic("uri " + pattern + " is exist ")
	}
	handlerMap[pattern] = &(handler{GetMethod,h})
}

func Service(pattern string,h func (ctx *fasthttp.RequestCtx))  {
	_,ok := handlerMap[pattern]
	if (ok){
		panic("uri " + pattern + " is exist ")
	}
	handlerMap[pattern] = &(handler{AllMethod,h})
}

func Handler(ctx *fasthttp.RequestCtx) {
	path := string(ctx.Path())
	method := string(ctx.Method())

	h,ok := handlerMap[path]
	if ok {
		if h.method == AllMethod || h.method == method {
			h.handlerFuc(ctx)
		}else {
			ctx.NotFound()
		}
	}else {
		ctx.NotFound()
	}
}

func ReqToStruct(m map[string][]string,s interface{}) {
	structutils.StringArrayMapToStruct(m,s)
}

func ReqToMap(ctx *fasthttp.RequestCtx) map[string][]string {
	m := make(map[string][]string)
	queryArgs := (string(ctx.QueryArgs().QueryString()))
	args := strings.Split(queryArgs,"&")
	for i := 0; i < len(args); i++ {
		arg := strings.Split(args[i],"=")
		if (len(arg) == 2){
			v,ok := m[arg[0]]
			if ok {
				s := v[:]
				s = append(s,arg[1])
				m[arg[0]] = s
			}else {
				m[arg[0]] = []string{arg[1]}
			}
		}
	}

	for k,v := range m {
		if (len(v) > 1){
			args := ctx.QueryArgs().PeekMulti(k)
			ss := make([]string,len(args))
			for i := 0; i < len(args); i++ {
				ss[i] = string(args[i])
			}
			m[k] = ss
		}else {
			m[k] = []string{string(ctx.QueryArgs().Peek(k))}
		}
	}

	return m
}
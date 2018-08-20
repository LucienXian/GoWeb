package helper

import (
	. "net/http"
)

type UrlParams struct {
	Name string
	Value string
}

type PathConf struct {
	TemplatePath string
	JsPath string
	CssPath string
	ImagePath string
}

type Context struct {
	w ResponseWriter
	r *Request
	p []UrlParams
}

type Router interface {
	AddHandler(pattern string, method string, handler func(ctx *Context) ) bool
	Handle(handler interface{}) bool
	GetMatch(url string, method string) func(ctx *Context)
}

type WebServer struct {
	router Router
}



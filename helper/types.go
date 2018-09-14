package helper

import (
	. "net/http"
	"regexp"
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
	AddHandler(pattern string, method string, handler func(*Context) ) bool
	Handle(handler interface{}) bool
	GetMatch(url string, method string) (func(*Context), []UrlParams)
}

type WebServer struct {
	Route Router
	P404 func(*Context)
}

type RouterEntry struct {
	pattern string
	reg *regexp.Regexp
	handler map[string]func(*Context)
}

type WebRouters struct {
	router []RouterEntry
}




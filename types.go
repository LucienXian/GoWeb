package tinyWeb

import (
	"net/http"
	"regexp"
)

type urlParams struct {
	Name string
	Value []string
}

type pathConf struct {
	templatePath string
	jsPath string
	cssPath string
	imagePath string
}

type Context struct {
	W http.ResponseWriter
	R *http.Request
	P []urlParams
}

type router interface {
	AddHandler(pattern string, method string, handler func(*Context) ) bool
	Handle(handler interface{}) bool
	GetMatch(url string, method string, params map[string][]string) (func(*Context), []urlParams)
}

type WebServer struct {
	Route router
	P404 func(*Context)
}

type routerEntry struct {
	pattern string
	reg *regexp.Regexp
	handler map[string]func(*Context)
}

type webRouters struct {
	router []routerEntry
}




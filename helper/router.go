package helper

import (
	"strings"
)

type Router interface {
	AddHandler(pattern string, method string, handler func(*Context) ) bool
	Handle(handler interface{}) bool
	GetMatch(url string, method string) (func(*Context), []UrlParams)
}

type RouterEntry struct {
	pattern string
	handler map[string]func(*Context)
}

type WebRouters struct {
	router []Router
}

func (r *RouterEntry) AddHandler(pattern string, method string, handler func(*Context) ) bool {
	if len(pattern) < 1 {
		return false
	}
	if !strings.HasPrefix(pattern, "/") {
		pattern = "/" + pattern
	}
	if !(strings.HasPrefix(pattern, "^") || strings.HasPrefix(pattern, "\\A")) {
		pattern = "^" + pattern
	}
	if !strings.HasSuffix(pattern, "/") {
		pattern = pattern + "/"
	}
	if !(strings.HasSuffix(pattern, "$") || strings.HasPrefix(pattern, "\\z")) {
		pattern = pattern + "$"
	}
	return true
}


func (r *RouterEntry) Handle(handler interface{}) bool {
	return true
}


func (r *RouterEntry) GetMatch(url string, method string) (f func(*Context),params []UrlParams) {
	return
}
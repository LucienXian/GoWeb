package helper

import (
	"regexp"
	"strings"
)

type Router interface {
	AddHandler(pattern string, method string, handler func(*Context) ) bool
	Handle(handler interface{}) bool
	GetMatch(url string, method string) (func(*Context), []UrlParams)
}

type RouterEntry struct {
	pattern string
	reg *regexp.Regexp
	handler map[string]func(*Context)
}

type WebRouters struct {
	router []RouterEntry
}

func (r *WebRouters) AddHandler(pattern string, method string, handler func(*Context) ) bool {
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
		pattern = pattern + "/?"
	}
	if !(strings.HasSuffix(pattern, "$") || strings.HasPrefix(pattern, "\\z")) {
		pattern = pattern + "$"
	}
	for _, e := range r.router {
		if e.pattern == pattern {
			e.handler[strings.ToUpper(method)] = handler
			return true
		}
	}
	reg, err := regexp.Compile(pattern)
	if err != nil {
		return false
	}
	routerEntry := RouterEntry{pattern:pattern, reg:reg, handler:map[string]func(*Context){}}
	routerEntry.handler[strings.ToUpper(method)] = handler
	r.router = append(r.router, routerEntry)

	return true
}


func (r *RouterEntry) Handle(handler interface{}) bool {
	return true
}


func (r *RouterEntry) GetMatch(url string, method string) (f func(*Context),params []UrlParams) {
	return
}
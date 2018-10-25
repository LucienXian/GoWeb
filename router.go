package tinyWeb

import (
	"regexp"
	"strings"
)



func (r *webRouters) AddHandler(pattern string, method string, handler func(*Context) ) bool {
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
	routerEntry := routerEntry{pattern:pattern, reg:reg, handler:map[string]func(*Context){}}
	routerEntry.handler[strings.ToUpper(method)] = handler
	r.router = append(r.router, routerEntry)

	return true
}


func (r *webRouters) Handle(handler interface{}) bool {
	return true
}

func parseUrl(params map[string][]string) []urlParams   {
	if len(params) < 1 {
		return []urlParams{}
	}
	ret := make([]urlParams, len(params))
	i := 0
	for k, v := range params {
		ret[i].Name = k
		ret[i].Value = v
		i += 1
	}
	return ret
}

func (r *webRouters) GetMatch(url string, method string, pars map[string][]string) (f func(*Context),params []urlParams) {
	var exist bool
	for _, rou := range r.router {
		if rou.reg.MatchString(url) {
			if f, exist = rou.handler[strings.ToUpper(method)]; exist {
				params = parseUrl(pars)
				return
			}
		}
	}
	return
}
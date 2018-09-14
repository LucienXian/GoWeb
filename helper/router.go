package helper

import (
	"regexp"
	"strings"
)

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


func (r *WebRouters) Handle(handler interface{}) bool {
	return true
}

func parseUrl(re *regexp.Regexp, url string) []UrlParams   {
	values := re.FindStringSubmatch(url)
	if len(values) <= 1 {
		return []UrlParams{}
	}
	ret := make([]UrlParams, len(values)-1)
	names := re.SubexpNames()
	if len(names) == len(values) {
		for i:=0; i<len(ret); i++ {
			ret[i].Name = names[i+1]
			ret[i].Value = values[i+1]
		}
	} else {
		for i:=0; i<len(ret); i++ {
			ret[i].Name = ""
			ret[i].Value = values[i+1]
		}
	}
	return ret
}

func (r *WebRouters) GetMatch(url string, method string) (f func(*Context),params []UrlParams) {
	var exist bool
	for _, rou := range r.router {
		if rou.reg.MatchString(url) {
			if f, exist = rou.handler[strings.ToUpper(method)]; exist {
				params = parseUrl(rou.reg, url)
				return
			}
		}
	}
	return
}
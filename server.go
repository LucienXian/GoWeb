package tinyWeb

import (
	"tinyWeb/helper"
	. "net/http"
)

type WebServe helper.WebServer

func (s *WebServe) init() {
	
}

func (s *WebServe) ServeHTTP(w ResponseWriter, r *Request) {
	fun, para := s.Route.GetMatch(r.URL.Path, r.Method)
	c := &helper.Context{w, r, para}
	if fun != nil {
		s.P404(c)
	} else {
		fun(c)
	}

}
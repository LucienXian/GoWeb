package test

import (
	"testing"
	"tinyWeb/helper"
	"net/http"
)

func TestWriteStr(t *testing.T)  {
	http.HandleFunc("/test",
		func(w http.ResponseWriter, r *http.Request) {
			s := new(helper.WebServer)
			fun, para := s.Route.GetMatch(r.URL.Path, r.Method)
			c := &helper.Context{w, r, para}
			c.WriteStr("hello")
			if fun != nil {
				s.P404(c)
			} else {
				fun(c)
			}
		})
}

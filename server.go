package tinyWeb

import (
	"tinyWeb/helper"
	"net/http"
	"fmt"
)

type WebServe helper.WebServer

func (s *WebServe) init() {
	
}

func NewServer() *helper.WebServer {
	return &helper.WebServer{
		Route: new(helper.WebRouters),
		P404: func(c *helper.Context) {
			c.WriteStr("404 Not Found")
		},
	}
}

func (s *WebServe) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fun, para := s.Route.GetMatch(r.URL.Path, r.Method)
	c := &helper.Context{W:w, R:r, P:para}
	if fun != nil {
		s.P404(c)
	} else {
		fun(c)
	}

}

func genContext(handler interface{}) func(*helper.Context) {
	var s string
	switch v := handler.(type) {
	case string:
		s = v
	case func() string:
		s = v()
	default:
		fmt.Println("unknown")
	}
	c := func(c *helper.Context) {
		c.WriteStr(s)
	}
	return c
}

func Get(route string, handler interface{}) {
	c := genContext(handler)
	mainServer.Get(route, c)
	//mainServer.Handle(handler)
}

func Run(port string) {
	mainServer.Run(port)
}

var mainServer = NewServer()



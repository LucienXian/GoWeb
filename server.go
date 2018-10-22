package main

import (
	"fmt"
	"net/http"
	"tinyWeb/helper"
	. "net/http"
)

type WebServe helper.WebServer

func (s *WebServe) init() {
	
}

func (s *WebServe) ServeHTTP(w ResponseWriter, r *Request) {
	fmt.Print("sss\n")
	fmt.Println(s.Route)
	fun, para := s.Route.GetMatch(r.URL.Path, r.Method)
	c := &helper.Context{W:w, R:r, P:para}
	if fun != nil {
		s.P404(c)
	} else {
		fun(c)
	}

}


func main() {
	webRouters := new(helper.WebRouters)
	res := webRouters.AddHandler("/test", "GET", func(*helper.Context){
		fmt.Print("hello, world")
	})
	if res == false {
		fmt.Printf("add handler error")
		return
	}
	s := new(helper.WebServer)
	s.Route = webRouters
	s.P404 = func(c *helper.Context) {
		c.W.Write([]byte("404 Not Foun"))
	}
	http.HandleFunc("/test",
		func(w http.ResponseWriter, r *http.Request) {
			fun, para := s.Route.GetMatch(r.URL.Path, r.Method)
			c := &helper.Context{W:w, R:r, P:para}
			//fmt.Print(fun)
			if fun == nil {
				s.P404(c)
			} else {
				c.W.Write([]byte("hello"))
				fun(c)
			}
		})
	http.ListenAndServe(":12345", nil)
}
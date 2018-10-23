package helper

import (
	"net/http"
)

func (s *WebServer) SetRouter(r Router){
	s.Route = r
}

func (s *WebServer) Get(pattern string, entry func(*Context)) bool {
	s.Route.AddHandler(pattern, "GET", entry)
	return true
}

func (s *WebServer) Post(pattern string, entry func(*Context)) bool {
	return s.Route.AddHandler(pattern, "POST", entry)
}

func (s *WebServer) All(pattern string, entry func(*Context)) bool {
	return s.Route.AddHandler(pattern, "GET", entry) &&
		s.Route.AddHandler(pattern, "POST", entry)
}

func (s *WebServer) Run(addr string) {
	http.HandleFunc("/",
		func(w http.ResponseWriter, r *http.Request) {
			fun, para := s.Route.GetMatch(r.URL.Path, r.Method)
			c := &Context{W:w, R:r, P:para}
			if fun == nil {
				s.P404(c)
			} else {
				fun(c)
			}
		})
		http.ListenAndServe(addr, nil)
}

func (s *WebServer) Handle(entries interface{})  {
	s.Route.Handle(entries)
}


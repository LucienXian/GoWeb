package helper

import "fmt"

func (s *WebServer) SetRouter(r Router){
	s.Route = r
}

func (s *WebServer) Get(pattern string, entry func(*Context)) bool {
	// s.Route.AddHandler(pattern, "GET", entry)
	fmt.Println(s.Route)
	return true
}

func (s *WebServer) Post(pattern string, entry func(*Context)) bool {
	return s.Route.AddHandler(pattern, "POST", entry)
}

func (s *WebServer) All(pattern string, entry func(*Context)) bool {
	return s.Route.AddHandler(pattern, "GET", entry) &&
		s.Route.AddHandler(pattern, "POST", entry)
}

func (s *WebServer) Handle(entries interface{})  {
	s.Route.Handle(entries)
}


package tinyWeb

import (
	"os"
	"time"
	"path"
	"fmt"
	"net/http"
)

func (s *WebServer) SetRouter(r router){
	s.Route = r
}

func (s *WebServer) Get(pattern string, entry func(*Context)) bool {
	Info.Println("Add handler(get) ", pattern)
	s.Route.AddHandler(pattern, "GET", entry)
	return true
}

func (s *WebServer) Post(pattern string, entry func(*Context)) bool {
	Info.Println("Add handler(post) ", pattern)
	return s.Route.AddHandler(pattern, "POST", entry)
}

func (s *WebServer) All(pattern string, entry func(*Context)) bool {
	Info.Println("Add handler(get & post) ", pattern)
	return s.Route.AddHandler(pattern, "GET", entry) &&
		s.Route.AddHandler(pattern, "POST", entry)
}

func (s *WebServer) Run(addr string) {
	http.HandleFunc("/",
		func(w http.ResponseWriter, r *http.Request) {
			requestPath := r.URL.Path
			method := r.Method
			Info.Println(fmt.Sprintf("tinyWeb handle the request path %s, method %s", requestPath, method))
			r.ParseForm()
			fun, para := s.Route.GetMatch(requestPath, method, r.Form)
			c := &Context{W:w, R:r, P:para}
			c.setHeader("Server", "tinyWeb")
			c.setHeader("Date", time.Now().UTC().String())
			s.handle(r, w, c, fun)
		})
		http.ListenAndServe(addr, nil)
}

func (s *WebServer) handle(r *http.Request, w http.ResponseWriter, c *Context, fun func(*Context))  {
	requestPath := r.URL.Path
	if r.Method == "GET" {
		if s.serveStaticFile(requestPath, r, w) {
			/*c.setHeader("Content-Type", "text/plain; charse99t=utf-8")
			if fun != nil {
				fun(c)
			}
			*/
			return
		}
	}
	if fun != nil {
		fun(c)
	}
}

func (s *WebServer) serveStaticFile(name string, r *http.Request, w http.ResponseWriter) bool{
	for _, staticDir := range defaultStaticDirs {
		staticFile := path.Join(staticDir, name)
		if _, err := os.Stat(staticFile); !os.IsNotExist(err) {
			f, err := os.Open(staticFile)
			if err != nil {
				Error.Println(fmt.Sprintf("Serve static file %s error ", staticFile))
				return false
			}
			http.ServeContent(w, r, staticFile, time.Now(), f)
			return true
		}
	}
	return false
}

func newServer() *WebServer {
	return &WebServer{
		Route: new(webRouters),
		P404: func(c *Context) {
			c.WriteStr("404 Not Found")
		},
	}
}


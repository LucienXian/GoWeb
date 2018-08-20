package tinyWeb

import (
	"tinyWeb/helper"
	. "net/http"
)

type WebServer helper.WebServer

func (s *WebServer) init() {
	
}

func (s *WebServer) ServeHTTP(w ResponseWriter, r *Request) {
	
}
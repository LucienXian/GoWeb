package helper

import (
	"testing"
	"fmt"
)

func TestAddHandler(t *testing.T)  {
	s := new(WebServer)
	route := new(Router)
	s.Route = *route
	s.Get("test", func(context *Context) {
		fmt.Println("Get")
	})
}
package test

import (
	"tinyWeb/helper"
	"testing"
	"fmt"
)

func TestAddHandler(t *testing.T)  {
	s := new(helper.WebServer)
	route := new(helper.Router)
	s.Route = *route
	s.Get("test", func(context *helper.Context) {
		fmt.Println("Get")
	})
}
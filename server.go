package tinyWeb

import (
	"fmt"
	"os"
	"path"
)

var defaultStaticDirs []string

func init() {
	folderpath, _ := os.Getwd()
	defaultStaticDirs = append(defaultStaticDirs, path.Join(folderpath, "static"))
	return
}

func genContext(handler interface{}) func(*Context) {
	var s string
	switch v := handler.(type) {
	case string:
		s = v
	case func() string:
		s = v()
	case func(*Context):
		return v
	default:
		fmt.Println("unknown")
	}
	c := func(c *Context) {
		c.WriteStr(s)
	}
	return c
}

func Get(route string, handler interface{}) {
	c := genContext(handler)
	mainServer.Get(route, c)
	//mainServer.Handle(handler)
}

func Post(route string, handler interface{}) {
	c := genContext(handler)
	mainServer.Post(route, c)
}

func Run(port string) {
	Info.Println("tinyWeb run in the port: ", port)
	mainServer.Run(port)
}

var mainServer = newServer()



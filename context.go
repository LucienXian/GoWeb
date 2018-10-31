package tinyWeb

import (
	"encoding/json"
	"encoding/xml"
)

func (c *Context) WriteStr(s string) {
	c.setHeader("Content-Type", "text/plain; charset=utf-8")
	c.W.Write([]byte(s))
}

func (c *Context) WriteJson(v interface{}) {
	js, err := json.Marshal(v)
	if err != nil {
		Error.Println("Json Marshal error")
		c.Abort(500, "Server Error (json Marshal error)")
		return
	}
	c.setHeader("Content-Type", "application/json")
	c.W.Write(js)
}

func (c *Context) WriteXml(v interface{}) {
	x, err := xml.MarshalIndent(v, "", "  ")
	if err != nil {
		Error.Println("Xml Marshal error")
		c.Abort(500, "Server Error (xml Marshal error)")
		return
	}
	c.setHeader("Content-Type", "application/xml")
	c.W.Write(x)
}

func (c *Context) setHeader(name string, value string) {
	c.W.Header().Set(name, value)
}

func (c *Context) addHeader(name string, value string) {
	c.W.Header().Add(name, value)
}

func (c *Context) Abort(status int, body string) {
	c.setHeader("Content-Type", "text/html; charset=utf-8")
	c.W.WriteHeader(status)
	c.W.Write([]byte(body))
}

package tinyWeb

func (c *Context) WriteStr(s string) {
	c.W.Write([]byte(s))
}

func (c *Context) setHeader(name string, value string) {
	c.W.Header().Set(name, value)
}

func (c *Context) addHeader(name string, value string) {
	c.W.Header().Add(name, value)
}

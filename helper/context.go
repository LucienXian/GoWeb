package helper

func (c *Context) WriteStr(s string) {
	c.W.Write([]byte(s))
}
package helper

type Router interface {
	AddHandler(pattern string, method string, handler func(*Context) ) bool
	Handle(handler interface{}) bool
	GetMatch(url string, method string) (func(*Context), []UrlParams)
}



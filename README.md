# TinyWeb



tinyWeb is the simplest way to write web applications in the Go programming language, which supported  to handle the GET/POST method.



## Intsall

```shell
go get github.com/LucienXian/GoWeb
```



## Usage

Before we use the tinyWeb, we must write the handler. And also, you should point out the port you'd like to listen.



### Hanlder 1

```go
func helloworld() string{
	return "helloworld"
}

func main() {
	tinyWeb.Get("/index", helloworld)
	tinyWeb.Run(":12345")
}
```



### Handler 2

You need to use the struct **tinyWeb.Context**.

```go
func helloworld_2(ctx *tinyWeb.Context) {
	for k, v := range ctx.P {
		fmt.Println(k, v)
	}
}

func main() {
	tinyWeb.Get("/index", helloworld_2)
	tinyWeb.Run(":12345")
}
```



### Regexp

```go
func helloworld() string{
	return "helloworld"
}

func main() {
    //You can curl the URL start with '/'
	tinyWeb.Get("/(.+)", helloworld_abort) 
	tinyWeb.Run(":12345")
}
```



### Json & Xml

You can return the data as Json or Xml

```go
type Profile struct {
	Name    string
	Hobbies []string
}

func helloworld_3(ctx *tinyWeb.Context) {
	profile := Profile{"Alex", []string{"snowboarding", "programming"}}
	ctx.WriteJson(profile)
	//ctx.WriteXml(profile)
}

func main() {
	tinyWeb.Get("/index", helloworld_3) 
	tinyWeb.Run(":12345")
}
```



### Static file

The program will find the static directory in the same path as the excutable program. 

```go
func main() {
	tinyWeb.Get("/index.html", helloworld)
	tinyWeb.Run(":12345")
}
```



## To Do

* Dynamic Config
* More http method
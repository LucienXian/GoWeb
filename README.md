# TinyWeb



tinyWeb is the simplest way to write web applications in the Go programming language.



## Example

```go
package main

import (
	"tinyWeb"
	"fmt"
)

type Profile struct {
	Name    string
	Hobbies []string
  }

func helloworld() string{
	return "helloworld"
}

func helloworld_2(ctx *tinyWeb.Context) {
	for k, v := range ctx.P {
		fmt.Println(k, v)
	}
}

func helloworld_3(ctx *tinyWeb.Context) {
	profile := Profile{"Alex", []string{"snowboarding", "programming"}}
	ctx.WriteJson(profile)
	//ctx.WriteXml(profile)
}

func helloworld_abort(ctx *tinyWeb.Context) {
	ctx.WriteJson(make(chan int))
}

func main() {
	tinyWeb.Get("/(.+)", helloworld_abort)
	tinyWeb.Run(":12345")
}
```


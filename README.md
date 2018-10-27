# TinyWeb



tinyWeb is the simplest way to write web applications in the Go programming language.



## Example

```go
package main

import (
	"tinyWeb"
	"fmt"
)

func helloworld() string{
	return "helloworld"
}

func helloworld_2(ctx *tinyWeb.Context) {
	for k, v := range ctx.P {
		fmt.Println(k, v)
	}
}

func main() {
	//tinyWeb.Get("/test", helloworld_2)
	tinyWeb.Get("/(.+)", helloworld_2) //curl -i localhost:12345/index.html?a=1
	tinyWeb.Run(":12345")
}
```


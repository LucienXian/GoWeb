# TinyWeb



tinyWeb is the simplest way to write web applications in the Go programming language.



## Example

```go
package main

import (
	"tinyWeb"
)

func helloworld() string{
	return "helloworld"
}

func main() {
	tinyWeb.Get("/test", helloworld)
	tinyWeb.Run(":12345")
}
```


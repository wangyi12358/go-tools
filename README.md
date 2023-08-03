# GO tools

## Install
```shell
go get github.com/wangyi12358/go-tools@latest
```

## Usage
```go
package main

import (
	"github.com/wangyi12358/go-tools/array"
)

func main() {
	array.find([]string{"a", "b", "c"}, func(v string) bool {
		return v == "a"
	})
}
```
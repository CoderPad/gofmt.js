package main

import (
	"github.com/gopherjs/gopherjs/js"
	"go/format"
)

func Format(source string) interface{} {
	ret, err := format.Source([]byte(source))
	if err == nil {
		return string(ret)
	} else {
		return false
	}
}

func main() {
	if js.Module != js.Undefined {
		js.Module.Set("exports", Format)
	} else {
		js.Global.Set("gofmt", Format)
	}
}

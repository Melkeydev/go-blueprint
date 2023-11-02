package template

type BeegoTemplates struct{}

func (b BeegoTemplates) Main() []byte {
	return MainTemplate()
}
func (g BeegoTemplates) Server() []byte {
	return MakeHTTPServer()
}
func (g BeegoTemplates) Routes() []byte {
	return MakeBeegoRoutes()
}

func MakeBeegoRoutes() []byte {
	return []byte(`package main

import (
	"github.com/beego/beego/v2/server/web"
	"fmt"
)

func main() {
	fmt.println("Running at localhost:8080")
	web.Run()
}

`)
}

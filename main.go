package main

import (
	"github.com/astaxie/beego"
	_ "fmt"
	_ "dockerbee/routers"
	_ "github.com/fsouza/go-dockerclient"
)

func main() {
	beego.TemplateLeft = "$$"
	beego.TemplateRight = "**"
	beego.Run()
}


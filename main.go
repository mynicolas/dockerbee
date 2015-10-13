package main

import (
	_ "fmt"
	_ "dockerbee/routers"
	"github.com/astaxie/beego"
	_ "github.com/fsouza/go-dockerclient"
)

func main() {
	beego.SetStaticPath("/polymer", "polymer")
	beego.Run()
}


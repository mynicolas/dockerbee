package main

import (
	_ "dockerbee/routers"
	"github.com/astaxie/beego"
	_ "github.com/fsouza/go-dockerclient"
)

func main() {
	beego.Run()
}


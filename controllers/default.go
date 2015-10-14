package controllers

import (
	_ "fmt"
	"github.com/astaxie/beego"
	"github.com/fsouza/go-dockerclient"
)

type TestController struct {
	beego.Controller
}
func (t *TestController) Get() {
	t.TplNames = "test/test.tpl"
	t.LayoutSections = make(map[string]string)
	t.LayoutSections["Title"] = "test/title.tpl"
	t.LayoutSections["DockerElements"] = "index/docker-main.tpl"
	t.Layout = "index/base.tpl"
}

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.TplNames = "index/index.tpl"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["Title"] = "index/title.tpl"
	c.LayoutSections["DockerElements"] = "index/docker-main.tpl"
	c.Layout = "index/base.tpl"
}

type OverviewController struct {
	beego.Controller
}

func (o *OverviewController) Get() {
	o.TplNames = "overview/overview.tpl"
	o.LayoutSections = make(map[string]string)
	o.LayoutSections["Title"] = "overview/title.tpl"
	o.LayoutSections["DockerElements"] = "overview/docker-main.tpl"
	o.Layout = "index/base.tpl"
}

type ContainersController struct {
	beego.Controller
}

func (o *ContainersController) Get() {
	endpoint := "tcp://172.16.4.112:2375"
	client, _ := docker.NewClient(endpoint)
	containers, _ := client.ListContainers(docker.ListContainersOptions{All: true})

	o.Data["Conts"] = containers
	o.TplNames = "containers/containers.tpl"
	o.LayoutSections = make(map[string]string)
	o.LayoutSections["Title"] = "containers/title.tpl"
	o.LayoutSections["DockerElements"] = "containers/docker-main.tpl"
	o.Layout = "index/base.tpl"
}
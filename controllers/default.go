package controllers

import (
	"github.com/astaxie/beego"
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
	Conts := []struct {
		Name string
		Status string		
	}{{"container1", "normal"},{"container2", "normal"},{"container3", "error"}}

	o.Data["Conts"] = Conts
	o.TplNames = "containers/containers.tpl"
	o.LayoutSections = make(map[string]string)
	o.LayoutSections["Title"] = "containers/title.tpl"
	o.LayoutSections["DockerElements"] = "containers/docker-main.tpl"
	o.Layout = "index/base.tpl"
}
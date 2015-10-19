package controllers

import (
	"github.com/astaxie/beego"
	"github.com/fsouza/go-dockerclient"
)

type TestController struct {
	beego.Controller
}
func (this *TestController) Get() {
	this.TplNames = "test/test.tpl"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["Title"] = "test/title.tpl"
	this.LayoutSections["DockerElements"] = "index/docker-main.tpl"
	this.Layout = "index/base.tpl"
}

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.TplNames = "index/index.tpl"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["Title"] = "index/title.tpl"
	this.LayoutSections["DockerElements"] = "index/docker-main.tpl"
	this.Layout = "index/base.tpl"
}

type OverviewController struct {
	beego.Controller
}

func (this *OverviewController) Get() {
	this.TplNames = "overview/overview.tpl"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["Title"] = "overview/title.tpl"
	this.LayoutSections["DockerElements"] = "overview/docker-main.tpl"
	this.Layout = "index/base.tpl"
}

type ContainersController struct {
	beego.Controller
}

func (this *ContainersController) Containers() {
	endpoint := beego.AppConfig.String("docker_endpoint")
	client, _ := docker.NewClient(endpoint)
	containers, _ := client.ListContainers(docker.ListContainersOptions{All: true})

	this.Data["Conts"] = containers
	this.TplNames = "containers/containers.tpl"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["Title"] = "containers/title.tpl"
	this.LayoutSections["DockerElements"] = "containers/docker-main.tpl"
	this.Layout = "index/base.tpl"
}

type ImagesController struct {
	beego.Controller
}

func (this *ImagesController) Images() {
	endpoint := beego.AppConfig.String("docker_endpoint")
	client, _ := docker.NewClient(endpoint)
	images, _ := client.ListImages(docker.ListImagesOptions{All: false})

	this.Data["Imgs"] = images
	this.TplNames = "images/images.tpl"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["Title"] = "images/title.tpl"
	this.LayoutSections["DockerElements"] = "images/docker-main.tpl"
	this.Layout = "index/base.tpl"
}

func (this *ContainersController) Pause() {
	containerId := this.Input().Get("id")
	
	endpoint := beego.AppConfig.String("docker_endpoint")
	client, _ := docker.NewClient(endpoint)
	result := client.PauseContainer(containerId)

	if result == nil {
		this.Ctx.WriteString("succeed")
	} else {
		this.Ctx.WriteString(result.Error())
	}
}

func (this *ContainersController) Unpause() {
	containerId := this.Input().Get("id")
	
	endpoint := beego.AppConfig.String("docker_endpoint")
	client, _ := docker.NewClient(endpoint)
	result := client.UnpauseContainer(containerId)

	if result == nil {
		this.Ctx.WriteString("succeed")
	} else {
		this.Ctx.WriteString(result.Error())
	}
}

func (this *ContainersController) Stop() {
	containerId := this.Input().Get("id")
	
	endpoint := beego.AppConfig.String("docker_endpoint")
	client, _ := docker.NewClient(endpoint)
	result := client.StopContainer(containerId, 3)

	if result == nil {
		this.Ctx.WriteString("succeed")
	} else {
		this.Ctx.WriteString(result.Error())
	}
}

func (this *ContainersController) Start() {
	containerId := this.Input().Get("id")
	
	endpoint := beego.AppConfig.String("docker_endpoint")
	client, _ := docker.NewClient(endpoint)
	result := client.StartContainer(containerId, &docker.HostConfig{PublishAllPorts: true})

	if result == nil {
		this.Ctx.WriteString("succeed")
	} else {
		this.Ctx.WriteString(result.Error())
	}
}

func (this *ContainersController) Restart() {
	containerId := this.Input().Get("id")
	
	endpoint := beego.AppConfig.String("docker_endpoint")
	client, _ := docker.NewClient(endpoint)
	result := client.RestartContainer(containerId, 3)

	if result == nil {
		this.Ctx.WriteString("succeed")
	} else {
		this.Ctx.WriteString(result.Error())
	}
}

func (this *ContainersController) Kill() {
	containerId := this.Input().Get("id")
	
	endpoint := beego.AppConfig.String("docker_endpoint")
	client, _ := docker.NewClient(endpoint)
	result := client.KillContainer(docker.KillContainerOptions{ID: containerId})

	if result == nil {
		this.Ctx.WriteString("succeed")
	} else {
		this.Ctx.WriteString(result.Error())
	}
}

func (this *ContainersController) Remove() {
	containerId := this.Input().Get("id")
	
	endpoint := beego.AppConfig.String("docker_endpoint")
	client, _ := docker.NewClient(endpoint)
	result := client.RemoveContainer(docker.RemoveContainerOptions{ID: containerId, Force: true})

	if result == nil {
		this.Ctx.WriteString("succeed")
	} else {
		this.Ctx.WriteString(result.Error())
	}
}

func (this *ContainersController) Inspect() {
	containerId := this.Input().Get("id")
	
	endpoint := beego.AppConfig.String("docker_endpoint")
	client, _ := docker.NewClient(endpoint)
	container, err := client.InspectContainer(containerId)

	if err == nil {
		this.Data["json"] = &container
		this.ServeJson()
	} else {
		this.Ctx.WriteString(err.Error())
	}
}
package controller

import(
	"release_manager/model"
	"release_manager/domain"
	"release_manager/service"
	"fmt"
)

type HomeController struct {
	BaseController
	siteId int
	siteInfo *domain.SiteConfig
}

func (this *HomeController) Before() {
	(&this.BaseController).Before()
	model.User.HasLoggedin(this.Ctx, true)

	this.siteId = this.Param.GetInt("siteId")
	this.View.Set("siteId", this.siteId)
	this.View.Set("siteInfo", nil)
	if this.siteId > 0 {
		var err error
		this.siteInfo, err = model.SiteConfig.GetById(this.siteId)
		if err!= nil {
			panic(err)
		}
		this.View.Set("siteInfo", this.siteInfo)
	}
}

func (this *HomeController) IndexAction() {
	var task = this.Param.Get("task")
	var releaseType = this.Param.Get("releaseType")
	var filterKeywords = this.Param.Get("keyWords")

	var frameLink = fmt.Sprintf(`src="/home/runCommand?siteId=%d&task=%s&releaseType=%s"`, 
		this.siteId, task, releaseType)
	this.View.Set("frameLink", frameLink)
	this.View.Set("releaseType", releaseType)
	this.View.Set("keyWords", filterKeywords)

	sites, err := model.SiteConfig.List(0, 1000)
	if err!=nil {
		panic(err)
	}
	this.View.Set("sites", sites)

	this.Render("home/index")
}

func (this *HomeController) RunCommandAction() {
	var command = this.Param.Get("task")
	var siteId = this.Param.GetInt("siteId")

	var releasService = service.NewReleaseService()
	releasService.RunCommand(this.Ctx, siteId, command)
	return
}

func (this *HomeController) HiAction() string {
	return "hi " + this.Param["username"]
}

func (this *HomeController) ErrAction() string {
	panic("this is a test!")
}

package controller

import(
	"release_manager/model"
	"release_manager/domain"
)

type SiteConfigController struct {
	BaseController
}

func (this *SiteConfigController) Before() {
	(&this.BaseController).Before()
	model.User.HasLoggedin(this.Ctx, true)
	model.User.IsAdmin(this.Ctx, true)
}

func (this *SiteConfigController) IndexAction() {
	//list in single page
	var page = 1
	var numPerPage= 1000

	var offset = (page - 1) * numPerPage
	sites, err := model.SiteConfig.List(offset, numPerPage)
	if err!= nil {
		panic(err)
	}

	this.View.Set("sites", sites)
	this.Render("siteconfig/list")
}

func (this *SiteConfigController) AddAction() {
	var siteconfig = &domain.SiteConfig{}
	this.View.Set("siteconfig", siteconfig)
	this.View.Set("errStr", "")
	this.Render("siteconfig/add")
}

func (this *SiteConfigController) DoAddAction() {
	
	var siteconfig = &domain.SiteConfig{
		Id: 0,
		Sitename: this.Param.Get("sitename"),
		Base_dir: this.Param.Get("base_dir"),
		Get_current_branch_command: this.Param.Get("get_current_branch_command"),
		Update_command: this.Param.Get("update_command"),
		Generate_command: this.Param.Get("generate_command"),
		Test_release_command: this.Param.Get("test_release_command"),
		Release_command: this.Param.Get("release_command"),
		Cache_dir: this.Param.Get("cache_dir"),
		Cache_exclude_dir: this.Param.Get("cache_exclude_dir"),
		Cache_urls: this.Param.Get("cache_urls"),
	}

	var errmsg = ""
	if siteconfig.Sitename == "" {
		errmsg = "Invalid sitename";
	}
			
	if errmsg == "" {
		_, err := model.SiteConfig.Add(siteconfig)
		if err!= nil {
			errmsg = "new siteconfig failed"
		}else{
			this.Redirect("/siteconfig")
			this.Exit()
		}
	}

	this.View.Set("siteconfig", siteconfig)
	this.View.Set("errStr", errmsg)
	this.Render("siteconfig/add")
}

func (this *SiteConfigController) EditAction() {
	
	var id = this.Param.GetInt("id")

	siteconfig, err := model.SiteConfig.GetById(id)
	if err!=nil {
		panic(err)
	}
	this.View.Set("siteconfig", siteconfig)
	this.View.Set("errStr", "")
	
	this.Render("siteconfig/edit")
}

func (this *SiteConfigController) DoEditAction() {

	var siteconfig = &domain.SiteConfig{
		Id: this.Param.GetInt("id"),
		Sitename: this.Param.Get("sitename"),
		Base_dir: this.Param.Get("base_dir"),
		Get_current_branch_command: this.Param.Get("get_current_branch_command"),
		Update_command: this.Param.Get("update_command"),
		Generate_command: this.Param.Get("generate_command"),
		Test_release_command: this.Param.Get("test_release_command"),
		Release_command: this.Param.Get("release_command"),
		Cache_dir: this.Param.Get("cache_dir"),
		Cache_exclude_dir: this.Param.Get("cache_exclude_dir"),
		Cache_urls: this.Param.Get("cache_urls"),
	}

	var errmsg = ""

	if siteconfig.Id < 1 {
		errmsg = "Invalid id";
	}
	if siteconfig.Sitename == "" {
		errmsg = "Please fill out the Site Name!";
	}
	if errmsg == "" {
		_, err := model.SiteConfig.Update(siteconfig)
		if err!=nil {
			errmsg = "update siteconfig failed"
		}else{
			this.Redirect("/siteconfig")
			this.Exit()
		}
	}

	this.View.Set("siteconfig", siteconfig)
	this.View.Set("errStr", errmsg)
	this.Render("siteconfig/edit")
}

func (this *SiteConfigController) DeleteAction() {

	var id = this.Param.GetInt("id")

	_,err := model.SiteConfig.Delete(id)
	if err!= nil {
		panic(err)
	}
	this.Redirect("/siteconfig")
}
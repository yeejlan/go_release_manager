package controller

import(
	"release_manager/model"
	"release_manager/domain"
)

type AdminController struct {
	BaseController
}

type userPost struct {
	domain.User
	Confirmpassword string
}

func (this *AdminController) Before() {
	(&this.BaseController).Before()
	model.User.HasLoggedin(this.Ctx, true)
	model.User.IsAdmin(this.Ctx, true)
}

func (this *AdminController) IndexAction() {
	//user list in single page
	var page = 1
	var numPerPage= 1000

	var offset = (page - 1) * numPerPage
	userList, err := model.Admin.ListUser(offset, numPerPage)
	if err!= nil {
		panic(err)
	}

	this.View.Set("userlist", userList)
	this.Render("admin/list")
}

func (this *AdminController) AddAction() {
	var user = &userPost{}
	this.View.Set("post", user)
	this.View.Set("errStr", "")
	this.Render("admin/add")
}

func (this *AdminController) DoAddAction() {
	
	var username = this.Param.Get("username")
	var password = this.Param.Get("password")
	var confirmpassword = this.Param.Get("confirmpassword")
	var role = this.Param.Get("role")

	var errmsg = ""

	if username == "" || password == "" || confirmpassword == "" {
		errmsg = "Please fill out Username/Password/Confirm Password!";
	}
	if password != confirmpassword {
		errmsg = "Confirm Password does NOT match Password!";
	}
			
	if errmsg == "" {
		_, err := model.Admin.NewUser(this.Ctx, username, password, role)
		if err!= nil {
			errmsg = "new user failed"
		}else{
			this.Redirect("/admin")
			this.Exit()
		}
	}
	var user = &domain.User{
		Username: username,
		Password: password,
		Role: role,
	}

	var userPost = &userPost{
		User: *user,
		Confirmpassword: confirmpassword,
	}
	this.View.Set("post", userPost)
	this.View.Set("errStr", errmsg)
	this.Render("admin/add")
}

func (this *AdminController) EditAction() {
	
	var id = this.Param.GetInt("id")

	user, err := model.User.GetUserById(id)
	if err != nil {
		panic(err)
	}

	this.View.Set("user", user)
	this.View.Set("errStr", "")
	this.Render("admin/edit")
}

func (this *AdminController) DoEditAction() {

	var id = this.Param.GetInt("id")
	var username = this.Param.Get("username")
	var password = this.Param.Get("password")
	var confirmpassword = this.Param.Get("confirmpassword")
	var role = this.Param.Get("role")

	var errmsg = ""

	if password == "" || confirmpassword == "" {
		errmsg = "Please fill out Password/Confirm Password!";
	}
	if password != confirmpassword {
		errmsg = "Confirm Password does NOT match Password!";
	}
	if errmsg == "" {
		_, err := model.Admin.UpdateUser(this.Ctx, id, password, role)
		if err != nil {
			errmsg = "update user failed"
		}else{
			this.Redirect("/admin")
			this.Exit()
		}
	}

	var user = &domain.User{
		Id: id,
		Username: username,
		Password: "",
		Role: role,
	}
	this.View.Set("user", user)
	this.View.Set("errStr", errmsg)
	this.Render("admin/edit")
}

func (this *AdminController) DeleteAction(){

	var id = this.Param.GetInt("id")

	_, err := model.Admin.DeleteUser(id)
	if err != nil {
		panic(err)
	}
	this.Redirect("/admin")
}
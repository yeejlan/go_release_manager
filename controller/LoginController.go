package controller

import(
	"github.com/yeejlan/maru"
	"release_manager/model"
	"release_manager/lib"
)

type LoginController struct {
	BaseController
}

func (this *LoginController) Before() {
	(&this.BaseController).Before()
	this.View.Set("errStr", "")
}

func (this *LoginController) IndexAction(){
	errno := this.Param.GetInt("err")

	if errno == 1 {
		this.View.Set("errStr", "Invalid username or password, Please try again!")
	}
	if errno == 2 {
		this.View.Set("errStr", "Internal error, Please try again later!")
	}

	this.Render("login/index")
}

func (this *LoginController) PostAction() {
	username := this.Param.Get("username")
	password := this.Param.Get("password")
	clientIp := lib.Utils.GetClientIp(this.Req)
	loginResult, err := model.User.Login(this.Ctx, username, password, clientIp)
	if err != nil {
		maru.Log("login", "login error: " + err.Error())
		this.Redirect("/login?err=2")
		return
	}
	if loginResult == true {
		this.Redirect("/")
		return
	}
	this.Redirect("/login?err=1")
	return
}

func (this *LoginController) ExitAction() {
	this.Session.Destroy()
	this.Redirect("/login")
	return
}

func (this *LoginController)  ChangePasswordAction() {
	model.User.HasLoggedin(this.Ctx, true)

	oldPassword := this.Param.Get("oldpassword")
	newPassword := this.Param.Get("newpassword")
	confirmPassword := this.Param.Get("confirmpassword")

	this.View.Set("msg", "")
	this.View.Set("oldPassword", oldPassword)
	this.View.Set("newPassword", newPassword)
	this.View.Set("confirmPassword", confirmPassword)
	if len(this.Param)> 0 {
		var msg = ""
		var userid = this.Session.GetInt("uid")
		verifyResult, err := model.User.VerifyPassword(this.Ctx, userid, oldPassword)
		if err!=nil {
			msg = "Internal error on verifyPassword"
		}
		if msg == "" && oldPassword != "" && verifyResult == false {
			msg = "Old Password is wrong!"
		}
		if msg == "" && (newPassword == "" || confirmPassword == "") {
			msg = "New Password or Comfirm New Passsword cannot be empty!"
		}
		if msg == "" && (newPassword != confirmPassword) {
			msg = "Confirm New Password do NOT match New Password!"
		}
		if msg == "" {
			_, err = model.User.ChangePassword(this.Ctx, userid, newPassword)
			if err!=nil {
				msg = "Internal error on changePassword"
			}else {
				msg = "Your password updated successfully!"
			}
		}
		this.View.Set("msg", msg)
	}
	this.Render("login/changepassword")
}
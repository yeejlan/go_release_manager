package model

import(
	"github.com/yeejlan/maru"
	"release_manager.domain"
)

var (
	User = &userModel{}
)

type userModel struct{}

func (this *userModel) CurrentUserId(ctx *maru.WebContext) Int {
	return ctx.Session.GetInt("uid")
}

func (this *userModel) CurrentRole(ctx *maru.WebContext) string {
	return ctx.Session.GetString("role")
}

func (this *userModel) CurrentUserInfo(ctx *maru.WebContext) (*domain.User, error) {
	userid := this.CurrentUserId(ctx)
	return this.GetUserById(userid)
}

/*call this function when a page need user auth*/
func (this *userModel) HasLoggedin(ctx *maru.WebContext, loginPageRedirect bool) bool {
	if this.CurrentUserId(ctx) > 0 {
		return true
	}
	if loginPageRedirect {
		ctx.Redirect("/login")
		ctx.Exit()
	}
	return false
}

/*call this function when a page need admin privilege*/
func (this *userModel) IsAdmin(pageRedirect bool) bool {
	if this.CurrentRole() == "admin" {
		return true
	}
	if(pageRedirect) {
		ctx.Redirect("/")
		ctx.Exit()
	}
	return false
}

func (this *userModel) GetUserById(userid int) (*domain.User, error) {
	return dao.User.getUserById(userid)
}





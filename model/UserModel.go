package model

import(
	"github.com/yeejlan/maru"
	"release_manager/domain"
	"release_manager/dal/dao"
	"crypto/md5"
	"fmt"
)

var (
	User = &userModel{}
)

type userModel struct{}

func (this *userModel) CurrentUserId(ctx *maru.Ctx) int {
	return ctx.Session.GetInt("uid")
}

func (this *userModel) CurrentRole(ctx *maru.Ctx) string {
	return ctx.Session.GetString("role")
}

func (this *userModel) CurrentUserInfo(ctx *maru.Ctx) (*domain.User, error) {
	userid := this.CurrentUserId(ctx)
	return this.GetUserById(userid)
}

/*call this function when a page need user auth*/
func (this *userModel) HasLoggedin(ctx *maru.Ctx, loginPageRedirect bool) bool {
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
func (this *userModel) IsAdmin(ctx *maru.Ctx, pageRedirect bool) bool {
	if this.CurrentRole(ctx) == "admin" {
		return true
	}
	if(pageRedirect) {
		ctx.Redirect("/")
		ctx.Exit()
	}
	return false
}

func (this *userModel) GetUserById(userid int) (*domain.User, error) {
	return dao.User.GetUserById(userid)
}

//user login
func (this *userModel) Login(ctx *maru.Ctx, username string, password string, clientIp string) (bool, error) {
	if username == "" || password == "" {
		return false, nil
	}

	user, err := dao.User.GetUserByName(username)
	if err != nil {
		return false, err
	}
	passwordMd5 := this.GetPasswordMd5(ctx, password)
	if user.Password == passwordMd5 { //login success
		session := ctx.Session
		session.Set("uid", user.Id)
		session.Set("username", user.Username)
		session.Set("role", user.Role)

		ActionLog.Add(user.Id, user.Username, "login", "Success", clientIp)

		return true, nil
	}

	ActionLog.Add(-1, username, "login", "Failed", clientIp)

	return false, nil
}

func (this *userModel) VerifyPassword(ctx *maru.Ctx, userid int, password string) (bool, error) { 
	if userid < 1 || password == "" {
		return false, nil
	}

	user, err := dao.User.GetUserById(userid)
	if err !=nil {
		return false, err
	}
	passwordMd5 := this.GetPasswordMd5(ctx, password)
	if user.Password == passwordMd5 {
		return true, nil
	}
	return false, nil
}

func (this *userModel) ChangePassword(ctx *maru.Ctx, userid int, password string) (bool, error) {
	if userid < 1 || password == "" {
		return false, maru.NewError("bad param")
	}
	passwordMd5 := this.GetPasswordMd5(ctx, password)
	_, err := dao.User.UpdatePassword(userid, passwordMd5)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (this *userModel) GetPasswordMd5(ctx *maru.Ctx, password string) string {
	sitePhrase := ctx.App.Config().Get("site.phrase")
	data := []byte(password + sitePhrase)
	return fmt.Sprintf("%x", md5.Sum(data))
}
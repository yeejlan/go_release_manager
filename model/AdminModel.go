package model

import(
	"github.com/yeejlan/maru"
	"release_manager/domain"
	"release_manager/dal/dao"
)

var (
	Admin = &adminModel{}
)

type adminModel struct{}

func (this *adminModel) NewUser(ctx *maru.WebContext, username string, password string, role string) (int, error) {
	passwordMd5 := User.GetPasswordMd5(ctx, password)
	user := &domain.User{
		Username: username,
		Password: passwordMd5,
		Role: role,
	}
	return dao.User.Add(user)
}

func (this *adminModel) UpdateUser(ctx *maru.WebContext, userid int, password string, role string) (int, error) {
	passwordMd5 := ""
	if password != "" {
		passwordMd5 = User.GetPasswordMd5(ctx, password)
	}
	return dao.User.Update(userid, passwordMd5, role)
}

func (this *adminModel) ListUser(offset int, pageSize int) (result []domain.User, err error) {
	return dao.User.List(offset, pageSize)
}

func (this *adminModel) isUsernameAvailable(username string) bool {
	user, _ := dao.User.GetUserByName(username)
	if user != nil && user.Id > 0 {
		return false
	}
	return true
}

func (this *adminModel) DeleteUser(userid int) (int, error) {
	return dao.User.Delete(userid)
}


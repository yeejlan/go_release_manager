package dao

import (
	"github.com/yeejlan/maru"
	"release_manager/dal"
	"release_manager/domain"
)

var (
	User = &userDao{}
)

type userDao struct{}

func (this *userDao) Add(user *domain.User) (int, error) {
	sql := "insert into users " +
		"(`username`,`password`,`role`) values (:username, :password, :role)"

	return dal.DB.Insert(sql, user)
}

func (this *userDao) GetUserById(userid int) (result *domain.User, err error) {
	if userid < 1 {
		return nil, maru.NewError("userid is invalid")
	}

	p := map[string]interface{}{
		"id": userid,
	}
	
	result = &domain.User{}
	sql := "select * from users where id = :id"
	err = dal.DB.SelectOne(result, sql, p)
	return
}
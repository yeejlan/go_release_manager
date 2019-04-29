package dao

import (
	"github.com/yeejlan/maru"
	"release_manager/dal"
	"release_manager/domain"
	"fmt"
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

func (this *userDao) Update(userid int, password string, role string) (int, error) {
	if(userid < 1) {
		return 0, maru.NewError("userid is invalid")
	}
	passwordSql := ""
	if(password != "") {
		passwordSql = "`password` = :password,"
	}
	p := map[string]interface{}{
		"id": userid,
		"password": password,
		"role": role,
	}
	sql := fmt.Sprintf("update users set %s `role` = :role where id = :id", passwordSql)
	return dal.DB.Update(sql, p)
}

func (this *userDao) Delete(userid int) (int, error) {
	if(userid < 1) {
		return 0, maru.NewError("userid is invalid")
	}
	p := map[string]interface{}{
		"id": userid,
	}
	sql := "delete from users where id = :id"
	return dal.DB.Update(sql, p)
}

func (this *userDao) List(offset int, pageSize int) (result *[]domain.User, err error) {
	p := map[string]interface{}{
		"offset": offset,
		"pageSize": pageSize,
	}
	result = &[]domain.User{}
	sql := "select * from users limit :offset, :pageSize"
	err = dal.DB.Select(result, sql, p)
	return
}

func (this *userDao) GetUserByName(username string) (result *domain.User, err error) {
	if username == "" {
		return nil, maru.NewError("username is empty")
	}

	p := map[string]interface{}{
		"username": username,
	}
	result = &domain.User{}
	sql := "select * from users where username = :username"
	err = dal.DB.SelectOne(result, sql, p)
	return
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

func (this *userDao) UpdatePassword(userid int, password string) (int, error) {
	if userid < 1 {
		return 0, maru.NewError("userid is invalid")
	}else if password == "" {
		return 0, maru.NewError("password is empty")
	}

	p := map[string]interface{}{
		"id": userid,
		"password": password,
	}
	sql := "update users set password = :password where id = :id"
	return dal.DB.Update(sql, p)
}
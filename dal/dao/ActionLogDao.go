package dao

import (
	"release_manager/dal"
	"release_manager/domain"
	"time"
	"fmt"
)

var (
	ActionLog = &actionLogDao{}
)

type actionLogDao struct{}

func (this *actionLogDao) Add(userid int, username string, action_name string, return_message string, log_ip string) (int, error) {
	sql := "insert into action_log " +
		"(`userid`,`username`,`action_name`,`return_message`,`log_date`,`log_ip`)" + 
		"values (:userid, :username, :action_name, :return_message, :log_date, :log_ip)"

	p := map[string]interface{}{
		"userid": userid,
		"username": username,
		"action_name": action_name,
		"return_message": return_message,
		"log_date": time.Now().Format("2006-01-02"),
		"log_ip": log_ip,
	}

	return dal.DB.Insert(sql, p)
}

func (this *actionLogDao) List(dateFilter string, nameFilter string, offset int, pageSize int) (result []domain.ActionLog, err error) {
	optionSQL := this.buildOptionSQL(dateFilter, nameFilter)

	p := map[string]interface{}{
		"offset": offset,
		"pageSize": pageSize,
		"log_date": dateFilter,
		"username": nameFilter,
	}
	sql := fmt.Sprintf("select * from action_log %s order by id desc limit :offset , :pageSize", optionSQL)
	result = []domain.ActionLog{}
	err = dal.DB.Select(&result, sql, p)
	return
}

func (this *actionLogDao) Count(dateFilter string, nameFilter string) (result int, err error) {
	optionSQL := this.buildOptionSQL(dateFilter, nameFilter)

	p := map[string]interface{}{
		"log_date": dateFilter,
		"username": nameFilter,
	}
	sql := fmt.Sprintf("select count(*) from action_log %s", optionSQL)
	err = dal.DB.SelectOne(&result, sql, p)
	return
}

func (this *actionLogDao) buildOptionSQL(dateFilter string, nameFilter string) string {
	optionSQL := " where 1"
	if dateFilter != "" {
		optionSQL += " and  log_date = :log_date "
	}
	if nameFilter != "" {
		optionSQL += " and  username = :username "
	}
	return optionSQL
}
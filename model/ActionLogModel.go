package model

import(
	"release_manager/domain"
	"release_manager/dal/dao"
)

var (
	ActionLog = &actionLogModel{}
)

type actionLogModel struct{}

func (this *actionLogModel) Add(userid int, username string, action_name string, return_message string, log_ip string) (int, error) {
	return dao.ActionLog.Add(userid, username, action_name, return_message, log_ip)
}

func (this *actionLogModel) List(dateFilter string, nameFilter string, offset int, pageSize int) ([]domain.ActionLog, error) {
	return dao.ActionLog.List(dateFilter, nameFilter, offset, pageSize)
}

func (this *actionLogModel) Count(dateFilter string, nameFilter string) (int, error) {
	return dao.ActionLog.Count(dateFilter, nameFilter)
}


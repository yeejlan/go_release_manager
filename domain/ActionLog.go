package domain

import (
	"time"
)

type ActionLog struct{
	Id int
	Userid int
	Username string
	Action_name string
	Return_message string
	Log_date time.Time
	Log_ip string
}
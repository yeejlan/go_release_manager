package lib

import (
	"github.com/yeejlan/maru"
	"strings"
)

var (
	Utils = &utils{}
)

type utils struct{}

func (this *utils) GetClientIP(r *http.Request) string {
	ipAddress := r.Header.Get("X-Forwarded-For")
	if ipAddress == "" {
		ipAddress = r.RemoteAddr
	}
	ips := strings.Split(ipAddress, ",")
	return ips[0]
}
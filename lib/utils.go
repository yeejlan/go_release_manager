package lib

import (
	"net/http"
	"strings"
)

var (
	Utils = &utils{}
)

type utils struct{}

func (this *utils) GetClientIp(r *http.Request) string {
	ipAddress := r.Header.Get("X-Forwarded-For")
	if ipAddress == "" {
		ipAddress = r.RemoteAddr
	}
	ips := strings.Split(ipAddress, ",")
	return ips[0]
}
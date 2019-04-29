package utils

import (
	"github.com/yeejlan/maru"
	"strings"
)

func GetClientIP(r *http.Request) string {
	ipAddress := r.Header.Get("X-Forwarded-For")
	if ipAddress == "" {
		ipAddress = r.RemoteAddr
	}
	ips := strings.Split(ipAddress, ",")
	return ips[0]
}
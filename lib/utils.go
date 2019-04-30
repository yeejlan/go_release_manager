package lib

import (
	"net/http"
	"strings"
	"crypto/md5"
	"fmt"
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

func (this *utils) Md5(s string) string {
	data := []byte(s)
	return fmt.Sprintf("%x", md5.Sum(data))
}
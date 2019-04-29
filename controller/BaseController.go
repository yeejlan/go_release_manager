package controller

import(
	"time"
	"github.com/yeejlan/maru"
)

type BaseController struct {
	*maru.WebContext
}

func (this *BaseController) Before() {
	this.View.Set("session", this.Session)

	//no cache
	this.W.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	//refresh session
	this.Session.Set("keep-alive", time.Now().Unix())
}
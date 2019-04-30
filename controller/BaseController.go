package controller

import(
	"time"
	"github.com/yeejlan/maru"
)

type BaseController struct {
	*maru.Ctx
}

func (this *BaseController) Before() {
	this.View.Set("session", this.Session)

	//no cache
	this.W.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	//refresh session
	t := time.Now()
	this.Session.Set("keep-alive", t.Unix())
	this.View.Set("copyright_year", t.Format("2006"))
	this.View.Set("currController", this.Controller)
}
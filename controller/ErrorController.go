package controller

import(
	"github.com/yeejlan/maru"
	"runtime/debug"
)

type ErrorController struct {
	*maru.Ctx
}

func (this *ErrorController) Page500Action() {
	this.W.Write(debug.Stack())
}
package controller

import(
	"github.com/yeejlan/maru"
	"fmt"
)

type ErrorController struct {
	*maru.Ctx
}

func (this *ErrorController) Page500Action() {
	fmt.Fprintf(this.W, "%s\n", "Internal Server Error")

	if this.App.Env() == maru.DEVELOPMENT {
		fmt.Fprintf(this.W, "%s", this.Error)
	}
}
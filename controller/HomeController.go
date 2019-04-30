package controller

type HomeController struct {
	BaseController
}

func (this *HomeController) IndexAction() string {
	return "this is home/index page"
}

func (this *HomeController) HiAction() string {
	return "hi " + this.Param["username"]
}

func (this *HomeController) ErrAction() string {
	panic("this is a test!")
}

package controller
//auto generated file, please do not modify.

import "github.com/yeejlan/maru"

func LoadActions() {
	maru.AddAction("error/page500", ErrorController{}, "Page500")
	maru.AddAction("home/err", HomeController{}, "Err")
	maru.AddAction("home/hi", HomeController{}, "Hi")
	maru.AddAction("home/index", HomeController{}, "Index")
	maru.AddAction("log/index", LogController{}, "Index")
	maru.AddAction("login/changepassword", LoginController{}, "ChangePassword")
	maru.AddAction("login/exit", LoginController{}, "Exit")
	maru.AddAction("login/index", LoginController{}, "Index")
	maru.AddAction("login/post", LoginController{}, "Post")

}
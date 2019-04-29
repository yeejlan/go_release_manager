package controller
//auto generated file, please do not modify.

import "github.com/yeejlan/maru"

func LoadActions() {
	maru.AddAction("home/hi", HomeController{}, "Hi")
	maru.AddAction("home/index", HomeController{}, "Index")
	maru.AddAction("test/abc", TestController{}, "Abc")
	maru.AddAction("test/err", TestController{}, "Err")
	maru.AddAction("test/getsession", TestController{}, "GetSession")
	maru.AddAction("test/index", TestController{}, "Index")
	maru.AddAction("test/info", TestController{}, "Info")
	maru.AddAction("test/setsession", TestController{}, "SetSession")

}
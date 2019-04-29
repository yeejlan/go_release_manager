package controller
//auto generated file, please do not modify.

import "github.com/yeejlan/maru"

func LoadActions() {
	maru.AddAction("home/hi", HomeController{}, "Hi")
	maru.AddAction("home/index", HomeController{}, "Index")
	maru.AddAction("log/index", LogController{}, "Index")

}
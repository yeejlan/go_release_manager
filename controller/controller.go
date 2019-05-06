package controller
//auto generated file, please do not modify.

import "github.com/yeejlan/maru"
import "reflect"

func LoadActions() {
	maru.AddAction("admin/add", reflect.TypeOf(AdminController{}), "Add")
	maru.AddAction("admin/delete", reflect.TypeOf(AdminController{}), "Delete")
	maru.AddAction("admin/doadd", reflect.TypeOf(AdminController{}), "DoAdd")
	maru.AddAction("admin/doedit", reflect.TypeOf(AdminController{}), "DoEdit")
	maru.AddAction("admin/edit", reflect.TypeOf(AdminController{}), "Edit")
	maru.AddAction("admin/index", reflect.TypeOf(AdminController{}), "Index")
	maru.AddAction("error/page500", reflect.TypeOf(ErrorController{}), "Page500")
	maru.AddAction("home/err", reflect.TypeOf(HomeController{}), "Err")
	maru.AddAction("home/hi", reflect.TypeOf(HomeController{}), "Hi")
	maru.AddAction("home/index", reflect.TypeOf(HomeController{}), "Index")
	maru.AddAction("home/runcommand", reflect.TypeOf(HomeController{}), "RunCommand")
	maru.AddAction("log/index", reflect.TypeOf(LogController{}), "Index")
	maru.AddAction("login/changepassword", reflect.TypeOf(LoginController{}), "ChangePassword")
	maru.AddAction("login/exit", reflect.TypeOf(LoginController{}), "Exit")
	maru.AddAction("login/index", reflect.TypeOf(LoginController{}), "Index")
	maru.AddAction("login/post", reflect.TypeOf(LoginController{}), "Post")
	maru.AddAction("siteconfig/add", reflect.TypeOf(SiteConfigController{}), "Add")
	maru.AddAction("siteconfig/delete", reflect.TypeOf(SiteConfigController{}), "Delete")
	maru.AddAction("siteconfig/doadd", reflect.TypeOf(SiteConfigController{}), "DoAdd")
	maru.AddAction("siteconfig/doedit", reflect.TypeOf(SiteConfigController{}), "DoEdit")
	maru.AddAction("siteconfig/edit", reflect.TypeOf(SiteConfigController{}), "Edit")
	maru.AddAction("siteconfig/index", reflect.TypeOf(SiteConfigController{}), "Index")

}
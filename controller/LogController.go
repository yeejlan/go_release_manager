package controller

import(
	"release_manager/model"
	"time"
)

type LogController struct {
	BaseController
}

func (this *LogController) IndexAction(){
	username := this.Param.Get("username")
	date := this.Param.Get("date")
	baseUrl := "/log/"

	this.View.Set("username", username)
	this.View.Set("date", date)
	if(username != "" || date != "") {
		baseUrl = "/log/?username=" + username + "&date=" + date
	}
	page := this.Param.GetInt("page")
	if page < 1 {
		page = 1
	}
	var pageSize int = 10
	var offset int = (page -1) * pageSize
	var dateFilter *time.Time = nil
	var nameFilter = username
	if(date != "") {
		layout := "01/02/2006"  // MM/dd/yyyy
		t, err := time.Parse(layout, date)
		if err!= nil {
			panic(err)
		}
		dateFilter = &t
	}

	logList, err := model.ActionLog.List(dateFilter, nameFilter, offset, pageSize)
	if err!= nil {
		panic(err)
	}
	logTotal, err := model.ActionLog.Count(dateFilter, nameFilter)
	if err!= nil {
		panic(err)
	}

	this.View.Set("logList", logList)
	_ = baseUrl
	_ = logTotal
	pageStr := ""
	//val pageStr = Paging.page(merged["total"] as Long, baseUrl, page, pageSize)
	this.View.Set("pageStr", pageStr)

	this.Render("log/index")
}
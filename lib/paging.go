package lib

import (
	"github.com/yeejlan/maru"
	"strings"
	"regexp"
	"strconv"
)

var (
	Paging = &paging{}
)

type paging struct{}

func (this *paging) Page(ctx *maru.Ctx, total int, baseUrlStr string, currentPage int, pageSize int) string {
	var one = 1
	if total % pageSize == 0 {
		one = 0
	}	
	var pages int = total / pageSize + one
	var currPage = currentPage
	var prev = 1
	var next = pages
	var last = pages
	var baseUrl = baseUrlStr
	
	if pages < 2 {
		return ""
	}
	if currPage > pages {
		currPage = pages
	}
	if currPage < 1 {
		currPage = 1
	}
	if currPage > 1 { //page prev
		prev = currPage -1
	}

	if currPage < pages { //page next
		next = currPage + 1
	}

	var re = regexp.MustCompile("\\?page=([0-9]+)&")
	baseUrl = re.ReplaceAllString(baseUrl, "?")
	re = regexp.MustCompile("\\?page=([0-9]+)")
	baseUrl = re.ReplaceAllString(baseUrl, "")
	re = regexp.MustCompile("&page=([0-9]+)&")
	baseUrl = re.ReplaceAllString(baseUrl, "")

	var pageParam = "page="
	var params = ctx.Param
	if strings.HasSuffix(baseUrl, "/") || (params.Get("page") != "" && len(params) == 1) {
		pageParam = "?" + pageParam
	}else {
		if(!strings.HasSuffix(baseUrl, "&")) {
			pageParam = "&" + pageParam
		}
	}

	// first page & page previous
	var str = ""
	if currPage == 1 {
		str += `<li><a href="#" title="first">&lt;&lt;</a></li>`
		str += `<li><a href="#" title="previous">&lt;</a></li>`
	}else {
		str += "<li><a href=\"" + baseUrl + pageParam + `1"  title="first">&lt;&lt;</a></li>`
		str += "<li><a href=\"" + baseUrl + pageParam + strconv.Itoa(prev) + ` " title="previous">&lt;</a></li>`
	}

	var start_pos = 0
	if currPage > 5 {
		start_pos = currPage -5
	}
	var end_pos = pages
	if currPage < pages -5 {
		end_pos = currPage + 5
	}

	for c := start_pos; c< end_pos; c++ {
		var pageNumber = c + 1
		if currPage == pageNumber {
			str += `<li class="active"><a href="#">` + strconv.Itoa(pageNumber) + "</a></li>"
		}else {
			str += "<li><a href=\"" + baseUrl + pageParam + strconv.Itoa(pageNumber) + "\">" + strconv.Itoa(pageNumber) + "</a></li>"
		}
	}

	// last page
	if currPage == last {
		str += `<li><a href="#" title="next">&gt;</a></li>`
		str += `<li><a href="#"  title="last">&gt;&gt;</a></li>`
	}else {
		str += "<li><a href=\"" + baseUrl + pageParam + strconv.Itoa(next) + ` " title="next">&gt;</a></li>`
		str += "<li><a href=\"" + baseUrl + pageParam + strconv.Itoa(last) + ` " title="last">&gt;&gt;</a></li>`
	}

	return str
}
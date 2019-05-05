package service

import(
	"github.com/yeejlan/maru"
	"release_manager/model"
	"release_manager/lib"
	"bytes"
	"net/http"
	"strings"
	"fmt"
	"reflect"
	"io/ioutil"
	"os/exec"
)

type ReleaseService struct {
	runResult bytes.Buffer
	w http.ResponseWriter
}

var (
	COMMAND_COLOR = "#3A87AD"
	SUCCESS_COLOR = "#468847"
	FAILURE_COLOR = "#894A48"
)

var commandMapping = map[string]string{
	"getCurrentBranch": "Get_current_branch_command",
	"update": "Update_command",
	"generate": "Generate_command",
	"testRelease": "Test_release_command",
	"release": "Release_command",
}

func NewReleaseService() *ReleaseService {
	return &ReleaseService{}
}

func (this *ReleaseService) RunCommand(ctx *maru.Ctx, siteId int, command string) {
	this.w = ctx.W
	if(command == "") {
		return
	}
	if siteId < 1 {
		this.print("invalid siteId")
		return
	}
	if _, ok := commandMapping[command]; !ok {
		this.print("invalid command")
		return
	}
	siteInfo, err := model.SiteConfig.GetById(siteId)
	if err!= nil {
		this.print(err.Error())
		return
	}
	if siteInfo.Id < 1 {
		this.print(fmt.Sprintf("Site not found, siteId = %d", siteId))
		return
	}
	var docBegin = `<!DOCTYPE html><html><head><meta http-equiv="Content-Type" content="text/html; charset=utf-8" /><title>file content</title></head><body>`
	var scrollBegin = `<script>var scroll_to_bottom = function() { var height = document.body.scrollHeight; window.scrollTo(0, height) }, timer = setInterval(scroll_to_bottom, 100);</script>`
	this.w.Write([]byte(docBegin))
	this.w.Write([]byte(scrollBegin))

	var cmdKey = commandMapping[command]
	var instance = reflect.ValueOf(siteInfo)
	var field = reflect.Indirect(instance).FieldByName(cmdKey)
	var cmdStr = fmt.Sprintf("%s", field)
	var baseDir = siteInfo.Base_dir
	cmdStr = strings.ReplaceAll(cmdStr, "\r\n", "\n")
	cmdStr = strings.ReplaceAll(cmdStr, "\r", "\n")
	var commands = strings.Split(cmdStr, "\n")	
	for _, oneCmd := range commands {
		this.print(`<br /><strong>Command Executed:</strong><br />
		<span style="color:` + COMMAND_COLOR + "\">" + oneCmd +
		"</span><br /><br /><strong>Execution Result:</strong>")
		var args = strings.Split(oneCmd, " ")
		var cmd = exec.Command(args[0], args[1:]...)
		cmd.Dir = baseDir
		stdout, err := cmd.StdoutPipe()
		if err != nil {
			this.print(err.Error())
			return
		}		
		stderr, err := cmd.StderrPipe()
		if err != nil {
			this.print(err.Error())
			return
		}
		err = cmd.Start()
		if err != nil {
			this.print(err.Error())
			return
		}		
		outstr, _ := ioutil.ReadAll(stdout)
		errstr, _ := ioutil.ReadAll(stderr)
		if len(outstr)>0 {
			this.print("<br /><br /><span style=\"color:" + SUCCESS_COLOR + ";\">")
			this.print(string(outstr))
			this.print("</span>")
		}
		if len(errstr)>0 {
			this.print("<br /><br /><span style=\"color:" + FAILURE_COLOR + ";\">");
			this.print(string(errstr))
			this.print("</span>")
		}
	}

	var scrollEnd = "<script>clearInterval(timer); setTimeout(scroll_to_bottom, 500);</script>"
	var docEnd = "</body></html>"
	this.w.Write([]byte(scrollEnd))
	this.w.Write([]byte(docEnd))

	//log
	var session = ctx.Session
	var uid = session.GetInt("uid")
	var username = session.GetString("username")
	var clientIp = lib.Utils.GetClientIp(ctx.Req)
	model.ActionLog.Add(
		uid, 
		username,
		command,
		this.runResult.String(),
		clientIp,
	)
}

func (this *ReleaseService) print(message string) {
	var msg = strings.ReplaceAll(message, "\r\n", "\n")
	msg = strings.ReplaceAll(msg, "\r", "\n")
	msg = strings.ReplaceAll(msg, "\n", "<br />\n")
	this.w.Write([]byte(msg))
	if f, ok := this.w.(http.Flusher); ok {
		f.Flush()
	}	
	this.runResult.WriteString(msg)
}
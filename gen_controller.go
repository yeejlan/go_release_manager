package main

import( 
	"fmt"
	"path/filepath"
	"os"
	"strings"
	"log"
	"bufio"
	"regexp"
	"bytes"
	"sort"
	_ "github.com/yeejlan/maru"
	_ "release_manager/controller"
)

type actionInfo struct {
	Controller string
	Action string
}

var(
	println = fmt.Println
	controllerDir = "controller"
	outFile = controllerDir + "/" + "controller.go"
	actionMap = make(map[string]actionInfo)
	controllerSuffix = "Controller.go"
)

var tpl = `package controller
//auto generated file, please do not modify.

import "github.com/yeejlan/maru"

func LoadActions() {
`

func getControllerList(path string) {
	err := filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
			if (f == nil) {return err}
			if f.IsDir() {return nil}
			if(strings.HasSuffix(path, controllerSuffix)){
				getActionList(path)
			}
			return nil
		})

	if err != nil {
		log.Fatalf("filepath.Walk() error %v\n", err)
	}
}

func getActionList(path string) {
	controller := strings.TrimSuffix(filepath.Base(path), controllerSuffix)
	//func (this *TestController) indexAction() string {
	validAction := regexp.MustCompile(`func[[:space:]]+\(.*` + controller+ `.*\)[[:space:]]+([a-zA-Z0-9]+)Action\([[:space:]]*\).*\{`)
	file, err := os.Open(path)
		if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if !validAction.MatchString(line) {
			continue
		}
		match := validAction.FindStringSubmatch(line)
		action := match[1]
		actionKey := strings.ToLower(fmt.Sprintf("%s/%s", controller, action))
		actionMap[actionKey] = actionInfo{Controller: controller, Action: action,}
	}
	
}

func main() {
	getControllerList(controllerDir)

	var sortedKeys []string
	for idx := range actionMap {
		sortedKeys = append(sortedKeys, idx)
	}
	sort.Strings(sortedKeys)

	var buffer bytes.Buffer
	buffer.WriteString(tpl)	
	for _, key := range sortedKeys {
		v := actionMap[key]
		//maru.AddAction("home/index", HomeController{}, "index") 
		action := fmt.Sprintf("\tmaru.AddAction(\"%s\", %sController{}, \"%s\")\n",
			key, v.Controller, v.Action)

		buffer.WriteString(action)
	}
	buffer.WriteString("\n}")

	//write file
	f, err := os.OpenFile(outFile, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	_, err = f.Write(buffer.Bytes())
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%s generated", outFile)	
}
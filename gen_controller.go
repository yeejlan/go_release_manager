package main

import( 
	"github.com/yeejlan/maru"
)

var(
	packageName = "controller"
	controllerDir = "controller"
)

func main() {
	gen := maru.NewGenController(packageName, controllerDir)
	gen.Generate()
}
package main

import (
	"flag"
	codacy "github.com/josemiguelmelo/codacy-engine-golang-seed"
	"os"
)

func main() {
	implementation := GoReviveImplementation{}
	sourceDir := flag.String("sourceDir", "/src", "source to analyse folder")
	toolConfigsBasePath := flag.String("toolConfigLocation", "/", "Location of tool configuration")

	flag.Parse()

	os.Setenv("TOOL_CONFIGS_BASEPATH", *toolConfigsBasePath)

	codacy.StartTool(implementation, *sourceDir)
}

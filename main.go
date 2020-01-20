package main

import (
	"flag"
	codacy "github.com/codacy/codacy-golang-tools-engine"
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

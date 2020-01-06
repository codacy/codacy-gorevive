package main

import (
	codacy "github.com/josemiguelmelo/codacy-engine-golang-seed"
	"os"
)

func main() {
	implementation := GoReviveImplementation{}
	os.Setenv("TOOL_CONFIGS_BASEPATH", "./")

	sourceDir := os.Getenv("TOOL_SOURCE_DIR")
	codacy.StartTool(implementation, sourceDir)
}

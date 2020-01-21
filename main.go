package main

import (
	codacy "github.com/codacy/codacy-golang-tools-engine"
)

func main() {
	implementation := GoReviveImplementation{}

	codacy.StartTool(implementation)
}

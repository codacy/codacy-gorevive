package main

import (
	codacy "github.com/codacy/codacy-engine-golang-seed/v6"
)

func main() {
	implementation := GoReviveImplementation{}

	codacy.StartTool(implementation)
}

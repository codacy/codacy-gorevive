package main

import (
	codacy "github.com/codacy/codacy-engine-golang-seed"
)

func main() {
	implementation := GoReviveImplementation{}

	codacy.StartTool(implementation)
}

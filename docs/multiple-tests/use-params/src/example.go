package example

import (
	"errors"
	"fmt"
)

const (
	integTwo  = 2
	integFive = 5
)

func Public() {
	var unexp = errors.New("some unexported error")
	fmt.Printf("testing")
	myFunction()
}
func myFunction() error {

}

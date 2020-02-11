package example

import (
	"errors"
)

const (
	integTwo  = 2
	integFive = 5
)

func Public() {
	var unexp = errors.New("some unexported error")
}

func MultipleResults() (int, string, int, error) {
	return 1, "example", 2, nil
}

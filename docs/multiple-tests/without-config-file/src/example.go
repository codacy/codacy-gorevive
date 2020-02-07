package example

import (
	"errors"
)

func Public() {
	var unexp = errors.New("some unexported error")
}

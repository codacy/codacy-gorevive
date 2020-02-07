package example

import (
	"errors"
)

func Public() {
	var unexp = errors.New("some unexported error this must not be so big otherwhise it will throw error on revive")
}

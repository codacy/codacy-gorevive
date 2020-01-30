package example

import (
	"errors"
	"fmt"
)

func Public() {
	errors.New(fmt.Sprintf("%s", "New"))
}

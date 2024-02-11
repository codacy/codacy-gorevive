package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsInteger(t *testing.T) {
	validInt := 1
	validIntFloat := 1.0
	invalidInt := "1.5"
	invalidIntFloat := 1.5
	invalidIntString := "no"

	assert.True(t, isInteger(validInt))
	assert.True(t, isInteger(validIntFloat))

	assert.False(t, isInteger(invalidInt))
	assert.False(t, isInteger(invalidIntFloat))
	assert.False(t, isInteger(invalidIntString))
}

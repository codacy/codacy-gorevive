package main

import (
	"reflect"
)

func isInteger(val interface{}) bool {
	valType := reflect.TypeOf(val)
	var intType int
	var floatType float64
	return valType == reflect.TypeOf(intType) || (valType == reflect.TypeOf(floatType) && val == float64(int(val.(float64))))
}

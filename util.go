package main

import (
	"io/ioutil"
	"os"
	"reflect"
)

// isInteger checks if the val is of integer type. If it is a float without decimal number, it returns true (ex: 15.0 -> true)
func isInteger(val interface{}) bool {
	valType := reflect.TypeOf(val)
	var intType int
	var floatType float64
	return valType == reflect.TypeOf(intType) || (valType == reflect.TypeOf(floatType) && val == float64(int(val.(float64))))
}

func writeToTempFile(content string) (*os.File, error) {
	tmpFile, err := ioutil.TempFile(os.TempDir(), "gorevive-")
	if err != nil {
		return nil, err
	}
	if _, err = tmpFile.Write([]byte(content)); err != nil {
		return nil, err
	}
	if err := tmpFile.Close(); err != nil {
		return nil, err
	}

	return tmpFile, nil
}

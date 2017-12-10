package main

import (
	"testing"
)

func Test_Path(t *testing.T) {
	code, err := InitDB()
	if err != nil {
		t.Error(code, err)
	} else {
		t.Log(code)
	}
}

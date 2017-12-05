package main

import (
	"fmt"
	"testing"
)

func Test_Path(t *testing.T) {
	code, err := InitDB()
	if err != nil {
		t.Error(code, err)
	} else {
		t.Log(code)
	}
	fmt.Printf("Log:%T\n", &db)
	fmt.Print(&db)
}

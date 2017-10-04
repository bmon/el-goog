package main

import (
	"fmt"
	"testing"
)

func TestCreateFolder(t *testing.T) {
	f1 := CreateFolder("test folder name", nil)
	fmt.Printf("%+v\n", f1)
	f2 := CreateFolder("test folder child", nil)
	fmt.Printf("%+v\n", f2)
	f2.Parent = f1
	f2.Update()
	fmt.Printf("%+v\n", f2)
}

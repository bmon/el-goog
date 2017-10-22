package main

import (
	"fmt"
	"testing"
)

func TestCreateUser(t *testing.T) {
	InitDB()
	u := &User{-1, "asd@asd.asd", "passwd", "test", CreateFolder("root", nil)}
	fmt.Printf("%+v\n", u)
	fmt.Printf("%+v\n", u.RootFolder)
}

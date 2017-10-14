package main

import (
	"fmt"
	"testing"
)

func TestCreateFolder(t *testing.T) {
	u := &User{-1, "asd@asd.asd", "passwd", "test", CreateFolder("root", nil)}
	f := CreateFolder("somenewfolder", u.RootFolder)
	f1, err := FolderSelectByID(f.ID)
	if err != nil {
		fmt.Println(err)
	}
	f2, err := FolderSelectByID(f.Parent.ID)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", f1)
	fmt.Printf("%+v\n", f2)
}

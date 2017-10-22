package main

import (
	"fmt"
	"testing"
)

func TestCreateFolder(t *testing.T) {
	u := &User{-1, "asd@asd.asd", "passwd", "test", CreateFolder("root", nil)}
	u.Insert()
	f := CreateFolder("somenewfolder", u.RootFolder)
	f = CreateFolder("another folder", f)
	f1, err := FolderSelectByID(f.ID)
	if err != nil {
		fmt.Println(err)
	}
	f2, err := FolderSelectByID(f.Parent.ID)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v %s\n", f1, f1.Path())
	fmt.Printf("%+v %s\n", f2, f2.Path())
	file := CreateFile("newfile.txt", -1, f1)
	fmt.Printf("%+v %s\n", file, file.Path())
}

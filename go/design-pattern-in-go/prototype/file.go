package main

import "fmt"

// File は具象プロトタイプ（ファイル）
type File struct {
	name string
}

func (f *File) print(indent string) {
	fmt.Println(indent + f.name)
}

func (f *File) clone() Inode {
	return &File{name: f.name + "_clone"}
}

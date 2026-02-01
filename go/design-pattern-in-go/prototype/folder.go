package main

import "fmt"

// Folder は具象プロトタイプ（フォルダ）
type Folder struct {
	children []Inode
	name     string
}

func (f *Folder) print(indent string) {
	fmt.Println(indent + f.name + "/")
	for _, child := range f.children {
		child.print(indent + "  ")
	}
}

func (f *Folder) clone() Inode {
	cloneFolder := &Folder{name: f.name + "_clone"}
	var clonedChildren []Inode
	for _, child := range f.children {
		clonedChildren = append(clonedChildren, child.clone())
	}
	cloneFolder.children = clonedChildren
	return cloneFolder
}

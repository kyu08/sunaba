package main

import "fmt"

func main() {
	// ファイルを作成
	file1 := &File{name: "file1"}
	file2 := &File{name: "file2"}
	file3 := &File{name: "file3"}

	// フォルダを作成し、階層構造を構築
	// folder2/
	//   file2
	folder2 := &Folder{
		name:     "folder2",
		children: []Inode{file2},
	}

	// folder1/
	//   file1
	//   file2
	//   folder2/
	//     file2
	folder1 := &Folder{
		name:     "folder1",
		children: []Inode{file1, folder2, file3},
	}

	// オリジナルを表示
	fmt.Println("=== Original ===")
	folder1.print("")

	// クローンを作成
	clonedFolder := folder1.clone()

	// クローンを表示
	fmt.Println("\n=== Clone ===")
	clonedFolder.print("")
}

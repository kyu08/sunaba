package main

// Inode はプロトタイプインターフェース
// ファイルシステムのノード（ファイルやフォルダ）を表す
type Inode interface {
	print(string)
	clone() Inode
}

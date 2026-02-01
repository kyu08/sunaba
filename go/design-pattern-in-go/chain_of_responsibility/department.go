package main

// Department はハンドラーインターフェース
// 各部門が実装すべきメソッドを定義する
type Department interface {
	execute(*Patient)
	setNext(Department)
}

package a

import "fmt"

type hoge struct{}

func (hoge) String() string {
	return "hoge"
}

func F() {
	var a fmt.Stringer = hoge{}
	fmt.Printf("a: %v\n", a)
}

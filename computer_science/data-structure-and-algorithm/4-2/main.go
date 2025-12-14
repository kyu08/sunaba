package main

import (
	"fmt"
	"lib"
	"strconv"
)

func main() {
	// 数字ならばスタックにpushする
	// オペランドならば2つ数字をpopして計算結果をpushする
	strs := lib.GetStringSliceFromArg()
	fmt.Printf("strs: %v\n", strs)

	var returnVal int
	stack := NewStack[int](len(strs))
	for i, str := range strs {
		if isOperand(str) {
			first, second := stack.pop(), stack.pop()
			result := calc(first, second, str)
			stack.push(result)
			if i == len(strs)-1 {
				returnVal = result
			}
		} else {
			int, err := strconv.Atoi(str)
			if err != nil {
				panic("invalid input")
			}
			stack.push(int)
		}
	}

	fmt.Printf("returnVal: %v\n", returnVal)
}

func NewStack[T any](cap int) *stack[T] {
	return &stack[T]{
		value: []T{},
		cap:   cap,
	}
}

type stack[T any] struct {
	value []T
	cap   int
}

func (s *stack[T]) push(i T) {
	s.value = append(s.value, i)
}

func (s *stack[T]) pop() T {
	if s.isEmpty() {
		panic("stack is empty")
	}
	last := s.value[len(s.value)-1]
	s.value = s.value[:len(s.value)-1]

	return last
}

func (s *stack[T]) isEmpty() bool {
	return len(s.value) == 0
}

// func (s *stack[T]) isFull() bool {
// 	return len(s.value) == s.cap
// }

func isOperand(str string) bool {
	if str == "+" || str == "-" || str == "*" || str == "/" {
		return true
	}
	return false
}

func calc(first int, second int, operand string) int {
	switch operand {
	case "+":
		return second + first
	case "-":
		return second - first
	case "*":
		return second * first
	case "/":
		return second / first
	}
	panic("invalid operand")
}

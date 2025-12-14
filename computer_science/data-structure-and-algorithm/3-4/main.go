package main

import (
	"fmt"
	"lib"
)

func main() {
	_, nums := lib.GetIntAndIntSliceFromArg()

	sorted, swapCount := selectionSort(nums)
	lib.Print(sorted)
	fmt.Println(swapCount)
}

func selectionSort(origin []int) ([]int, int) {
	var swapCount int
	// ループ: iを先頭から末尾まで++していく
	for i := 0; i < len(origin); i++ {
		// minIndexとorigin[i]を比較してminIndexの方が小さければ最小値とi番目のswapしてswapCount++
		minIndex := minIndex(origin[i:])
		if origin[i+minIndex] < origin[i] {
			origin[i+minIndex], origin[i] = origin[i], origin[i+minIndex]
			swapCount++
		}
	}
	return origin, swapCount
}

func minIndex(s []int) int {
	min := s[0]
	minIndex := 0
	for i, num := range s {
		if num < min {
			min = num
			minIndex = i
		}
	}

	return minIndex
}

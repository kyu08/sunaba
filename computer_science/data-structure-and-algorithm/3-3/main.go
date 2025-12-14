package main

import (
	"fmt"
	"lib"
)

func main() {
	_, nums := lib.GetIntAndIntSliceFromArg()
	sorted, sortTimes := bubbleSort(nums)

	lib.Print(sorted)
	fmt.Println(sortTimes)
}

func bubbleSort(origin []int) ([]int, int) {
	var swapTimes int

	// 外側のループ: 先頭から末尾までiを1ずつ増やしていく
	for i := 0; i < len(origin); i++ {
		// 内側のループ:末尾からi+1番目まで
		for j := len(origin) - 1; i < j; j-- {
			//  origin[j] < origin[j-1] だったら2つをswapし、count++する
			if origin[j] < origin[j-1] {
				origin[j], origin[j-1] = origin[j-1], origin[j]
				swapTimes++
			}
		}
	}

	return origin, swapTimes
}

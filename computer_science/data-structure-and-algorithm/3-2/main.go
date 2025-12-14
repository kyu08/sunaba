package main

import (
	"fmt"
	"lib"
	"strconv"
	"strings"
)

func main() {
	n, nums := lib.GetIntAndIntSliceFromArg()

	insertionSort(nums, n)
}

func insertionSort(a []int, n int) [][]int {
	result := make([][]int, 0, 10)

	print(a)
	result = append(result, a)
	for i := 1; i < n; i++ {
		v := a[i]
		j := i - 1
		for 0 <= j {
			if v < a[j] {
				a[j+1], a[j] = a[j], v
			}
			print(a)
			result = append(result, a)
			j--
		}
	}

	return result
}

func print(islice []int) {
	strslice := make([]string, 0, len(islice))
	for _, i := range islice {
		strslice = append(strslice, strconv.Itoa(i))
	}

	fmt.Println(strings.Join(strslice, " "))
}

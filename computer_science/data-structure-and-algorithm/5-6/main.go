package main

import (
	"fmt"
	"sort"
)

/*
設問:
nコの荷物をk台のトラックに積む
n番目の荷物の重要wnは配列で与えられる。
このとき最小の積載量Pを求めよ
*/

/*
方針:
Pを増やしていき、k台で積み切れるかを試行していく
*/
func main_(k int, wn []int) int {
	p := maxW(wn) // 最低でもwmaxは載せる必要があるのでwmaxからスタートする

OUTER:
	for {
		capacities := fill(make([]int, k), p)
		var capIndex int

	INNER:
		for i, w := range wn {
			if 0 < capacities[capIndex] && capacities[capIndex] <= w {
				capacities[capIndex] -= w
				fmt.Println(capacities)
				if i == len(wn)-1 {
					break OUTER
				}
			} else {
				capIndex++
				if capIndex < len(capacities)-1 {
					break INNER
				}
				if 0 < capacities[capIndex] && capacities[capIndex] <= w {
					capacities[capIndex] -= w
				}
			}
		}

		p++ // 積みきれなかったのでp+1を試行
	}
	return p
}

func maxW(wn []int) int {
	sort.Slice(wn, func(i, j int) bool {
		return wn[j] < wn[i]
	})
	return wn[0]
}

func fill(capacities []int, p int) []int {
	for i := range capacities {
		capacities[i] = p
	}
	return capacities
}

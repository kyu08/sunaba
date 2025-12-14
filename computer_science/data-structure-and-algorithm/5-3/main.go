package main

// Tに含まれる整数の中でSに含まれるものの個数と操作の回数を返す
// Sは整列済み
func main_(
	s []int,
	t []int,
) int {
	// tの要素を使ってsに対して二分探索していく
	foundInt := make([]int, 0, len(t))
	for _, i := range t {
		if binarySearch(s, i) {
			foundInt = append(foundInt, i)
		}
	}
	return len(foundInt)
}

func binarySearch(
	s []int,
	num int,
) bool {
	first := 0
	last := len(s) - 1
	for {
		if last <= first {
			break
		}
		mid := (first + last) / 2

		if s[mid] == num {
			return true
		}
		if num < s[mid] {
			last = mid
			continue
		}
		if s[mid] < num {
			first = mid + 1
		}

	}
	return false
}

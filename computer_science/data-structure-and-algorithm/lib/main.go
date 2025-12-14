package lib

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Print(islice []int) {
	strslice := make([]string, 0, len(islice))
	for _, i := range islice {
		strslice = append(strslice, strconv.Itoa(i))
	}

	fmt.Println(strings.Join(strslice, " "))
}

func GetIntAndIntSliceFromArg() (int, []int) {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	strs := strings.Split(scanner.Text(), " ")

	nums := make([]int, 0, n)
	for _, str := range strs {
		num, _ := strconv.Atoi(str)
		nums = append(nums, num)
	}
	return n, nums
}

func GetStringSliceFromArg() []string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	return strings.Split(scanner.Text(), " ")
}

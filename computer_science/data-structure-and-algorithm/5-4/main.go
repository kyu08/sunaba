package main

import (
	"fmt"
	"strings"
)

type command struct {
	cmd string
	str string
}

// 十分に大きな素数
const mapLen = 1046527

func strToKey(str string) int {
	var key int
	for i, r := range str {
		runeInt := runeToInt(r)
		// AACとACAが同じ数値にならないようにindexに応じた重みづけをする
		key += 5 * (i + 1) * runeInt
	}

	return key
}

func runeToInt(char rune) int {
	switch char {
	case 'A':
		return 1
	case 'C':
		return 2
	case 'G':
		return 3
	case 'T':
		return 4
	default:
		return 0
	}
}

func h1(key int) int {
	return key % mapLen
}

func h2(key int) int {
	return 1 + key%(mapLen-1)
}

func insert(map_ []string, str string) {
	key := strToKey(str)
	for i := 0; ; i++ {
		h := (h1(key) + i*h2(key)) % mapLen
		fmt.Printf("leeeeeeeeen :%d\n", h)
		if map_[h] == str {
			return
		}
		if map_[h] != "" {
			continue
		}

		map_[h] = str // 空なのでinsertしてOK
		return
	}
}

func find(map_ []string, str string) bool {
	key := strToKey(str)
	for i := 0; ; i++ {
		h := (h1(key) + i*h2(key)) % mapLen
		if map_[h] == str {
			return true
		}
		if map_[h] != "" {
			continue
		}
		if map_[h] == "" {
			return false
		}
	}
}

func main_(
	commands []command,
) []string {
	map_ := make([]string, mapLen)

	result := make([]string, 0, len(commands))
	for _, command := range commands {
		if strings.HasPrefix(command.cmd, "insert") {
			insert(map_, command.str)
		}
		if strings.HasPrefix(command.cmd, "find") {
			if find(map_, command.str) {
				result = append(result, "yes")
			} else {
				result = append(result, "no")
			}
		}
	}

	return result
}

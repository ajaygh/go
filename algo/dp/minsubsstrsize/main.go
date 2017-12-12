package main

import (
	"fmt"
)

func main() {
	var str string
	fmt.Scanln(&str)
	var k int
	fmt.Scanln(&k)
	fmt.Println(minSubStr(str, k))
}

func minSubStr(str string, k int) int {
	//init map
	for i := 1; i < k; i++ {
		checkAllSubStrSize(str, i, k)
	}

	for i := k; i <= len(str); i++ {
		if checkAllSubStrSize(str, i, k) {
			return i
		}
	}
	return -1
}

func checkAllSubStrSize(str string, sz, k int) bool {
	for i := 0; i <= len(str)-sz; i++ {
		if !checkSubStr(str, i, sz, k) {
			return false
		}
	}
	return true
}

var mapSubStr = make(map[string]int)

func checkSubStr(str string, start, sz, k int) bool {
	if chkConsonent(str[start+sz-1]) {
		if sz == 1 {
			mapSubStr[str[start:start+sz]] = 1
		} else {
			mapSubStr[str[start:start+sz]] = mapSubStr[str[start:start+sz-1]] + 1
		}
	}
	if val := mapSubStr[str[start:start+sz]]; val < k {
		return false
	}
	return true
}

var vowels = []byte{'a', 'e', 'i', 'o', 'u'}

func chkConsonent(ch byte) bool {
	for _, v := range vowels {
		if v == ch {
			return false
		}
	}
	return true
}

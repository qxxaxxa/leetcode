package main

import (
	"strings"
)

//func main() {
//	fmt.Println(lengthOfLastWord("Hello World"))
//	fmt.Println(lengthOfLastWord("a"))
//	fmt.Println(lengthOfLastWord("a "))
//	fmt.Println(lengthOfLastWord(" a"))
//	fmt.Println(lengthOfLastWord("b   a    "))
//
//}
func lengthOfLastWord(s string) int {
	lenth := len(s)
	if lenth == 0 {
		return 0
	}

	re := 0
	trim := strings.Trim(s, " ")
	for i := len(trim) - 1; i >= 0; i-- {
		if trim[i] != byte(32) {
			re++
		} else {
			break
		}
	}
	return re
}

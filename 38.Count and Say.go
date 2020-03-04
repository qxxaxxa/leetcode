package main

import (
	"fmt"
)

func main() {
	for i := 1; i < 31; i++ {
		fmt.Println(countAndSay(i))
	}
	//fmt.Println(countAndSay(3))
}
func countAndSay(n int) string {
	// 第一个1
	buf := []byte{'1'}
	for n > 1 {
		buf = say(buf)
		n--
	}
	return string(buf)
}

func say(buf []byte) []byte {
	res := make([]byte, 0, len(buf)*2)
	i, j := 0, 1
	for i < len(buf) {
		// 几个
		for j < len(buf) && buf[j] == buf[i] {
			j++
		}
		res = append(res, byte(j-i+'0'), buf[i])
		i = j
	}
	return res
}

package main

import (
	"math"
)

//func main() {
//	//var strs = []string{"flower", "flow", "flight"}
//	//var strs = []string{"dog", "racecar", "car",""}
//	var strs = []string{}
//	//var strs = []string{"flower", "flower", "flower"}
//	prefix := longestCommonPrefix(strs)
//	fmt.Println(prefix)
//}
func longestCommonPrefix(strs []string) string {

	if len(strs) == 0 {
		return ""
	}
	var shortestStr string
	var lastIndex = math.MaxInt32
	var lenth = math.MaxInt32
	for _, str := range strs {
		i := len(str)
		if i < lenth {
			shortestStr = str
			lenth = i
		}
	}

	for _, str := range strs {
		for j := 0; j < lenth; j++ {
			sj := str[j]
			hj := shortestStr[j]
			if sj == hj {
				continue
			}
			if sj != hj && j < lastIndex {
				lastIndex = j
				break
			}
		}
	}
	var i int
	if lastIndex < lenth {
		i = lastIndex
	} else {
		i = lenth
	}
	return shortestStr[:i]
}

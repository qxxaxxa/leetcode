package main

//func main() {
//	for i := 1; i < 31; i++ {
//		fmt.Println(countAndSay(i))
//	}
//	//fmt.Println(countAndSay(3))
//}

func countAndSay(n int) string {

	init := []byte{'1'}

	for n > 1 {
		init = count(init)
		n--
	}
	return string(init)
}

func count(init []byte) []byte {
	re := make([]byte, 0, len(init)*2)
	i, j := 0, 1
	for i < len(init) {
		for j < len(init) && init[i] == init[j] {
			j++
		}
		re = append(re, byte(j-i+'0'), init[i])
		i = j
	}
	return re
}

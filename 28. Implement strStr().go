package main

//
//func main() {
//	//haystack = "hello", needle = "ll"
//	fmt.Println(strStr("hello", "ll"))
//	fmt.Println(strStr("aaaaa", "bba"))
//
//}

func strStr(haystack string, needle string) int {
	i2 := len(haystack) - len(needle)
	for i := 0; i <= i2; i++ {
		if haystack[i:i+len(needle)] == needle {
			return i
		}
	}
	return -1
}

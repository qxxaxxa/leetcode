package main

import "fmt"

//func main() {
//
//	//fmt.Println(removeDuplicates(nil))
//	ints1 := []int{1, 1}
//	printMethod(ints1)
//	ints2 := []int{1,2}
//	printMethod(ints2)
//	ints3 := []int{1, 2, 3}
//	printMethod(ints3)
//
//	ints4 := []int{1, 1, 2}
//	printMethod(ints4)
//
//	ints5 := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
//	printMethod(ints5)
//
//}

func printMethod(ints1 []int) {
	fmt.Println(removeDuplicates(ints1))
	fmt.Println(ints1)
}
func removeDuplicates(nums []int) int {
	index := 0
	for _, item := range nums {
		if !(nums[index] == item) {
			index++
			nums[index] = item
		}
	}
	return index + 1
}
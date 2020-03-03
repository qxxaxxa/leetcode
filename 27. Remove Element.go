package main

//func main() {
//	ints := []int{3, 2, 2, 3}
//	fmt.Println(removeElement(ints, 3))
//	fmt.Println(ints)
//
//	int1 := []int{0, 1, 2, 2, 3, 0, 4, 2}
//	fmt.Println(removeElement(int1, 2))
//	fmt.Println(int1)
//
//}
func removeElement(nums []int, val int) int {
	if 0 == len(nums) {
		return 0
	}
	index := 0
	for _, num := range nums {
		if num != val {
			nums[index] = num
			index++
		}
	}
	return index
}

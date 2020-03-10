package main

//func main() {
//	fmt.Println(plusOne([]int{9}))
//	fmt.Println(plusOne([]int{1, 2, 3}))
//	fmt.Println(plusOne([]int{4, 3, 2, 1}))
//	fmt.Println(plusOne([]int{4, 3, 2, 9}))
//	fmt.Println(plusOne([]int{9, 9, 9, 9}))
//}
func plusOne(digits []int) []int {
	quo := 1
	for i, re, sum := len(digits)-1, 0, 0; i >= 0; i-- {
		sum = digits[i] + quo
		quo = sum / 10
		re = sum % 10
		digits[i] = re
		if quo == 0 {
			return digits
		}
	}

	if quo != 0 {
		return append([]int{quo}, digits...)
	}
	return nil
}

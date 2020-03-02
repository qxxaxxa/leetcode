package main

//func main() {
//	nums := []int{1, 2, 3, 4}
//	elist := decompressRLElist(nums)
//	for _, i := range elist {
//		fmt.Println(i)
//	}
//}

func decompressRLElist(nums []int) []int {
	lenth := 0
	for i, num := range nums {
		if i%2 == 0 {
			lenth += num
		}

	}
	re := make([]int, lenth)
	sp := 0
	for i := 0; i < len(nums); i += 2 {
		times := nums[i]
		cbyte := nums[i+1]
		for j := 0; j < times; j++ {
			re[sp] = cbyte
			sp++
		}
	}
	return re
}

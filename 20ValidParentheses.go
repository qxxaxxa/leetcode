package main

import (
	"container/list"
)

//func main() {
//	fmt.Println(isValid("]["))
//}

func isValid(s string) bool {
	if s == "" {
		return true
	}
	if len(s) == 1 {
		return false
	}
	ll := list.New()
	var pmap = map[int32]int32{
		40:  41,
		91:  93,
		123: 125,
	}
	for _, i := range s {
		switch i {
		case 40, 91, 123:
			ll.PushBack(i)
		case 41, 93, 125:
			back := ll.Back()
			if back == nil {
				return false
			}
			s3 := back.Value.(int32)
			if pmap[s3] == i {
				ll.Remove(back)
			} else {
				return false
			}
		}
	}
	if ll.Len() != 0 {
		return false
	}
	return true
}

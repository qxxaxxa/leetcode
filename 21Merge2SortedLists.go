package main

import "fmt"

//func main() {
//
//	l1 := &ListNode{
//		Val: 1,
//		Next: &ListNode{
//			Val: 2,
//			Next: &ListNode{
//				Val:  4,
//				Next: nil,
//			},
//		},
//	}
//	l2 := &ListNode{
//		Val: 1,
//		Next: &ListNode{
//			Val: 3,
//			Next: &ListNode{
//				Val:  4,
//				Next: nil,
//			},
//		},
//	}
//	//var l1 *ListNode
//	//var l2 *ListNode
//	node := mergeTwoLists(l1, l2)
//
//	node.print()
//
//}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	HEAD := new(ListNode)
	t := HEAD
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			t.Next = l1
			t = t.Next
			l1 = l1.Next
		} else {
			t.Next = l2
			t = t.Next
			l2 = l2.Next
		}
	}
	if l1 != nil {
		t.Next = l1
	}
	if l2 != nil {
		t.Next = l2
	}

	return HEAD.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func (n *ListNode) print() {
	for {
		if n == nil {
			fmt.Println("/")
			return
		}
		fmt.Print(n.Val)
		n = n.Next
	}

}

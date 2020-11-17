package main

import (
	"fmt"
	"hello/list/reverse"
)

func main() {
	//head := new(reverse.ListNode)
	//head.Value = 1
	//ln2 := new(reverse.ListNode)
	//ln2.Value = 2
	//ln3 := new(reverse.ListNode)
	//ln3.Value = 3
	//ln4 := new(reverse.ListNode)
	//ln4.Value = 4
	//ln5 := new(reverse.ListNode)
	//ln5.Value = 5
	//head.Next = ln2
	//ln2.Next = ln3
	//ln3.Next = ln4
	//ln4.Next = ln5

	slice := []int{1, 2, 3, 4, 5, 6}

	head := reverse.InitTailList(slice)
	for head != nil {
		fmt.Println(head.Value)
		head = head.Next
	}

	pre := reverse.Reverse(head)
	for pre != nil {
		fmt.Println(pre.Value)
		pre = pre.Next
	}
}

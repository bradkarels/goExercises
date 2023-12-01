package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func oddEvenList(head *ListNode) *ListNode {
	//isEven := true
	current := head
	var oone *ListNode
	for {
		// Some next nil checking...
		//if isEven {
		oone = current.Next                                       // 2
		fmt.Printf("oone: %d\n", oone.Val)                        //2
		current.Next = current.Next.Next                          // 3
		fmt.Printf("1. current.Next.Val: %d\n", current.Next.Val) //3
		//current.Next.Next = oone
		//fmt.Printf("1. current.Next.Next.Val: %d\n", current.Next.Next.Val) //2
		//printListNode(current)
		//} else {
		//}
		if current.Next == nil {
			break
		} else {
			current = current.Next
			continue
		}
		break
	}

	return head
}

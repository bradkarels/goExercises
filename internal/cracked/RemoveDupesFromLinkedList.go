package cracked

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func remDup(h *ListNode) *ListNode {
	if h == nil || h.Next == nil {
		return h
	}
	for h.Next != nil {
		n := h.Next
		//t := n.Val
		//c := n.Next
		for n.Next != nil {
			n = n.Next
		}
		h = h.Next
	}
	fmt.Println(h.Val)
	return h
}

func deleteNode(h *ListNode, t int) *ListNode {
	n := h
	if n.Val == t {
		return h.Next // Drop head
	}

	for n.Next != nil {
		if n.Next.Val == t {
			n.Next = n.Next.Next
			return h // No change to head
		}
		n = n.Next
	}
	return h
}

func printListNode(head *ListNode) {
	current := head
	for {
		if current.Next != nil {
			fmt.Printf("[%d, HasNext],", current.Val)
			current = current.Next
			continue
		} else {
			fmt.Printf("[%d, nil]\n", current.Val)
			break
		}
	}
}

package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteMiddle(head *ListNode) *ListNode {
	lllen := countEmUp(head, 0)

	mIdx := (lllen / 2) - 1
	idx := 0
	current := head
	switch {
	case lllen == 1:
		return nil
	case lllen == 2:
		current.Next = nil
		return current
	}
	for {
		fmt.Printf("idx: %d, mIdx: %d\n", idx, mIdx)
		if idx == mIdx {
			fmt.Printf("*current.Val = %d\n", current.Val)
			current.Next = current.Next.Next
			break
		} else {
			idx++
			fmt.Printf("current.Val = %d\n", current.Val)
			current = current.Next
			continue
		}
	}
	printListNode(head)
	return head
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

func countEmUp(head *ListNode, cnt int) int {
	cnt++
	if head.Next == nil {
		return cnt
	} else {
		return countEmUp(head.Next, cnt)
	}
}

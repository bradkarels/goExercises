package cracked

type Node struct {
	Val  int
	Next *Node
}

type LinkedList struct {
	head *Node
}

func (ll *LinkedList) IsEmpty() bool {
	return ll.head == nil
}

// Append to end O(n)
func (ll *LinkedList) append(i int) {
	n := Node{Val: i}
	if ll.IsEmpty() {
		ll.head = &n
		return
	}
	ptr := ll.head
	for {
		if ptr.Next == nil {
			ptr.Next = &n
			break
		}
		ptr = ptr.Next
	}
}

// if pos > len - add to end - not super efficient.
func (ll *LinkedList) InsertAt(pos, value int) {
	n := Node{Val: value}
	if ll.IsEmpty() {
		ll.head = &n
	}

	if pos == 0 {
		// replace head
		n.Next = ll.head
		ll.head = &n
	}

	i := 0
	ptr := ll.head
	// adjust for zero based index...
	for {
		i++
		if ptr.Next == nil {
			ptr.Next = &n
			break
		} else if i == pos {
			n.Next = ptr.Next
			ptr.Next = &n
			break
		} else {
			ptr = ptr.Next
		}
	}

}

func (n *Node) IsEmpty() bool {
	return n == nil
}

func Addz(h *Node, i int) *Node {
	// 1. allocate node
	// 2. put in the data
	// 3. Make next of new node as head
	newNode := &Node{Val: i, Next: h}
	// 4. Move the head to point to
	// the new node
	return newNode
}

func (n *Node) Add(i int) *Node {
	return &Node{Val: i, Next: n} // Replace head with added "top of stack _OR_ back of queue...
}

// Alias for Add - queue vs. stack
func (n *Node) Put(i int) *Node {
	return &Node{Val: i, Next: n} // Replace head with added "top of stack _OR_ back of queue...
}

func (n *Node) Pop() (*Node, int) {
	retVal := n.Val
	//n = n.Next
	return n.Next, retVal // advance head
}

func (n *Node) Peek() *Node {
	if n.IsEmpty() {
		return nil
	}
	previous := n
	current := n.Next
	for {
		if current == nil {
			return previous
		}
		previous = current
		current = current.Next
	}
}

func (n *Node) Remove() *Node {
	t := n
	if t.IsEmpty() {
		return nil
	}
	for t.Next != nil {
		if t.Next.Next == nil {
			ret := *t.Next
			t.Next = nil // Drop last Elem pointer?
			return &ret
		}
		t = t.Next
	}
	return nil
}

//func DoPop(s Stack) int {
//	return s.Pop().Val
//}

//func DoPut(i int, stack Stack) Node {
//	stack.Put(i)
//	return stack.Pop()
//}

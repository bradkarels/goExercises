package trees

import "fmt"

type Node struct {
	value int
	left  *Node
	right *Node
}

func visit(n *Node) {
	fmt.Printf("Visiting node: %d\n", n.value)
}

func InOrderTraversal(n *Node) {
	if n != nil {
		InOrderTraversal(n.left)
		visit(n)
		InOrderTraversal(n.right)
	}
}

func PreOrderTraversal(n *Node) {
	if n != nil {
		visit(n)
		PreOrderTraversal(n.left)
		PreOrderTraversal(n.right)
	}
}

func PostOrderTraversal(n *Node) {
	if n != nil {
		PostOrderTraversal(n.left)
		PostOrderTraversal(n.right)
		visit(n)
	}
}

package trees

func Insert(head, nieuwe *Node) {
	if head == nil {
		head = nieuwe
		return
	}
	//nieweVal := nieuwe.value
	current := head
	//for {
	//	if current.right != nil {
	//		current = current.right
	//		continue
	//	} else {
	//		current.right = nieuwe
	//		break
	//	}
	//}
	doInsert(current, nieuwe)
	PostOrderTraversal(head)
}

func swap(current, nieuwe *Node) {
	// TODO: Implement recursive
}

func doInsert(current, nieuwe *Node) {
	if current.right == nil {
		current.right = nieuwe
		return
	} else {
		doInsert(current.right, nieuwe)
	}
}

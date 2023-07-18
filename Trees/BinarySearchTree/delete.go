package main

func (n *Node) minValueNode() *Node {

	current := n

	for current.Left != nil {
		current = current.Left
	}

	return current

}

func (n *Node) delete(value int) *Node {

	if n == nil {
		return nil
	}

	if value < n.Value {
		n.Left = n.Left.delete(value)
	} else if value > n.Value {
		n.Right = n.Right.delete(value)
	} else {

		if n.Left == nil {
			return n.Right
		} else if n.Right == nil {
			return n.Left
		}

		minValueNode := n.Right.minValueNode()
		n.Value = minValueNode.Value
		n.Right = n.Right.delete(n.Value)

	}

	return n

}

func (t *BinaryTree) DeleteNode(value int) {

	t.Root = t.Root.delete(value)

}

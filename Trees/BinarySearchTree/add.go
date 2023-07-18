package main

func (t *BinaryTree) AddNode(value int) {

	if t.Root == nil {
		t.Root = &Node{Value: value}
	} else {
		t.Root.add(value)
	}
}

func (n *Node) add(value int) {

	if n.Value > value {

		if n.Left == nil {
			n.Left = &Node{Value: value}
		} else {
			n.Left.add(value)
		}

	} else if n.Value < value {

		if n.Right == nil {
			n.Right = &Node{Value: value}
		} else {
			n.Right.add(value)
		}

	}
}

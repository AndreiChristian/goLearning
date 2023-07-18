package main

func (t *BinaryTree) Has(value int) bool {

	if t.Root == nil {
		return false
	} else {
		return t.Root.search(value) != nil
	}

}

func (t *BinaryTree) SearchNode(value int) *Node {

	if t.Root == nil {
		return nil
	} else {
		return t.Root.search(value)
	}

}

func (n *Node) search(value int) *Node {

	if n == nil {
		return nil
	} else if n.Value == value {
		return n
	} else if n.Value > value {
		return n.Left.search(value)
	} else {
		return n.Right.search(value)
	}

}

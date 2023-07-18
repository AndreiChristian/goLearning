package main

import "fmt"

func (t *BinaryTree) IsBST() bool {

	if t.Root == nil {
		return true
	} else {
		return t.Root.IsBST()
	}

}

func (n *Node) IsBST() bool {

	isBST := true

	for isBST {

		isBST = n.isBSNode()
		n.Left.IsBST()
		n.Left.IsBST()

	}

	return isBST

}

func (n *Node) isBSNode() bool {

	if n == nil {

		return true
	} else if n.Left != nil && n.Right != nil {

		fmt.Println("1", &n)
		return ((n.Left.Value < n.Value) && (n.Value < n.Right.Value))

	} else if n.Left == nil {

		fmt.Println("2", &n)
		// we have only right

		return n.Value < n.Right.Value

	} else {
		fmt.Println("3", &n)

		return n.Left.Value < n.Value

	}

}

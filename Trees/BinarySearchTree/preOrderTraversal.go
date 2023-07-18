package main

import "fmt"

func (n *Node) preOrderTraversal() {

	if n != nil {
		// ROOT
		fmt.Println(n.Value)
		// LEFT
		n.Left.preOrderTraversal()
		// RIGHT
		n.Right.preOrderTraversal()
	}

}

func (t *BinaryTree) PreOrderTraversal() {

	if t.Root != nil {

		fmt.Println("Starting Pre Order traversal")
		t.Root.preOrderTraversal()

	}

}

package main

import "fmt"

func (n *Node) inOrderTraversal() {

	if n != nil {
		// LEFT
		n.Left.inOrderTraversal()
		// ROOT
		fmt.Println(n.Value)
		// RIGHT
		n.Right.inOrderTraversal()
	}

}

func (t *BinaryTree) InOrderTraversal() {

	if t.Root != nil {

		fmt.Println("Starting In Order traversal")
		t.Root.inOrderTraversal()

	}

}

package main

import "fmt"

func (n *Node) postOrderTraversal() {

	if n != nil {

		n.Left.postOrderTraversal()
		n.Right.postOrderTraversal()
		fmt.Println(n.Value)

	}

}

func (t *BinaryTree) PostOrderTraversal() {

	if t.Root != nil {
		fmt.Println("Starting Post Order Traversal")
		t.Root.postOrderTraversal()
	}

}

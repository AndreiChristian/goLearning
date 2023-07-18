package main

import "fmt"

func (t *BinaryTree) Invert() {

	fmt.Println("After being inverted, this will no longer be a BST")

	if t.Root != nil {
		t.Root.invert()
	}

}

func (n *Node) invert() {

	if n == nil {
		return
	}

	temp := n.Left
	n.Left = n.Right
	n.Right = temp

	n.Left.invert()
	n.Right.invert()

}

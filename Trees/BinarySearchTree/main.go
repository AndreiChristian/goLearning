package main

import (
	"math/rand"
)

type BinaryTree struct {
	Root *Node
}

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func main() {

	BT := NewBinaryTree()

	i := 100

	for i > 0 {

		BT.AddNode(rand.Intn(20))
		i--

	}

	print(BT.Root, "", true)

	BT.DeleteNode(4)
	BT.DeleteNode(5)
	BT.DeleteNode(7)
	BT.DeleteNode(13)
	BT.DeleteNode(14)
	BT.DeleteNode(19)
	BT.DeleteNode(1)
	print(BT.Root, "", true)
}

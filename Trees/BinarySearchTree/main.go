package main

type BinaryTree struct {
	Root *Node
}

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func NewBinaryTree() *BinaryTree {
	return &BinaryTree{}
}

func main() {

	BT := NewBinaryTree()
	BT.AddNode(100)
	BT.AddNode(7)
	BT.AddNode(2)
	BT.AddNode(4)
	BT.AddNode(1)
	BT.AddNode(9)
	BT.AddNode(8)

	print(BT.Root, "", true)
}

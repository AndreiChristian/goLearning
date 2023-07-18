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

func main() {

	BT := NewBinaryTree()
	BT.AddNode(100)
	BT.AddNode(7)
	BT.AddNode(2)
	BT.AddNode(4)
	BT.AddNode(1)
	BT.AddNode(9)
	BT.AddNode(8)

	Print(BT.Root, "", true)
}

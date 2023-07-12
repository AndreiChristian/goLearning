package tree

var T Tree

type Node struct {
	name     string
	size     int
	children []*Node
	parent   *Node
}

type Tree struct {
	root *Node
}

func NewTree() {

	T = Tree{
		root: &Node{},
	}

}

func CreateNode(parent *Node) *Node {

	return &Node{
		parent: parent,
	}

}

func (n *Node) SetNodeName(name string) {
	n.name = name
}

func (n *Node) SetSize(size int) {
	n.size = size
}

func (n *Node) SetChildren(nodes []*Node) {

	n.children = nodes

}

func MoveUP(node *Node) *Node {

	return node.parent

}

func MoveDown(node *Node, childName string) *Node {

	for _, child := range node.children {
		if child.name == childName {
			return child
		}
	}

	return nil

}

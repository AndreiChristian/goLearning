package main

import "fmt"

type BST struct{
	root *Node
}

type Node struct{
	Value int
	Left *Node
	Right *Node
}

func (t *BST) InsertValue(value int) *BST{

	if t.root == nil {

		t.root = &Node{Value: value}
	} else {

		t.root.Insert(value)
	}

	return t
}

func Print(n *Node, value int) { 

	if n == nil{
		return
	}

	space := ""
	i:=0

	for i<=value{

		space = space + " "
}

	fmt.Println(space,n.Value)

	Print(n.Left,value + 1)
	Print(n.Right, value + 1)
	
}

func (n *Node) Insert(value int) {
	if n==nil{
		return
	} else if value <= n.Value{

		if n.Left == nil{

			n.Left = &Node{Value: value}

		} else {

			n.Left.Insert(value)
		}

	} else {

		if n.Right == nil {

			n.Right = &Node{Value: value}

		} else {

			n.Right.Insert(value)
		}
	}
}

func (t *BST) SearchNode(value int) *Node{
	return t.root.Search(value)
}

func (n *Node) Search(value int ) *Node {
	fmt.Println("wORKING")
	if n == nil {
	
		return nil
	}

	switch{
	case n.Value < value :
		return n.Right.Search(value)
	case n.Value > value :
		return n.Left.Search(value)
	default :	
		return n
	}
}

func main(){

	bst := BST{}

	bst.InsertValue(4)
	bst.InsertValue(1)
	bst.InsertValue(5)
	bst.InsertValue(8)
	bst.InsertValue(9)
	bst.InsertValue(2)
	bst.InsertValue(14)
	bst.InsertValue(0)
	bst.InsertValue(3)
	bst.InsertValue(100)
	bst.InsertValue(15)

	fmt.Println("done inserting")

	Print(bst.root, 0)

	node := bst.SearchNode(4)

	fmt.Println(node)
}


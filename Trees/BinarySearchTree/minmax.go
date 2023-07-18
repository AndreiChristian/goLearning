package main

import "errors"

func (t *BinaryTree) Min() (int, error) {

	if t.Root == nil {
		return 0, errors.New("the tree is null")
	} else {

		currentNode := t.Root

		for currentNode.Left != nil {
			currentNode = currentNode.Left
		}

		return currentNode.Value, nil

	}

}

func (t *BinaryTree) Max() (int, error) {

	if t.Root == nil {
		return 0, errors.New("the tree is null")
	} else {

		currentNode := t.Root

		for currentNode.Right != nil {
			currentNode = currentNode.Right
		}

		return currentNode.Value, nil

	}

}

package main

import "fmt"

func Print(node *Node, prefix string, isTail bool) {
    if node == nil {
        return
    }

    fmt.Println(prefix + (map[bool]string{false: "├── ", true: "└── "}[isTail]) + fmt.Sprintf("%v", node.Value))
    var child []*Node
    if node.Left != nil {
        child = append(child, node.Left)
    }
    if node.Right != nil {
        child = append(child, node.Right)
    }
    for i := 0; i < len(child); i++ {
        print(child[i], prefix+(map[bool]string{false: "│   ", true: "    "}[isTail]), i == len(child)-1)
    }
}

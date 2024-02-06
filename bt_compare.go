package main

import (
    "fmt"
)

type BinaryTreeNode struct {
    value int
    left *BinaryTreeNode
    right *BinaryTreeNode
}

func compare(a *BinaryTreeNode, b *BinaryTreeNode) bool {

    if a == nil && b == nil {
        return true
    }
    
    if a == nil || b == nil {
        return false
    }

    if a.value != b.value {
        return false
    }
    return compare(a.left, b.left) && compare(a.right, b.right)
}

func main() {
    a := &BinaryTreeNode{value: 5}
    a.left = &BinaryTreeNode{value: 4}
    a.right = &BinaryTreeNode{value: 3}
    b := &BinaryTreeNode{value: 5}
    b.left = &BinaryTreeNode{value: 4}
    b.right = &BinaryTreeNode{value: 2}
    result := compare(a, b)
    fmt.Println(result)
}

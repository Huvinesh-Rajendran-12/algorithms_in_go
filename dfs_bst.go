package main

import (
    "fmt"
)

type BinaryTreeNode struct {
    value int
    left *BinaryTreeNode
    right *BinaryTreeNode
}

func dfs_search(curr *BinaryTreeNode, needle int) bool {
    if curr == nil {
        return false
    }
    if curr.value == needle {
        return true
    }
    if curr.value < needle {
        return dfs_search(curr.right, needle)
    }
    if curr.value > needle {
        return dfs_search(curr.left, needle)
    }
    return false

}


func main() {
    head := BinaryTreeNode{value: 7, left: nil, right: nil}
    head.left = &BinaryTreeNode{value: 4}
    head.right = &BinaryTreeNode{value: 23}
    head.left.left = &BinaryTreeNode{value: 2}
    head.left.right = &BinaryTreeNode{value: 3}
    result := dfs_search(&head, 10)
    fmt.Println(result)
}

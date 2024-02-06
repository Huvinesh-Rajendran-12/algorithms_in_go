package main

import (
    "fmt"
)

type BinaryTreeNode struct {
    value int
    left *BinaryTreeNode
    right *BinaryTreeNode
}

func post_order_walk(curr *BinaryTreeNode, path *[]int) {

    if curr == nil{
        return
    }
    post_order_walk(curr.left, path)
    post_order_walk(curr.right, path)
    *path = append(*path, curr.value)

}

func post_order_search(head *BinaryTreeNode) []int {
    path := []int{}
    post_order_walk(head, &path)
    return path
}

func main() {
    head := BinaryTreeNode{value: 7, left: nil, right: nil}
    head.left = &BinaryTreeNode{value: 23}
    head.right = &BinaryTreeNode{value: 4}
    head.left.left = &BinaryTreeNode{value: 9}
    head.left.right = &BinaryTreeNode{value: 3}
    path := post_order_search(&head)
    fmt.Println(path)
}

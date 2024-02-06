
package main

import (
    "fmt"
)

type BinaryTreeNode struct {
    value int
    left *BinaryTreeNode
    right *BinaryTreeNode
}

func in_order_walk(curr *BinaryTreeNode, path *[]int) {

    if curr == nil{
        return
    }


    in_order_walk(curr.left, path)
    *path = append(*path, curr.value)
    in_order_walk(curr.right, path)

}

func in_order_search(head *BinaryTreeNode) []int {
    path := []int{}
    in_order_walk(head, &path)
    return path
}

func main() {
    head := BinaryTreeNode{value: 7, left: nil, right: nil}
    head.left = &BinaryTreeNode{value: 23}
    head.right = &BinaryTreeNode{value: 4}
    head.left.left = &BinaryTreeNode{value: 9}
    head.left.right = &BinaryTreeNode{value: 3}
    path := in_order_search(&head)
    fmt.Println(path)
}

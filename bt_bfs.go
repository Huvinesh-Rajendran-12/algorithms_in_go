package main

import (
    "fmt"
)

type BinaryTreeNode struct {
    value int
    left *BinaryTreeNode
    right *BinaryTreeNode
}

type Node struct{
    item *BinaryTreeNode
    next *Node
}

type Queue struct {
    length int
    head *Node
    tail *Node
}

func (q *Queue) Enqueue(value *BinaryTreeNode) {
    node := &Node{value, nil}
    if q.head == nil {
        q.head = node
        q.tail = node
    } else {
        q.tail.next = node
        q.tail = node
    }
    q.length ++
    fmt.Printf("Enqueued %d\n", value)
}

func (q *Queue) Dequeue() *BinaryTreeNode {
    if q.tail == nil {
       panic("Queue is empty")
    }
    value := q.head.item
    q.head = q.head.next
    if q.head ==  nil {
	q.tail = nil 
    }
    q.length --
    return value
}

func (q *Queue) peek() {
    value := q.head.item
    fmt.Printf("The value of the head :", value)
	
}

func bfs(head *BinaryTreeNode, needle int) bool{
    q := Queue{}
    q.Enqueue(head)
    for q.length > 0 {
        curr := q.Dequeue()
        if curr.value == needle {
            return true
        }
        if curr.left != nil {
            q.Enqueue(curr.left)
        }
        if curr.right != nil {
            q.Enqueue(curr.right)
        }

    }
    return false
} 

func main() {
    head := BinaryTreeNode{value: 7, left: nil, right: nil}
    head.left = &BinaryTreeNode{value: 23}
    head.right = &BinaryTreeNode{value: 4}
    head.left.left = &BinaryTreeNode{value: 9}
    head.left.right = &BinaryTreeNode{value: 3}
    result := bfs(&head, 40)
    fmt.Println(result)
}

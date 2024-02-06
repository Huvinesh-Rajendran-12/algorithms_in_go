package main

import (
    "fmt"
)

type Node struct {
    Value int
    Next *Node
}

type Queue struct {
    Head *Node
    Tail *Node
}

func (q *Queue) Enqueue(value int) {
    node := &Node{value, nil}
    if q.Head == nil {
        q.Head = node
        q.Tail = node
    } else {
        q.Tail.Next = node
        q.Tail = node
    }
    fmt.Printf("Enqueued %d\n", value)
}

func (q *Queue) Dequeue() int {
    if q.Tail == nil {
       panic("Queue is empty")
    }
    value := q.Head.Value
    q.Head = q.Head.Next
    if q.Head ==  nil {
	q.Tail = nil 
    }
    return value
}

func (q *Queue) peek() {
    value := q.Head.Value
    fmt.Printf("The value of the head :", value)
	
}

func main()  {
    q := Queue{}
    q.Enqueue(1)
    q.Enqueue(2)
    q.Enqueue(3)
    fmt.Println(q.Dequeue())
}

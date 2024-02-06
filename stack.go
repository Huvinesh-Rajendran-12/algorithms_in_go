package main

import (
    "fmt"
    "math"
)

type StackNode struct {
    value int
    prev *StackNode
}

type Stack struct {
    head *StackNode
    length int
}

func (s *Stack) push(value int)  {
    s.length += 1
    node := &StackNode{value, nil}
    if s.head == nil {
        s.head = node
        return
    }

    node.prev = s.head
    s.head = node
}

func (s *Stack) pop() int {

    s.length = int(math.Max(0, float64(s.length) - 1))
    if s.length == 0 {
        head := s.head 
        s.head = nil
        if head == nil {
            return -1
        }
        return head.value
    } else {
        head := s.head 
        s.head = s.head.prev 
        return head.value
    }
}

func main() {
    s := Stack{}
    s.push(1)
    s.push(2)
    s.push(3)

    fmt.Println(s.pop())
    fmt.Println(s.pop())
    fmt.Println(s.pop())
    fmt.Println(s.pop())
}


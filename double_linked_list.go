package main

import (
    "fmt"
)

type DoubleNode struct {
    value int
    next *DoubleNode
    prev *DoubleNode
}

type DoubleLinkedList struct {
    head *DoubleNode
    tail *DoubleNode
    length int

}

func (d *DoubleLinkedList) prepend(value int) {
    node := &DoubleNode{value: value}
    d.length += 1
    if d.head == nil {
        d.head = node
        d.tail = node
        return
    }
    node.next = d.head
    d.head.prev = node
    d.head = node
}

func (d *DoubleLinkedList) insertAt(value int, idx int) {

    if idx > d.length {
        panic("Index out of bounds")
    } else if idx == d.length {
        d.append(value)
    
    } else if idx == 0 {
        d.prepend(value)
    }
    node := &DoubleNode{value: value}

    curr_pointer := d.head
    for i := 0; i < idx; i++ {
        curr_pointer = curr_pointer.next
    }
    node.next = curr_pointer
    node.prev = curr_pointer.prev
    if curr_pointer.prev != nil {
        curr_pointer.prev.next = node
    }
    curr_pointer.prev = node
    d.length += 1
}

func (d *DoubleLinkedList) append(value int) {
    node := &DoubleNode{value: value}
    d.length += 1
    if d.tail == nil {
        d.head = node
        d.tail = node 
        return
    }
    node.prev = d.tail
    d.tail.next = node
    d.tail = node
}

func (d *DoubleLinkedList) remove(value int) {
    curr_pointer := d.head
    for i := 0; i < d.length; i++ {
        if curr_pointer.value == value {
            break
        }
        curr_pointer = curr_pointer.next
    }
    if curr_pointer == nil {
        return 
    }
    if curr_pointer.prev != nil {
        curr_pointer.prev.next = curr_pointer.next
    }
    if curr_pointer.next != nil {
        curr_pointer.next.prev = curr_pointer.prev
    }
    if curr_pointer == d.head {
        d.head = curr_pointer.next
    }
    if curr_pointer == d.tail {
        d.tail = curr_pointer.prev
    }
    curr_pointer.prev = nil
    curr_pointer.next = nil
    d.length -= 1

}

func (d *DoubleLinkedList) get(idx int) int {
    if idx > d.length {
        panic("Index out of bounds")
    }
    curr_pointer := d.head
    for i := 0; i < idx; i++ {
        curr_pointer = curr_pointer.next
    }
    return curr_pointer.value
}

func (d *DoubleLinkedList) removeAt(idx int) {

    if idx > d.length {
        panic("Index out of bounds")
    }
    curr_pointer := d.head
    for i := 0; i < idx; i++ {
        curr_pointer = curr_pointer.next
    }
    if curr_pointer.prev != nil {
        curr_pointer.prev.next = curr_pointer.next
    }
    if curr_pointer.next != nil {
        curr_pointer.next.prev = curr_pointer.prev
    }
    if curr_pointer == d.head {
        curr_pointer.next.prev = nil
        d.head = curr_pointer.next
    }
    if curr_pointer == d.tail {
        curr_pointer.prev.next = nil
        d.tail = curr_pointer.prev
        curr_pointer.prev = nil
    }
    d.length -= 1
}

func main() {
    d := DoubleLinkedList{}
    d.prepend(1)
    d.append(2)
    d.append(3)
    fmt.Println(d.get(2))
    d.removeAt(2)
    curr_pointer := d.head
    for i := 0; i < d.length; i++ {
        if curr_pointer != nil {
            fmt.Println(curr_pointer.value)
            fmt.Println("length", d.length)
            fmt.Println("head", d.head.value)
            fmt.Println("tail", d.tail.value)
        }

        curr_pointer = curr_pointer.next
    }
}


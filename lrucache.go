package main

import (
    "fmt"
)

type ValContainer struct {
    val int
    next *ValContainer
    prev *ValContainer
}

type LRU struct {
    head *ValContainer
    tail *ValContainer
    length int
    lookup map[int]ValContainer
    reverseLookup map[ValContainer]int
    capacity int
}

func (l *LRU) update(key int, val int) {
    Container , ok:= l.lookup[key]
    if !ok {
        container := ValContainer{val: val}
        l.prepend(&container)
        l.length++
        l.trimCache()
        l.lookup[key] = container
        l.reverseLookup[container] = key
    } else {
        l.detach(&Container)
        l.prepend(&Container)
        Container.val = val
    }
}

func (l *LRU) get(key int) int {
    valContainer, ok := l.lookup[key]
    if !ok {
        return -1
    }
    l.detach(&valContainer)
    l.prepend(&valContainer)
    return valContainer.val
}

func (l *LRU) detach(val *ValContainer) {
    if val.next != nil {
        val.next.prev = val.prev
    }
    if val.prev != nil {
        val.prev.next = val.next
    }
    if l.head == val {
        l.head = l.head.next
    }
    if l.tail == val {
        l.tail = l.tail.prev
    }
    val.next = nil
    val.prev = nil
}

func (l *LRU) prepend(val *ValContainer) {
    if l.head == nil {
        l.head = val
        l.tail = l.head
    }
    l.head.prev = val
    val.next = l.head
    l.head = val
}

func (l *LRU) trimCache(){
    if l.length <= l.capacity {
        return       
    }
    tail := l.tail
    l.detach(tail)
    key := l.reverseLookup[*tail]
    delete(l.lookup, key)
    delete(l.reverseLookup, *tail)
    l.length--
}

func main() {
    lru := LRU{length: 0, lookup: map[int]ValContainer{}, 
    reverseLookup: map[ValContainer]int{}, capacity: 10}
    lru.update(1, 41)
    lru.update(21, 51)
    lru.update(20, 40)
    lru.update(18, 39)
    lru.update(13, 26)
    lru.update(14, 15)
    lru.update(12, 76)
    lru.update(11, 23)
    lru.update(7, 24)
    lru.update(10, 21)
    lru.update(6, 7)
    result := lru.get(12)
    fmt.Println(result)
}

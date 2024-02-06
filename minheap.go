package main

import "fmt"

type MinHeap struct {
    length int
    data []int
}

func (m *MinHeap) insert(val int) {
    m.data = append(m.data, val)
    heapifyUp(m, m.length)
    m.length ++
}

func (m *MinHeap) delete() int {
    if m.length == 0 {
        return -1
    }
    out := m.data[0]
    m.length --
    if m.length == 0 {
        m.data = m.data[:0]
        return out
    }
    m.data[0] = m.data[m.length]
    heapifyDown(m, 0)
    return out
}

func parent(idx int) int {
    return (idx - 1) / 2
}

func left_child(idx int) int {
    return idx * 2 + 1 
}

func right_child(idx int) int {
    return idx * 2 + 2
}

func swap_val(m *MinHeap ,idx1 int, idx2 int, val1 int, val2 int) {
    m.data[idx1] = val2
    m.data[idx2] = val1
}

func  heapifyUp(m *MinHeap , idx int) {
    if idx == 0 {
        return 
    }
    parent := parent(idx)
    parent_val := m.data[parent]
    val := m.data[idx]
    if parent_val > val {
        swap_val(m, idx, parent, val, parent_val)
        heapifyUp(m, parent)
    }
}

func heapifyDown(m *MinHeap, idx int) {
    if idx >= m.length {
        return
    }
    left_idx := left_child(idx)
    right_idx := right_child(idx)

    if left_idx >= idx {
        return 
    }
    if right_idx >= idx {
        return
    }

    left_val := m.data[left_idx]
    right_val := m.data[right_idx]
    val := m.data[idx]

    if left_val > right_val && val > right_val {
        swap_val(m, idx, right_idx, val, right_val)
    }else if right_val > left_val && val > left_val {
        swap_val(m, idx, left_idx, val, left_val)
    }
}

func main() {
    m := MinHeap{length: 0, data: []int{}}
    m.insert(5)
    m.insert(20)
    m.insert(6)
    m.insert(7)
    result := m.delete()
    fmt.Println(result)
}

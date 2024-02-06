package main

import (
	"fmt"
)

type AdjacencyMatrix struct {
    data [][]int
    length int
}

func reverseArray(arr *[]int) {
    for i := 0; i < len(*arr)/2; i++ {
        j := len(*arr) - i - 1
        (*arr)[i], (*arr)[j] = (*arr)[j], (*arr)[i]
    }
}


func bfs(graph AdjacencyMatrix, source int, needle int) []int {
    seen := make([]bool, graph.length)
    prev := make([]int, graph.length)
    for i := range prev {
        prev[i] = -1
    }
    seen[source] = true
    queue := []int{source}
    for graph.length > 0 {
        curr := queue[0]
        queue = queue[1:]
        if curr == needle {
            break
        }
        for i := 0; i < graph.length; i++{
            if graph.data[curr][i] == 0 {
                continue
            }
            if seen[i] == true {
                continue
            }
            seen[i] = true
            prev[i] = curr
            queue = append(queue, i)

        }
    }
    curr := needle 
    out := []int{}
    for prev[curr] != -1 {
        out = append(out, curr)
        curr = prev[curr]
        
    }
    if len(out) > 0 {
        reverseArray(&out)
        out = append([]int{source}, out...)
        return out
    }
    return []int{}

}

func main() {
    graph := AdjacencyMatrix{
		data: [][]int{
            {0, 3, 1, 0, 0, 0, 0},
            {0, 0, 0, 0, 1, 0, 0},
            {0, 0, 7, 0, 0, 0, 0},
            {0, 0, 0, 0, 0, 0, 0},
            {0, 1, 0, 5, 0, 2, 0},
            {0, 0, 18, 0, 0, 0, 1},
            {0, 0, 0, 1, 0, 0, 1},
		},
		length: 7,
	}

	startVertex := 0
	needleVertex := 6
	fmt.Printf("BFS traversal starting from vertex %d and searching for vertex %d\n", startVertex, needleVertex)
	result := bfs(graph, startVertex, needleVertex)
	fmt.Println("Result:", result)
}

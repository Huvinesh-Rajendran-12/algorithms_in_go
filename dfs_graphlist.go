package main

import (
    "fmt"
)

type GraphEdge struct {
    to int
    weight int
}

type AdjacencyList struct {
    data [][]GraphEdge
    length int
}

func walk(graph AdjacencyList, curr int, needle int, seen []bool, path *[]int) bool {
    if seen[curr] == true {
        return false
    }
    seen[curr] = true
    // pre
    *path = append(*path, curr)
    if curr == needle {
        return true
    }
    // recurse
    list := graph.data[curr]
    for i := 0; i < len(list); i++ {
        edge := list[i]
        if walk(graph, edge.to, needle, seen, path) == true {
            return true
        }
    }

    //post 
    *path = (*path)[:len(*path) - 1]
    return false
}

func dfs(graph AdjacencyList, source int, needle int) []int {
    seen := make([]bool, graph.length)
    path := []int{}
    walk(graph, source, needle, seen, &path)
    return path
}

func main() {
    list2 := make([][]GraphEdge, 7)

	list2[0] = []GraphEdge{
		{to: 1, weight: 3},
		{to: 2, weight: 1},
	}
	list2[1] = []GraphEdge{
		{to: 4, weight: 1},
	}
	list2[2] = []GraphEdge{
		{to: 3, weight: 7},
	}
	list2[3] = []GraphEdge{}
	list2[4] = []GraphEdge{
		{to: 1, weight: 1},
		{to: 3, weight: 5},
		{to: 5, weight: 2},
	}
	list2[5] = []GraphEdge{
		{to: 2, weight: 18},
		{to: 6, weight: 1},
	}
	list2[6] = []GraphEdge{
		{to: 3, weight: 1},
	}
    fmt.Println("length of the list:", len(list2))
    graph := &AdjacencyList{data: list2, length: len(list2)}
    source := 0
    needle := 6
    path := dfs(*graph, source, needle)
    fmt.Println(path)
}

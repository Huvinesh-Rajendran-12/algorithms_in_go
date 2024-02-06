package main

import (
	"fmt"
	"math"
)

type GraphEdge struct {
    to int
    weight int
}

type AdjacencyList struct {
    data [][]GraphEdge
    length int
}

func hasUnseen(seen []bool, dist []int) bool {
    for i := 0; i < len(seen); i++ {
        if seen[i] == false || dist[i] == int(math.Inf(1)) {
            return true
        }

    }
    return false
}
func reverseArray(arr *[]int) {
    for i := 0; i < len(*arr)/2; i++ {
        j := len(*arr) - i - 1
        (*arr)[i], (*arr)[j] = (*arr)[j], (*arr)[i]
    }
}

func getLowestUnvisited(seen []bool, dist []int) int {
    idx := -1
    lowestDist := int(math.Inf(1))
    for i := 0; i < len(seen); i++ {
        if seen[i] == true {
            continue
        }
        if lowestDist > dist[i] {
            lowestDist = dist[i]
            idx = i
        }
    }
    return idx
}

func djikstra_list(source int, sink int, arr *AdjacencyList) []int {
    seen := make([]bool, arr.length)
    dist := make([]int, arr.length)
    prev := make([]int, arr.length)
    for i := 0; i < arr.length; i++ {
        dist[i] = int(math.Inf(1))
    }
    for i:=0; i < arr.length; i++{
        prev[i] = -1
    }
    dist[source] = 0

    for hasUnseen(seen, dist) == true {

        curr := getLowestUnvisited(seen, dist)
        seen[curr] = true
        for i := 0; i < len(arr.data[curr]); i++ {
            edge := arr.data[curr][i]
            if seen[edge.to] == true {
                continue
            }
            dist_ := dist[curr] + edge.weight
            if dist_ < dist[edge.to] {
                dist[edge.to] = dist_
                prev[edge.to] = curr
            }

        }
    }

    out := []int{}
    curr := sink
    for prev[curr] != -1 {
        out = append(out, curr)
        curr = prev[curr]
    }
    reverseArray(&out)
    out = append([]int{source}, out...)
    return out
}

func main() {
    list1 := make([][]GraphEdge, 7)

    list1[0] = []GraphEdge{{to: 1, weight: 3}, {to: 2, weight: 1}}
    list1[1] = []GraphEdge{{to: 0, weight: 3}, {to: 2, weight: 4}, {to: 4, weight: 1}}
    list1[2] = []GraphEdge{{to: 1, weight: 4}, {to: 3, weight: 7}, {to: 0, weight: 1}}
    list1[3] = []GraphEdge{{to: 2, weight: 7}, {to: 4, weight: 5}, {to: 6, weight: 1}}
    list1[4] = []GraphEdge{{to: 1, weight: 1}, {to: 3, weight: 5}, {to: 5, weight: 2}}
    list1[5] = []GraphEdge{{to: 6, weight: 1}, {to: 4, weight: 2}, {to: 2, weight: 18}}
    list1[6] = []GraphEdge{{to: 3, weight: 1}, {to: 5, weight: 1}}
    graph := AdjacencyList{data: list1, length: len(list1)}
    result := djikstra_list(0,6,&graph)
    fmt.Println(result)
}



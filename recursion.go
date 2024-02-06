package main


import (
    "fmt"
)

type Point struct{
    x int
    y int
}


var direction = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func walk(maze [][]string, wall string, curr Point, end Point, seen [][]bool, path *[]Point) bool {
    if (curr.x < 0 || curr.x >= len(maze[0])|| curr.y < 0 || curr.y >= len(maze)) {
        return false
    }

    if (maze[curr.x][curr.y] == wall) {
        return false 
    }

    if (curr.x == end.x && curr.y == end.y) {
        *path = append(*path, end)
        return true
    }
    if (seen[curr.x][curr.y]) {
        return false
    }
    
    seen[curr.x][curr.y] = true
    *path = append(*path, curr)

    for i := 0; i < len(direction); i++ {
        x := curr.x + direction[i][0]
        y := curr.y + direction[i][1]
        new_curr := Point{x, y}
        if walk(maze, wall, new_curr, end, seen, path) == true {
            return true
        }
    }

    *path = (*path)[:len(*path) -1]
    return false
    
}


func solve(maze [][]string, wall string, curr Point, end Point) []Point {
    seen := make([][]bool, len(maze))
    for i := range seen {
        seen[i] = make([]bool, len(maze[0]))
    }
    var path []Point
    walk(maze, wall, curr, end, seen, &path)
    return path
}

func main() {
    maze := [][]string{
        {"0", "0", "1", "0", "0"},
        {"0", "1", "1", "0", "1"},
        {"0", "0", "0", "0", "0"},
        {"1", "0", "1", "1", "0"},
        {"0", "0", "0", "0", "0"},
    }

    wall := "1"
    start := Point{0, 0}
    end := Point{4, 4}

    path := solve(maze, wall, start, end)

    for _, p := range path {
        fmt.Printf("(%d, %d) -> ", p.x, p.y)
    }
    fmt.Println("End")
}

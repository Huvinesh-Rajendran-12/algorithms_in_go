package main

import (
	"fmt"
    "time"
)

func bubble_sort(arr []int) []int {
    for i := 0; i < len(arr); i++ {
        for j := 0; j < len(arr)-i-1; j++ {
            if arr[j] > arr[j+1] {
                arr[j], arr[j+1] = arr[j+1], arr[j]
            }
        }
    }
    return arr
}

func main() {
    start := time.Now()
    arr := []int{1, 3, 2, 5, 4, 6, 9, 8, 7}
    elapsed := time.Since(start)
    fmt.Println(bubble_sort(arr))
    fmt.Println("Elapsed time: ", elapsed)
}
	


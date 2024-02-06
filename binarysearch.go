package main

import (
	"fmt"
	"time"
)

func binarySearch(arr []int, target int) bool {
	lower := 0
	highest := len(arr) - 1
	for lower <= highest {
		mid := (lower + highest) / 2
		val := arr[mid]
		if val == target {
			return true
		} else if val > target {
			highest = mid
		} else {
			lower = mid + 1
		}

	}
	return false
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	var input int
	fmt.Printf("Enter an integer input:\n")
	fmt.Scan(&input)
	start := time.Now()
	res := binarySearch(arr, input)
	elapsed := time.Since(start)
	if res == true {
		fmt.Printf("Success!\n")
	} else {
		fmt.Printf("Failed to find")
	}
	fmt.Printf("Time elapsed: ", elapsed)
}

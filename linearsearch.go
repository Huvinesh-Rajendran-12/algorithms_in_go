package main

import (
	"fmt"
	"math/rand"
	"time"
)

func linear_search(arr []int, target int) int {
	for i, num := range arr {
		if num == target {
			return i
		}
	}
	return -1
}

func main() {
	rand.Seed(time.Now().UnixNano()) // Set the seed for random number generation

	arr := make([]int, 6)
	for i := 0; i < 6; i++ {
		arr[i] = rand.Intn(100) // Generate a random integer between 0 and 99
	}
	var input int
	fmt.Printf("Enter an integer input (0-99) :\n")
	fmt.Scan(&input)
	index := linear_search(arr, input)
	if index == -1 {
		fmt.Printf("False")
	} else {
		fmt.Printf("True")
	}
}

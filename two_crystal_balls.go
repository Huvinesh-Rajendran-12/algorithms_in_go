package main

import "math"

func two_crystal_balls(arr []bool) int {
	jumpAmount := math.Floor(math.Sqrt(float64(len(arr))))
	i := int(jumpAmount)
	for ; i < len(arr); i += int(jumpAmount) {
		if arr[i] == true {
			break
		}
	}

	i -= int(jumpAmount)

	for j := 0; j < int(jumpAmount) && i < len(arr); j += 1 {
		if arr[i] == true {
			return i
		}
		i += 1
	}
	return -1
}

func main() {
	arr := []bool{false, false, false, true, true, true}
	result := two_crystal_balls(arr)
  if result != -1 {
    
	println("result :", result)
  }
}

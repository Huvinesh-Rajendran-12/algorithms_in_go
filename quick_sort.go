package main

import "fmt"

 func qs(arr []int, low int, high int) {
     if low >= high {
         return
     }
     pivotIdx := partition(arr, low, high)
     qs(arr, low, pivotIdx - 1)
     qs(arr, pivotIdx + 1, high)
 }

 func partition(arr []int, low int, high int) int {

     pivot := arr[high]
     idx := low - 1

     for i := low; i < high; i++ {
         if arr[i] <= pivot {
            idx++
            tmp := arr[i]
            arr[i] = arr[idx]
            arr[idx] = tmp
         }
     }
     idx++
     arr[high] = arr[idx]
     arr[idx] = pivot
     
     return idx
 }

func quick_sort(arr *[]int) {
    qs(*arr, 0, len(*arr)-1)
}

func main() {
    arr := []int{3, 1, 2, 4, 9, 7, 6, 8}
    quick_sort(&arr)

    fmt.Println(arr)
}


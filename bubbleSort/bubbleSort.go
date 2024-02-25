package main

import "fmt"

func main() {
	arr := []int{64, 25, 12, 22, 11}
	fmt.Println("Unsorted Array: ", arr)

	arr = bubbleSort(arr)
	fmt.Println("Sorted Array: ", arr)
}

func bubbleSort(arr []int) []int {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				tmp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = tmp
			}
		}
	}

	return arr
}

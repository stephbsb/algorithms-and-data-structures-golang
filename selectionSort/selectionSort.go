package main

import "fmt"

func main() {
	arr := [5]int{64, 25, 12, 22, 11}
	fmt.Println("Unsorted Array: ", arr)

	arr = selectionSort(arr)
	fmt.Println("Sorted Array: ", arr)
}

// Selection Sort
func selectionSort(array [5]int) [5]int {
	for i := 0; i < len(array); i++ {
		minIndex := i

		for j := i + 1; j < len(array); j++ {
			if array[j] < array[minIndex] {
				minIndex = j
			}
		}

		if i != minIndex {
			aux := array[i]
			array[i] = array[minIndex]
			array[minIndex] = aux
		}
	}

	return array
}

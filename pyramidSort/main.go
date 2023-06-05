package main

import "fmt"

func pyramidSort(slice []int) []int {

	for i := len(slice)/2 - 1; i >= 0; i-- {
		slice = heapify(slice, i, len(slice))
	}

	for i := len(slice) - 1; i >= 1; i-- {
		slice[0], slice[i] = slice[i], slice[0]
		slice = heapify(slice, 0, i)
	}

	return slice
}

func heapify(slice []int, i int, sliceLen int) []int {

	maxChild := 0

	for i*2+1 < sliceLen {
		if i*2+1 == sliceLen-1 {
			maxChild = i*2 + 1
		} else if slice[i*2+1] > slice[i*2+2] {
			maxChild = i*2 + 1
		} else {
			maxChild = i*2 + 2
		}

		if slice[i] < slice[maxChild] {
			slice[i], slice[maxChild] = slice[maxChild], slice[i]
			i = maxChild
		} else {
			break
		}
		fmt.Printf("Sorting ...:\t%v\n", slice)
	}

	return slice
}

func main() {
	fmt.Println(pyramidSort([]int{4, 6, 3, 7, 5, 2, 8, 1, 9}))
	fmt.Println(pyramidSort([]int{4, 6, 3, 7, 5, 2, 8, 1, 10, 9}))
}

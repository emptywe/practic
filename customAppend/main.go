package main

import (
	"fmt"
)

func cAppend[Type any](slice []Type, val ...Type) []Type {

	l := len(slice)
	total := len(slice) + len(val)
	if total > cap(slice) {

		newSize := total*2 + 1
		newSlice := make([]Type, total, newSize)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[:total]
	copy(slice[l:], val)

	fmt.Println(slice)
	return slice
}

func main() {
	arr := make([]int, 1, 2)
	//arr[0] = 1
	//arr[1] = 2
	//arr = append(arr, 1)
	//fmt.Println(arr)
	fmt.Println("Slice:", arr)
	arr = cAppend(arr, 1)
	fmt.Println("Slice:", arr)
	arr = cAppend(arr, 1)
	fmt.Println("Slice:", arr)
	arr = cAppend(arr, 1, 2)
	fmt.Println("Slice:", arr)
	arr = cAppend(arr, 1, 2, 3)
	fmt.Println("Slice:", arr)
}

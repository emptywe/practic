package main

import "fmt"

func shakeSort(slice []int) []int {
	l := len(slice)
	if l == 0 {
		return slice
	}
	left := 0
	right := l - 1
	for left < right {
		for i := right; i > left; i-- {
			if slice[i-1] > slice[i] {
				slice[i-1], slice[i] = slice[i], slice[i-1]
			}
		}
		left++
		for i := left; i < right; i++ {
			if slice[i+1] < slice[i] {
				slice[i+1], slice[i] = slice[i], slice[i+1]
			}
		}
		right--
	}
	return slice
}

func main() {
	fmt.Println(shakeSort([]int{4, 6, 3, 7, 5, 2, 8, 1, 9}))
	fmt.Println(shakeSort([]int{4, 6, 3, 7, 5, 2, 8, 1, 10, 9}))
}

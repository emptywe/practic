package main

import "fmt"

func reverse(slice []int) []int {
	for i := 0; i < len(slice)/2; i++ {
		slice[i], slice[len(slice)-1-i] = slice[len(slice)-1-i], slice[i]
	}
	return slice
}

func main() {
	var slice = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(reverse(slice))
}

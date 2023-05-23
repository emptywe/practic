package main

import "fmt"

func removeAppearance(slice []int, k int) []int {

	countMap := make(map[int]int)
	for _, v := range slice {
		countMap[v]++
	}

	someCounter := 0

	for i := range slice {

		for countMap[slice[i+someCounter]] >= k && i+someCounter < len(slice)-1 {
			slice[i+someCounter] = 0
			someCounter++
		}
		if i+someCounter >= len(slice)-1 {
			break
		}
		slice[i], slice[i+someCounter] = slice[i+someCounter], slice[i]

	}
	if someCounter == 0 {
		someCounter--
	}
	return slice[0 : len(slice)-1-someCounter]
}

func main() {
	var slice = []int{1, 2, 2, 3, 3, 3, 4, 4, 4, 4, 5, 5, 5, 5, 5}
	slice = removeAppearance(slice, 4)
	fmt.Println(slice)
	slice = []int{1, 2, 3, 3}
	slice = removeAppearance(slice, 2)
	fmt.Println(slice)
	slice = []int{1, 2, 3, 3, 1, 1, 2, 3, 1, 1}
	slice = removeAppearance(slice, 4)
	fmt.Println(slice)
}

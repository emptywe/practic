package main

import "fmt"

func maxUnicSum(slice []int) int {

	var (
		start     int
		sum       int
		maxSum    int
		sliceKeys = make(map[int]int)
	)
	for end := 0; end < len(slice); end++ {

		sliceKeys[slice[end]] = end
		sum += slice[end]
		if len(sliceKeys) > 2 && slice[end] != slice[end-1] {
			start = sliceKeys[slice[start]] + 1
			end = start - 1
			sum = 0
			delete(sliceKeys, slice[start-1])
		}
		maxSum = max(sum, maxSum)
	}
	return maxSum
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func main() {
	var slice = []int{1}
	fmt.Println(maxUnicSum(slice))
	slice = []int{1, 2}
	fmt.Println(maxUnicSum(slice))
	slice = []int{10, 10, 10, 5, 5, 3}
	fmt.Println(maxUnicSum(slice))
	slice = []int{10, 10, 3, 5, 5}
	fmt.Println(maxUnicSum(slice))
	slice = []int{5, 5, 3, 10, 10}
	fmt.Println(maxUnicSum(slice))
	slice = []int{1, 2, 5, 5, 5, 10, 10, 10, 20}
	fmt.Println(maxUnicSum(slice))
	slice = []int{1, 2, 3, 7, 7, 12, 10, 10, 12, 12, 10}
	fmt.Println(maxUnicSum(slice))
}

package main

import "fmt"

func main() {
	slice := []int{4, 6, 3, 7, 5, 2, 8, 1, 9}
	sort(slice, 0, len(slice)-1)
	fmt.Println(slice)
	slice = []int{4, 6, 3, 7, 5, 2, 8, 1, 10, 9}
	sort(slice, 0, len(slice)-1)
	fmt.Println(slice)
}

func sort(s1 []int, start, end int) {
	if start >= end {
		return
	}

	pivot := s1[start]
	i := start + 1

	for j := start; j <= end; j++ {
		if pivot > s1[j] {
			s1[i], s1[j] = s1[j], s1[i]
			i++
		}
		fmt.Printf("Sorting ...:\t%v\n", s1)
	}

	s1[start], s1[i-1] = s1[i-1], s1[start]

	sort(s1, start, i-2)
	sort(s1, i, end)
}

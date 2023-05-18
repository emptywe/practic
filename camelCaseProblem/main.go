package main

import "fmt"

func wordsCount(s string) int {
	if len(s) == 0 {
		return 0
	}
	counter := 0
	for i := range s {
		if s[i] > 64 && s[i] < 91 {
			counter++
		}
	}
	return counter + 1
}

func main() {
	fmt.Println(wordsCount("camelCaseExample"))
	fmt.Println(wordsCount("someRandomWordsInCamelCase"))
}

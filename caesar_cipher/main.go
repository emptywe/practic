package main

import "fmt"

func encode(s string, factor uint8) string {
	if len(s) == 0 {
		return s
	}
	var result = make([]byte, len(s))
	for i := range s {
		switch {
		case s[i] > 64 && s[i] < 91:
			symbol := s[i] + factor
			if symbol > 90 {
				symbol -= 26
			}
			result[i] = symbol
		case s[i] > 96 && s[i] < 123:
			symbol := s[i] + factor
			if symbol > 122 {
				symbol -= 26
			}
			result[i] = symbol
		default:
			result[i] = s[i]
		}

	}
	return string(result)
}

func decode(s string, factor uint8) string {
	if len(s) == 0 {
		return s
	}
	var result = make([]byte, len(s))
	for i := range s {
		switch {
		case s[i] > 64 && s[i] < 91:
			symbol := s[i] - factor
			if symbol < 65 {
				symbol += 26
			}
			result[i] = symbol
		case s[i] > 96 && s[i] < 123:
			symbol := s[i] - factor
			if symbol < 97 {
				symbol += 26
			}
			result[i] = symbol
		default:
			result[i] = s[i]
		}

	}
	return string(result)
}

func main() {
	// Try this
	code := encode("What a lovely day!", 4)
	fmt.Println(code)
	fmt.Println(decode(code, 4))
}

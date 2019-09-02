package main

import "fmt"

func reverseString(s []byte) {
	if s == nil || len(s) == 1 {
		return
	}
	reverseString1(s, 0)
}

func reverseString1(s []byte, index int) {
	if len(s)-1-index <= index {
		return
	}
	s[index], s[len(s)-1-index] = s[len(s)-1-index], s[index]
	reverseString1(s, index+1)
}

func main() {
	bytes := []byte{'h', 'e', 'l', 'l', 'o'}
	reverseString(bytes)
	fmt.Println(bytes)
}

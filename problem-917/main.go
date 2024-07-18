package main

import (
	"fmt"
)

func main() {
	fmt.Println("problem 917")
	fmt.Println(reverseOnlyLetters("ab-cDZz=_123"))
}

func reverseOnlyLetters(s string) string {
	length := len(s)
	chars := []byte(s)

	return recursion(0, length-1, chars)
}

func recursion(start int, end int, chars []byte) string {
	if start >= end {
		return string(chars)
	}

	if !isLetter(chars[start]) {
		return recursion(start+1, end, chars)
	}

	if !isLetter(chars[end]) {
		return recursion(start, end-1, chars)
	}

	temp := chars[start]
	chars[start] = chars[end]
	chars[end] = temp

	return recursion(start+1, end-1, chars)
}

func isLetter(r byte) bool {
	if r >= 'A' && r <= 'Z' || r >= 'a' && r <= 'z' {
		return true
	} else {
		return false
	}
}

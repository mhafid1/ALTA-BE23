package main

import (
	"fmt"
	"strings"
)

func palindrome(input string) bool {
	input = strings.ToLower(input)
	runes := []rune(input)

	for i := 0; i < len(runes)/2; i++ {
		if runes[i] != runes[len(runes)-1-i] {
			return false
		}
	}
	return true
}

func main() {
	var text string
	fmt.Println("masukkan kata")
	fmt.Scanln(&text)

	if palindrome(text) {
		fmt.Println(text, "benar kata tersebut palindrome")
	} else {
		fmt.Println(text, "salah kata tersebut bukanlah palindrome")
	}
}

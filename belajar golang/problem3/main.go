// menentukan bilangan prima yang di input benar atau salah

package main

import (
	"fmt"
)

func primeNumber(number int) bool {
	if number <= 1 {
		return false
	}
	for i := 2; i <= number/2; i++ {
		if number%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var number int
	fmt.Println("masukkan bilangan prima: ")
	fmt.Scanln(&number)

	if primeNumber(number) {
		fmt.Println(number, "adalah benar bilangan prima")
	} else {
		fmt.Println(number, "adalah salah bukanlah bilangan prima")
	}
}

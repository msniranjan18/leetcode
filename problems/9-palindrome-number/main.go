package main

import (
	"fmt"
)

func main() {
	var num int
	for {
		fmt.Println("Enter the no:")
		fmt.Scan(&num)
		fmt.Println(reverseNumber(num))
		//fmt.Println(isPalindrome(num))
	}
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	} else if x == reverseNumber(x) {
		return true
	}
	return false
}

func reverseNumber(num int) int {
	reversed := 0
	for num != 0 {
		rem := num % 10
		num = num / 10
		reversed = reversed*10 + rem
	}
	return reversed
}

func isPalindromev2(num int) bool {
	if num < 0 {
		return false
	}

	x := num
	reversed := 0
	for x != 0 {
		reversed = 10*reversed + x%10
		x /= 10
	}

	return (reversed == num)
}

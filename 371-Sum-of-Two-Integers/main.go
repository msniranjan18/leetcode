package main

import (
	"fmt"
)

func main() {
	fmt.Println(getSum(5,5))
}

func getSum(a int, b int) int {
	for b != 0 {
		calculateSumWithoutCarry := a^b // calculate sum without carry
		carry := a&b // find the carry
		carry = carry << 1 // left sift by one so carry will be added to next digit.
		a = calculateSumWithoutCarry
		b = carry
	}
	return a
}

// Explanation
/*
a = 5
b = 5
a^b = 101 ^ 101 = 000 (this sis sum without carry)
a&b = 101 & 101 = 101
a&b << 1 = 101 << 1 = 1010
now repeate this process

a^b = 000 ^ 1010 = 1010
a&b = 000 & 1010 = 0000
a&b<<1 = 00000

nor carry is 0 that mean our addition is finished. return a
*/
package main

import (
	"fmt"
)

func main() {
	fmt.Println(plusOne([]int{9,9,9}))
}

func plusOne(digits []int) []int {
    return solution1(digits)
}

func solution1(digits []int) []int {
    carry:=1 // this is add 1
    for i:=len(digits)-1; i>=0; i-- {
        total := digits[i] + carry
        digits[i] = total%10
        carry = total / 10
    }
    if carry !=0 {
        digits = append([]int{carry}, digits...)
    }
    return digits
}
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(reverseInt32(1534236469))
}

func reverseInt32(x int) int {
	if x > math.MaxInt32 || x < math.MinInt32 {
		return 0
	}
	reversed := 0
	for x != 0 {
		reversed = reversed*10 + x%10
		x /= 10
	}
	if reversed > math.MaxInt32 || reversed < math.MinInt32 {
		return 0
	}
	return reversed
}

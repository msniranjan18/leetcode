/*
Solved on: 
30-May-26
*/
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(reverse(1534236469))
}

func reverse(x int) int {
    fmt.Println(math.MaxInt32, x ,math.MinInt32) // this is required to know the boundries
    if x > math.MaxInt32 || x < math.MinInt32 {
        return 0
    }

    rev:=0
    for x!=0 {
        digit := x%10

        if rev > math.MaxInt32/10 || rev == math.MaxInt32/10 && digit > 7 {
            return 0
        }
        if rev < math.MinInt32/10 || rev == math.MinInt32/10 && digit < -8 {
            return 0
        }

        rev = rev * 10 + digit
        x = x/10
    }
    return rev
}
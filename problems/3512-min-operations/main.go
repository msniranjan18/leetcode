package main

import ( 
	"fmt"
)

func main() {
	fmt.Println(minOperations([]int{1,2,3}, 9))
	fmt.Println(minOperations([]int{1,2,3}, 2))
}

func minOperations(nums []int, k int) int {
    sum:=0
    for _, num := range nums {
        sum+=num
    }
    return sum%k
}
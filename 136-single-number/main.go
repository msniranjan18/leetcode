package main

import (
	"fmt"
)

func main() {
	fmt.Println(singleNumber([]int{2,2,1}))
	
	fmt.Println(singleNumber([]int{3,2,3,2,0,1,1}))
}

func singleNumber(nums []int) int {
	num := 0
	for i:=0; i<len(nums); i++ {
		num ^= nums[i]
	}
	return num
}
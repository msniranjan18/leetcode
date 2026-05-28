package main

import (
	"fmt"
)

func main() {
	fmt.Println(minimumOperations([]int{1,2,3,4,2,3,3,5,7}))
}

func minimumOperations(nums []int) int {
    opCount := 0
    for !containsDistinct(nums) {
        if len(nums) > 3 {
        nums=nums[3:]
        } else {
            nums=nums[:0]
        }
        opCount += 1
    }
    return opCount
}

func containsDistinct(nums []int) bool {
    m := make(map[int]struct{})
    for i:=0; i<len(nums); i++ {
        if _,ok:=m[nums[i]];ok{
            return false
        } else {
            m[nums[i]] = struct{}{}
        }
    }
    return true
}
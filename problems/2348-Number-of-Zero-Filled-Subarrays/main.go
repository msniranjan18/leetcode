package main

import (
  "fmt"
)

func main() {
  fmt.Println(zeroFilledSubarray([]int{1,3,0,0,2,0,0,4}))
  fmt.Println(zeroFilledSubarray([]int{0,0,0,2,0,0}))
}
func zeroFilledSubarray(nums []int) int64 {
    
    subArray := func(n int) int {
        return n*(n+1)/2
    }

    res:=0
    for i:=0; i<len(nums); i++ {
            consicutiveZero:=0
            for i<len(nums)&&nums[i]==0 { // count the cconsicutibe zeros
                consicutiveZero++
                i++
            }
            res+=subArray(consicutiveZero)
    }
    return int64(res)
}

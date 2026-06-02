package main

import (
   "fmt"
)

func main() {
   fmt.Println(missingNumber([]int{0,2}))
   fmt.Println(missingNumber([]int{0,2,1,3,5}))
   fmt.Println(missingNumber([]int{1}))
}


func missingNumber(nums []int) int {
   // AP sum formula
   // sum = n/2(2a+(n-1)d)
   // where a is starting number and d is different 
   // in this case a=0 and d=1
   // so formula will be n/2(2*0+(n-1)*1) = n(n-1)/2
   sum:=0
   n := len(nums)+1
   totalSum := n*(n-1)/2
   for _,num := range nums {
    sum += num
   }
   return totalSum-sum
}

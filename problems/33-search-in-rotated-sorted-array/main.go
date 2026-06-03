package main

import (
	"fmt"
)

func main() {
    fmt.Println(search([]int{4,5,6,7,0,1,2},0))
}

func search(nums []int, target int) int {
    
    left, right := 0, len(nums)-1

    for left <= right {
        mid := left + (right - left)/2
        //fmt.Println(left, mid, right)
        if target == nums[mid] {
            return mid
        } else if nums[mid] >= nums[left] { // left sorted 
            if target >= nums[left] && target < nums[mid] {
                right = mid - 1
            } else {
                left = mid + 1
            }
        } else { // right side sorted
            if target <= nums[right] && target > nums[mid] {
                left = mid + 1
            } else {
                right = mid - 1
            }            
        }
    }
    return -1
}
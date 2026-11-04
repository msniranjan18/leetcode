// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	fmt.Println(findPeakElement([]int{1, 2, 3}))
	fmt.Println(findPeakElement([]int{3, 2, 1}))
	fmt.Println(findPeakElement([]int{3}))
}

func findPeakElement(nums []int) int {

	peak := 0

	// [1,2,3] // [3,2,1] // [1]
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] > nums[i-1] {
			peak = i
			if i < len(nums)-1 && nums[i] > nums[i+1] {
				return i
			}
		}
	}
	return peak
}

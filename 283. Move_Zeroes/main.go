/*
Given an integer array nums, move all 0's to the end of it while maintaining the relative order of the non-zero elements.

Note that you must do this in-place without making a copy of the array.

 

Example 1:

Input: nums = [0,1,0,3,12]
Output: [1,3,12,0,0]
Example 2:

Input: nums = [0]
Output: [0]
 

Constraints:

1 <= nums.length <= 104
-231 <= nums[i] <= 231 - 1
*/

package main

import "fmt"

type TestCase struct {
	GivenNumbers []int
}

func main() {
	testCases := []TestCase{
		{[]int{0, 1, 0, 3, 12}},
		{[]int{0}},
		{[]int{0, 0, 0, 0, 12}},
	}

	for _, testCase := range testCases {
		moveZeroes(testCase.GivenNumbers)
		fmt.Printf("output: %+v\n", testCase.GivenNumbers)
	}

}

func moveZeroes(nums []int) {
	i := 0
	for i < len(nums)-1 {
		if nums[i] != 0 {
			i++
			continue
		}
		j := i + 1
		for j < len(nums) && nums[j] == 0 {
			j++
		}
		if j < len(nums) {
			nums[i], nums[j] = nums[j], nums[i]
		} else {
			break
		}
		i++
	}
}

/*
Output:
output: [1 3 12 0 0]
output: [0]
output: [12 0 0 0 0]
*/

/*
Given an array nums, return true if the array was originally sorted in non-decreasing order, then rotated some number of positions (including zero). Otherwise, return false.

There may be duplicates in the original array.

Note: An array A rotated by x positions results in an array B of the same length such that A[i] == B[(i+x) % A.length], where % is the modulo operation.



Example 1:

Input: nums = [3,4,5,1,2]
Output: true
Explanation: [1,2,3,4,5] is the original sorted array.
You can rotate the array by x = 3 positions to begin on the element of value 3: [3,4,5,1,2].
Example 2:

Input: nums = [2,1,3,4]
Output: false
Explanation: There is no sorted array once rotated that can make nums.
Example 3:

Input: nums = [1,2,3]
Output: true
Explanation: [1,2,3] is the original sorted array.
You can rotate the array by x = 0 positions (i.e. no rotation) to make nums.


Constraints:

1 <= nums.length <= 100
1 <= nums[i] <= 100
*/


package main

import "fmt"

func main() {
	testcases := [][]int{
		{3, 4, 5, 1, 2},
		{2, 1, 3, 4},
		{1, 2, 3},
		{1, 1, 1},
		{1, 1, 1, 2, 2, 3, 3, 4},
		{2, 2, 3, 3, 4, 4, 1, 1, 1},
		{10, 1, 1, 10},
		{1, 3, 2},
	}
	for _, testcase := range testcases {
		fmt.Printf("Testcase:%+v, result:%v\n", testcase, check(testcase))
	}

}

func check(nums []int) bool {
	i := 1
	for i < len(nums) {
		if nums[i] >= nums[i-1] {
			i++
			continue
		} else {
			break
		}
	}
	if i == len(nums) { // this means array is alreay sorted
		return true
	}

	i++

	for i < len(nums) {
		if nums[i] >= nums[i-1] && nums[i] <= nums[0] {
			i++
			continue
		} else {
			return false
		}
	}
	return nums[len(nums)-1] <= nums[0]
}

/*
Output:
Testcase:[3 4 5 1 2], result:true
Testcase:[2 1 3 4], result:false
Testcase:[1 2 3], result:true
Testcase:[1 1 1], result:true
Testcase:[1 1 1 2 2 3 3 4], result:true
Testcase:[2 2 3 3 4 4 1 1 1], result:true
Testcase:[10 1 1 10], result:true
Testcase:[1 3 2], result:false
*/

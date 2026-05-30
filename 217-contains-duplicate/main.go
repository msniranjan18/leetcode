/*
Given an integer array nums, return true if any value appears at least twice in the array, and return false if every element is distinct.



Example 1:

Input: nums = [1,2,3,1]

Output: true

Explanation:

The element 1 occurs at the indices 0 and 3.

Example 2:

Input: nums = [1,2,3,4]

Output: false

Explanation:

All elements are distinct.

Example 3:

Input: nums = [1,1,1,3,3,4,3,2,4,2]

Output: true



Constraints:

1 <= nums.length <= 105
-109 <= nums[i] <= 109
*/

package main

import "fmt"

func main() {
	testcases := [][]int{
		{1, 2, 3, 1},
		{1, 2, 3, 4},
		{1, 1, 1, 3, 3, 4, 3, 2, 4, 2},
		{1},
	}

	for i := 0; i < len(testcases); i++ {
		fmt.Printf("testcase:%+v, result:%+v \n", testcases[i], containsDuplicate(testcases[i]))
	}

}

func containsDuplicate(nums []int) bool {
	num_map := make(map[int]int)
	for _, num := range nums {
		num_map[num]++
		if num_map[num] > 1 {
			return true
		}
	}
	return false
}

/*
Output:
testcase:[1 2 3 1], result:true 
testcase:[1 2 3 4], result:false 
testcase:[1 1 1 3 3 4 3 2 4 2], result:true 
testcase:[1], result:false 
*/

/*
Given a non-empty array of integers nums, every element appears twice except for one. Find that single one.

You must implement a solution with a linear runtime complexity and use only constant extra space.



Example 1:

Input: nums = [2,2,1]
Output: 1
Example 2:

Input: nums = [4,1,2,1,2]
Output: 4
Example 3:

Input: nums = [1]
Output: 1


Constraints:

1 <= nums.length <= 3 * 104
-3 * 104 <= nums[i] <= 3 * 104
Each element in the array appears twice except for one element which appears only once.
*/

package main

import "fmt"

func main() {
	testcases := [][]int{
		{2, 2, 1},
		{4, 1, 2, 1, 2},
		{1},
		{1, 0, 3, 0, 1},
		{2, 2, 1, 3, 3},
		{1, 2, 3, 2, 1},
		{1, 2, 2},
		{2, 1, 2},
	}

	for _, testcase := range testcases {
		fmt.Printf("testcase:%+v, result:%d \n", testcase, singleNumber(testcase))
	}

}

func singleNumber(nums []int) int {
	//bruteforce method
	// i:=0
	// for ;i<len(nums);i++{
	//     if nums[i] == 0 {
	//         continue
	//     }
	//     for j:=i+1; j<len(nums);j++{
	//         if nums[i] == nums[j] {
	//             nums[i] = 0
	//             nums[j] = 0
	//             break
	//         }
	//         if j+1 == len(nums) {
	//             return nums[i]
	//         }
	//     }
	// }
	// return nums[i-1]

	//map approach
	// num_map := make(map[int]int)
	// for i:=0;i<len(nums);i++ {
	//     num_map[nums[i]] += 1
	// }
	// for key, val := range num_map {
	//     if val == 1 {
	//         return key
	//     }
	// }
	// return 0

	//xor approach
	var result int
	for i := 0; i < len(nums); i++ {
		result ^= nums[i]
	}
	return result
}

/*
Output:
testcase:[2 2 1], result:1 
testcase:[4 1 2 1 2], result:4 
testcase:[1], result:1 
testcase:[1 0 3 0 1], result:3 
testcase:[2 2 1 3 3], result:1 
testcase:[1 2 3 2 1], result:3 
testcase:[1 2 2], result:1 
testcase:[2 1 2], result:1 
*/

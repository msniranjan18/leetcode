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

// func moveZeroes(nums []int) {
// 	i := 0
// 	for i < len(nums)-1 {
// 		if nums[i] != 0 {
// 			i++
// 			continue
// 		}
// 		j := i + 1
// 		for j < len(nums) && nums[j] == 0 {
// 			j++
// 		}
// 		if j < len(nums) {
// 			nums[i], nums[j] = nums[j], nums[i]
// 		} else {
// 			break
// 		}
// 		i++
// 	}
// }

/*
the above code is not optimal:
1. Unnecessary i++ after swapping
After swapping nums[i] and nums[j], you increment i. However, the value at index i might still be 0 after the swap, which causes an unnecessary iteration to check again.
Example: nums = [0, 1]. After swapping, nums becomes [1, 0], but i is incremented to 1, skipping over the correct element.
Fix: Increment i only when a non-zero element is encountered, or the loop will needlessly recheck values.

2. Nested while loop to find j is inefficient
The inner loop for j<len(nums) && nums[j]==0 { j++ } is scanning for the next non-zero value, which can result in multiple iterations for each zero. This makes the algorithm less efficient for large input arrays with many consecutive zeroes.
Optimization: Use a single pointer to keep track of the position for swapping rather than searching repeatedly.

3. Loop condition i < len(nums) - 1 can cause issues
The condition i < len(nums) - 1 prevents the last element from being processed in some cases.
Example: nums = [0, 1]. The last element (index 1) won't be checked, leaving a zero in its place.
Fix: Change the loop condition to i < len(nums).

4. Unnecessary swaps
The algorithm swaps elements even when nums[i] and nums[j] are already in the correct positions. This results in unnecessary writes and degrades performance.
Fix: Only swap when nums[i] == 0 and nums[j] != 0.
*/


// Optimal code:
func moveZeroes(nums []int)  {
    j:=0
    for i:=0; i<len(nums);i++ {
        if nums[i]!=0 {
            nums[j] = nums[i]
            j++
        }
    }
    // fill the remaining zeros
    for j<len(nums) {
        nums[j] = 0
        j++
    }
}
/*
Output:
output: [1 3 12 0 0]
output: [0]
output: [12 0 0 0 0]
*/

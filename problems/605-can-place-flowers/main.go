/*
You have a long flowerbed in which some of the plots are planted, and some are not. However, flowers cannot be planted in adjacent plots.

Given an integer array flowerbed containing 0's and 1's, where 0 means empty and 1 means not empty, and an integer n, return true if n new flowers can be planted in the flowerbed without violating the no-adjacent-flowers rule and false otherwise.



Example 1:

Input: flowerbed = [1,0,0,0,1], n = 1
Output: true
Example 2:

Input: flowerbed = [1,0,0,0,1], n = 2
Output: false


Constraints:

1 <= flowerbed.length <= 2 * 104
flowerbed[i] is 0 or 1.
There are no two adjacent flowers in flowerbed.
0 <= n <= flowerbed.length
*/

package main

import "fmt"

func main() {
	testcases := [][]int{
		{1, 0, 0, 0, 0, 1}, {2},
		{1, 0, 0, 0, 1, 0, 0}, {1},
		{0, 0, 1, 0, 0}, {2},
		{0, 0, 1, 0, 0, 1}, {2},
		{0, 0, 1, 0, 0, 0, 1}, {2},
		{0}, {1},
		{1}, {1},
		{0, 1}, {1},
		{1, 0, 0, 1}, {1},
		{0}, {0},
	}

	for i := 0; i < len(testcases); i += 2 {
		// Copy the elements from the original slice to the new slice
		flowerbed := make([]int, len(testcases[i]))
		copy(flowerbed, testcases[i])

		/*
		Key Points:
			    Slices share the underlying array if directly assigned, so use copy to create independent slices.
			    Arrays are value types in Go, so assigning an array to a new variable creates a copy automatically.
			      // Original array
			          original := [5]int{1, 2, 3, 4, 5}
			      // Copy the array
			          copied := original

			    Directly assigning slices like above array, will cause issue:
			    newSlice := originalSlice
			    Will only copy the slice header (length, capacity, and pointer to the underlying array).
			    Both newSlice and originalSlice will reference the same data, so modifying one affects the other.
			    To avoid this, always use copy for independent slices.
			      // Original slice
			          original := []int{1, 2, 3, 4, 5}
			      // Create a new slice with the same length as the original
			          copied := make([]int, len(original))
			          // Copy the elements from the original slice to the new slice
			          copy(copied, original)
		*/

		fmt.Printf("testcase: flowerbed=%+v, extraflower=%+v, canBePlace:%+v, \tflowerbedAfterPlacingExtraFlower:%+v \n", testcases[i], testcases[i+1], canPlaceFlowers(flowerbed, testcases[i+1][0]), flowerbed)
	}

}

func canPlaceFlowers(flowerbed []int, n int) bool {
	if n == 0 {
		return true
	}
	for i := 0; i < len(flowerbed) && n > 0; i++ {
		if flowerbed[i] == 0 && (i == 0 || flowerbed[i-1] == 0) && (i == len(flowerbed)-1 || flowerbed[i+1] == 0) {
			// flower can be place
			flowerbed[i] = 1
			n--
		}
	}
	if n != 0 {
		return false
	}
	return true
}

/*
Output:
testcase: flowerbed=[1 0 0 0 0 1], extraflower=[2], canBePlace:false, 	flowerbedAfterPlacingExtraFlower:[1 0 1 0 0 1]
testcase: flowerbed=[1 0 0 0 1 0 0], extraflower=[1], canBePlace:true, 	flowerbedAfterPlacingExtraFlower:[1 0 1 0 1 0 0]
testcase: flowerbed=[0 0 1 0 0], extraflower=[2], canBePlace:true, 	flowerbedAfterPlacingExtraFlower:[1 0 1 0 1]
testcase: flowerbed=[0 0 1 0 0 1], extraflower=[2], canBePlace:false, 	flowerbedAfterPlacingExtraFlower:[1 0 1 0 0 1]
testcase: flowerbed=[0 0 1 0 0 0 1], extraflower=[2], canBePlace:true, 	flowerbedAfterPlacingExtraFlower:[1 0 1 0 1 0 1]
testcase: flowerbed=[0], extraflower=[1], canBePlace:true, 	flowerbedAfterPlacingExtraFlower:[1]
testcase: flowerbed=[1], extraflower=[1], canBePlace:false, 	flowerbedAfterPlacingExtraFlower:[1]
testcase: flowerbed=[0 1], extraflower=[1], canBePlace:false, 	flowerbedAfterPlacingExtraFlower:[0 1]
testcase: flowerbed=[1 0 0 1], extraflower=[1], canBePlace:false, 	flowerbedAfterPlacingExtraFlower:[1 0 0 1]
testcase: flowerbed=[0], extraflower=[0], canBePlace:true, 	flowerbedAfterPlacingExtraFlower:[0]
*/

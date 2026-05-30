/*
There are n kids with candies. You are given an integer array candies, where each candies[i] represents the number of candies the ith kid has, and an integer extraCandies, denoting the number of extra candies that you have.

Return a boolean array result of length n, where result[i] is true if, after giving the ith kid all the extraCandies, they will have the greatest number of candies among all the kids, or false otherwise.

Note that multiple kids can have the greatest number of candies.



Example 1:

Input: candies = [2,3,5,1,3], extraCandies = 3
Output: [true,true,true,false,true]
Explanation: If you give all extraCandies to:
- Kid 1, they will have 2 + 3 = 5 candies, which is the greatest among the kids.
- Kid 2, they will have 3 + 3 = 6 candies, which is the greatest among the kids.
- Kid 3, they will have 5 + 3 = 8 candies, which is the greatest among the kids.
- Kid 4, they will have 1 + 3 = 4 candies, which is not the greatest among the kids.
- Kid 5, they will have 3 + 3 = 6 candies, which is the greatest among the kids.
Example 2:

Input: candies = [4,2,1,1,2], extraCandies = 1
Output: [true,false,false,false,false]
Explanation: There is only 1 extra candy.
Kid 1 will always have the greatest number of candies, even if a different kid is given the extra candy.
Example 3:

Input: candies = [12,1,12], extraCandies = 10
Output: [true,false,true]


Constraints:

n == candies.length
2 <= n <= 100
1 <= candies[i] <= 100
1 <= extraCandies <= 50
*/

package main

import "fmt"

func main() {
	testcases := [][]int{
		{2, 3, 5, 1, 3}, {3},
		{4, 2, 1, 1, 2}, {1},
		{12, 1, 12}, {10},
	}

	for i := 0; i < len(testcases); i += 2 {
		fmt.Printf("testcase: candies=%+v, extraCandies=%+v, result:%+v \n", testcases[i], testcases[i+1], kidsWithCandies(testcases[i], testcases[i+1][0]))
	}

}

func kidsWithCandies(candies []int, extraCandies int) []bool {
	max := getMax(candies)
	var result []bool
	for i := 0; i < len(candies); i++ {
		if (candies[i] + extraCandies) >= max {
			result = append(result, true)
		} else {
			result = append(result, false)
		}
	}
	return result
}

func getMax(candies []int) int {
	var max int
	if len(candies) > 0 {
		max = candies[0]
	}
	for i := 1; i < len(candies); i++ {
		if candies[i] > max {
			max = candies[i]
		}
	}
	return max
}

/*
Output:
testcase: candies=[2 3 5 1 3], extraCandies=[3], result:[true true true false true] 
testcase: candies=[4 2 1 1 2], extraCandies=[1], result:[true false false false false] 
testcase: candies=[12 1 12], extraCandies=[10], result:[true false true] 
*/

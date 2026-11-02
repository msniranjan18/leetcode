/*
2300. Successful Pairs of Spells and Potions
Solved
Medium
Topics
premium lock icon
Companies
Hint
You are given two positive integer arrays spells and potions, of length n and m respectively, where spells[i] represents the strength of the ith spell and potions[j] represents the strength of the jth potion.

You are also given an integer success. A spell and potion pair is considered successful if the product of their strengths is at least success.

Return an integer array pairs of length n where pairs[i] is the number of potions that will form a successful pair with the ith spell.

 

Example 1:

Input: spells = [5,1,3], potions = [1,2,3,4,5], success = 7
Output: [4,0,3]
Explanation:
- 0th spell: 5 * [1,2,3,4,5] = [5,10,15,20,25]. 4 pairs are successful.
- 1st spell: 1 * [1,2,3,4,5] = [1,2,3,4,5]. 0 pairs are successful.
- 2nd spell: 3 * [1,2,3,4,5] = [3,6,9,12,15]. 3 pairs are successful.
Thus, [4,0,3] is returned.
Example 2:

Input: spells = [3,1,2], potions = [8,5,8], success = 16
Output: [2,0,2]
Explanation:
- 0th spell: 3 * [8,5,8] = [24,15,24]. 2 pairs are successful.
- 1st spell: 1 * [8,5,8] = [8,5,8]. 0 pairs are successful. 
- 2nd spell: 2 * [8,5,8] = [16,10,16]. 2 pairs are successful. 
Thus, [2,0,2] is returned.
 

Constraints:

n == spells.length
m == potions.length
1 <= n, m <= 105
1 <= spells[i], potions[i] <= 105
1 <= success <= 1010
*/

package main

import (
	"fmt"
	"slices"
)

func main() {
	testcases := [][]int{
		{5, 1, 3}, {1, 2, 3, 4, 5}, {7}, {4, 0, 3},
		{3, 1, 2}, {8, 5, 8}, {16}, {2, 0, 2},
	}
	for i := 0; i < len(testcases); i = i + 4 {
		fmt.Printf("testcase: input=[spells:%+v potions=%+v], expacted=%+v, result:%+v \n", testcases[i], testcases[i+1], testcases[i+2], successfulPairs(testcases[i], testcases[i+1], int64(testcases[i+2][0])))
	}

}

func successfulPairs(spells []int, potions []int, success int64) []int {
	return sol2(spells, potions, success)
}

func sol2(spells []int, potions []int, success int64) []int { // binary search
	res := make([]int, len(spells))

	//slices.Sort(spells)
	slices.Sort(potions)

	for s := 0; s < len(spells); s++ {
		//count:=0
		// binary serach on the potions array
		l := 0
		r := len(potions) - 1
		m := 0
		idx := len(potions) // incase if there are no successful combination then res should have zero
		for l <= r {
			m = l + (r-l)/2
			prod := spells[s] * potions[m]
			if int64(prod) >= success {
				r = m - 1
				idx = m
			} else {
				l = m + 1
			}
		}
		res[s] = len(potions) - idx // all other bigger potion will be successfull too
	}
	return res
}

func sol1(spells []int, potions []int, success int64) []int { // TTE // incorrect solution.
	res := make([]int, len(spells))

	//slices.Sort(sp ells)
	slices.Sort(potions)

	for s := 0; s < len(spells); s++ {
		count := 0
		for p := 0; p < len(potions); p++ {
			prod := spells[s] * potions[p]
			if int64(prod) >= success {
				count = count + len(potions) - p // all other bigger potion will be successfull too
				break
			}
		}
		res[s] = count
	}
	return res
}
/*
testcase: input=[spells:[5 1 3] potions=[1 2 3 4 5]], expacted=[7], result:[4 0 3] 
testcase: input=[spells:[3 1 2] potions=[5 8 8]], expacted=[16], result:[2 0 2] 
*/

/*
Write a function to find the longest common prefix string amongst an array of strings.

If there is no common prefix, return an empty string "".



Example 1:

Input: strs = ["flower","flow","flight"]
Output: "fl"
Example 2:

Input: strs = ["dog","racecar","car"]
Output: ""
Explanation: There is no common prefix among the input strings.


Constraints:

1 <= strs.length <= 200
0 <= strs[i].length <= 200
strs[i] consists of only lowercase English letters.
*/

package main

import "fmt"

func main() {
	testcases := [][]string{
		{"flower", "flow", "flight"},
		{"flower", "g"},
		{"flower"},
		{"f", "flower", "flower"},
		{"dog", "racecar", "car"},
		{"flower", "flow"},
		{"flower", "flower"},
	}

	for _, testcase := range testcases {
		fmt.Printf("testcase:%+v, result:%+v \n", testcase, longestCommonPrefix(testcase))
	}

}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	for i := 0; i < len(strs[0]); i++ {
		char := strs[0][i]
		for j := 1; j < len(strs); j++ {
			if i >= len(strs[j]) || char != strs[j][i] {
				return strs[0][0:i]
			}
		}
	}
	return strs[0]
}

/*
Output:
testcase:[flower flow flight], result:fl 
testcase:[flower g], result: 
testcase:[flower], result:flower 
testcase:[f flower flower], result:f 
testcase:[dog racecar car], result: 
testcase:[flower flow], result:flow 
testcase:[flower flower], result:flower 
*/

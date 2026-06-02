/*
You are given an array of strings words and a string pref.

Return the number of strings in words that contain pref as a prefix.

A prefix of a string s is any leading contiguous substring of s.



Example 1:

Input: words = ["pay","attention","practice","attend"], pref = "at"
Output: 2
Explanation: The 2 strings that contain "at" as a prefix are: "attention" and "attend".
Example 2:

Input: words = ["leetcode","win","loops","success"], pref = "code"
Output: 0
Explanation: There are no strings that contain "code" as a prefix.


Constraints:

1 <= words.length <= 100
1 <= words[i].length, pref.length <= 100
words[i] and pref consist of lowercase English letters.
*/

package main

import "fmt"

func main() {
	testcases := [][]string{
		{"pay", "attention", "practice", "attend"}, {"at"},
		{"leetcode", "win", "loops", "success"}, {"code"},
		{"pay", "aention", "aractice", "attend"}, {"at"},
		{"leetcode"}, {"c"},
		{"leetcode"}, {"l"},
		{"leet", "leetcode"},
		{"leetcode", "leetcode"},
	}

	for i := 0; i < len(testcases); i = i + 2 {
		fmt.Printf("testcase:%+v, prefix:%+v result:%+v \n", testcases[i], testcases[i+1], prefixCount(testcases[i], testcases[i+1][0]))
	}

}
func prefixCount(words []string, pref string) int {
	count := 0
	for i := 0; i < len(words); i++ {
		j := 0
		for ; j < len(pref); j++ {
			if len(pref) > len(words[i]) || pref[j] != words[i][j] {
				break
			}
		}
		if j >= len(pref) {
			count++
		}
	}
	return count
}

/*
Output:
testcase:[pay attention practice attend], prefix:[at] result:2 
testcase:[leetcode win loops success], prefix:[code] result:0 
testcase:[pay aention aractice attend], prefix:[at] result:1 
testcase:[leetcode], prefix:[c] result:0 
testcase:[leetcode], prefix:[l] result:1 
testcase:[leet leetcode], prefix:[leetcode leetcode] result:1 
*/

/*
You are given a string s and a list dictionary[] of words. Your task is to determine whether the string s can be formed by concatenating one or more words from the dictionary[].

Note: From dictionary[], any word can be taken any number of times and in any order.

Examples :

Input: s = "ilike", dictionary[] = ["i", "like", "gfg"]
Output: true
Explanation: s can be breakdown as "i like".
Input: s = "ilikegfg", dictionary = ["i", "like", "man", "india", "gfg"]
Output: true
Explanation: s can be breakdown as "i like gfg".
Input: s = "ilikemangoes", dictionary = ["i", "like", "man", "india", "gfg"]
Output: false
Explanation: s cannot be formed using dictionary[] words.
Constraints:
1 ≤ s.size() ≤ 3000
1 ≤ dictionary.size() ≤ 1000
1 ≤ dictionary[i].size() ≤ 100
*/


/*
Instead of a greedy two-pointer approach, Dynamic Programming (DP) is the best approach.

Idea
Use a boolean DP array (dp[i]), where:

dp[i] = true if s[0:i] can be formed using words in dict.

dp[0] = true (empty string is always valid).

Iterate over the string, checking substrings to see if they exist in dict.

Use dp[j] to check if s[j:i] can extend a valid substring.
*/

package main

import (
	"fmt"
)

func main() {
	fmt.Println(wordBreak("ilikemangoes", []string{"i", "like", "man", "india", "gfg"}))
	fmt.Println(wordBreak("ilikegfg", []string{"i", "like", "man", "india", "gfg"}))
	fmt.Println(wordBreak("catsanddogs", []string{"cat", "cats", "and", "dog", "dogs"}))
}

func wordBreak(s string, dict []string) bool {
	mapDict := make(map[string]bool, len(dict))

	// fill the map
	for _, v := range dict {
		mapDict[v] = true
	}

	// use dynamic programming
	length := len(s)
	dp := make([]bool, length+1)
	dp[0] = true // this is for empty string.

	for i := 1; i <= length; i++ {
		for j := 0; j < i; j++ {
			if dp[j] && mapDict[s[j:i]] {
				dp[i] = true
			}
		}
	}

	fmt.Println(dp)
	return dp[length]
}

/*
Time Complexity Analysis
O(n²): Double loop (i from 1 to n, j from 0 to i).

Space Complexity: O(n + m) (O(n) for dp[], O(m) for dictionary storage).
*/

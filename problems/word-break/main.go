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

// more efficient code; use max length of the biggest word from the dictionary.
/*
Precompute max word length (maxLen) to limit the number of substrings checked.

Use a boolean DP array (dp[i]) where:

dp[i] = true → The substring s[0:i] can be formed using dictionary words.

dp[0] = true (empty string is always breakable).

Iterate i from 1 to n:

Check only substrings of length at most maxLen (reduces time complexity).

If dp[j] is true and s[j:i] exists in wordSet, mark dp[i] = true.
*/

/*
Time Complexity Analysis
Building wordSet → O(D), where D is the total words in dictionary.

DP Traversal → O(N * maxLen), much better than brute-force O(N²).

Overall Complexity → O(N + D), which is efficient.
*/

/*
i	j Range	s[j:i]	dp[j]	Exists in wordSet?	dp[i]
1	[0]	"c"	❌	❌	❌
2	[0,1]	"ca", "a"	❌❌	❌❌	❌
3	[0,1,2]	"cat"	✅	✅	✅
4	[1,2,3]	"cats"	✅	✅	✅
5	[2,3,4]	"s", "sa", "san"	❌	❌	❌
6	[3,4,5]	"sand"	✅	✅	✅
7	[4,5,6]	"and"	✅	✅	✅
8	[5,6,7]	"d"	❌	❌	❌
9	[6,7,8]	"dog"	✅	✅	✅
10	[7,8,9]	"dogs"	✅	✅	✅
✅ Final dp[10] = True, so the function returns true!
*/

/*
Why This is Optimized
✅ Avoids checking unnecessary substrings (unlike brute-force approaches).
✅ Uses a hash set (wordSet) for O(1) lookups.
✅ Limits the inner loop using maxLen, making it faster than O(N²) solutions.
*/


func wordBreak(s string, dict []string) bool {
	mapDict := make(map[string]bool, len(dict))

	// fill the map and find the max length of words
	var maxWordLength int
	for _, v := range dict {
		mapDict[v] = true
		if len(v) > maxWordLength { // finding the maximum lenght of the word so we'll check only till that length
			maxWordLength = len(v)
		}
	}

	// use dynamic programming
	length := len(s)
	dp := make([]bool, length+1)
	dp[0] = true // this is for empty string.

	for i := 1; i <= length; i++ {
		for j := max(0, i-maxWordLength); j < i; j++ {
			if dp[j] && mapDict[s[j:i]] {
				dp[i] = true
			}
		}
	}

	fmt.Println(dp)
	return dp[length]
}


/*
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
*/

/*
Time Complexity Analysis
O(n²): Double loop (i from 1 to n, j from 0 to i).

Space Complexity: O(n + m) (O(n) for dp[], O(m) for dictionary storage).
*/

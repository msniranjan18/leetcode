/*
You are given two strings word1 and word2. Merge the strings by adding letters in alternating order, starting with word1. If a string is longer than the other, append the additional letters onto the end of the merged string.

Return the merged string.

 

Example 1:

Input: word1 = "abc", word2 = "pqr"
Output: "apbqcr"
Explanation: The merged string will be merged as so:
word1:  a   b   c
word2:    p   q   r
merged: a p b q c r
Example 2:

Input: word1 = "ab", word2 = "pqrs"
Output: "apbqrs"
Explanation: Notice that as word2 is longer, "rs" is appended to the end.
word1:  a   b 
word2:    p   q   r   s
merged: a p b q   r   s
Example 3:

Input: word1 = "abcd", word2 = "pq"
Output: "apbqcd"
Explanation: Notice that as word1 is longer, "cd" is appended to the end.
word1:  a   b   c   d
word2:    p   q 
merged: a p b q c   d
 

Constraints:

1 <= word1.length, word2.length <= 100
word1 and word2 consist of lowercase English letters.
*/

package main

import "fmt"

func main() {
	testcases := [][]string{
		{"abc"}, {"pqr"},
		{"a"}, {"pqr"},
		{"abc"}, {"p"},
		{""}, {"pqr"},
		{"abc"}, {""},
	}

	for i := 0; i < len(testcases); i += 2 {
		fmt.Printf("testcase: word1=%+v, word2=%+v, result:%+v \n", testcases[i], testcases[i+1], mergeAlternately(testcases[i][0], testcases[i+1][0]))
	}

}

func mergeAlternately(word1 string, word2 string) string {
	var result []byte
	i := 0
	for ; i < len(word1) && i < len(word2); i++ {
		result = append(result, word1[i], word2[i])
	}
	for ; i < len(word1); i++ {
		result = append(result, word1[i])
	}
	for ; i < len(word2); i++ {
		result = append(result, word2[i])
	}
	return string(result)
}

/*
Output:
testcase: word1=[abc], word2=[pqr], result:apbqcr 
testcase: word1=[a], word2=[pqr], result:apqr 
testcase: word1=[abc], word2=[p], result:apbc 
testcase: word1=[], word2=[pqr], result:pqr 
testcase: word1=[abc], word2=[], result:abc 
*/

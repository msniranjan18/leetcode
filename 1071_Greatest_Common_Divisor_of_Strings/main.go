/*
For two strings s and t, we say "t divides s" if and only if s = t + t + t + ... + t + t (i.e., t is concatenated with itself one or more times).

Given two strings str1 and str2, return the largest string x such that x divides both str1 and str2.



Example 1:

Input: str1 = "ABCABC", str2 = "ABC"
Output: "ABC"
Example 2:

Input: str1 = "ABABAB", str2 = "ABAB"
Output: "AB"
Example 3:

Input: str1 = "LEET", str2 = "CODE"
Output: ""


Constraints:

1 <= str1.length, str2.length <= 1000
str1 and str2 consist of English uppercase letters.
*/

package main

import "fmt"

func main() {
	testcases := [][]string{
		{"ABCABC"}, {"ABC"},
		{"A"}, {"A"},
		{"A"}, {"B"},
		{"ABCDABCDABCD"}, {"ABCD"},
		{"AAA"}, {"A"},
		{"ABABABABABABABABABABABABABABABABABABABAB"}, {"ABABABABABABABABAB"},
	}

	for i := 0; i < len(testcases); i += 2 {
		fmt.Printf("testcase: word1=%+v, word2=%+v, result:%+v \n", testcases[i], testcases[i+1], gcdOfStrings(testcases[i][0], testcases[i+1][0]))
	}

}

func gcdOfStrings(str1 string, str2 string) string {
	if (str1 + str2) != (str2 + str1) {
		return ""
	}
	return str2[:gcd(len(str1), len(str2))]
}

func gcd(num1, num2 int) int {
	for num2 != 0 {
		temp := num2
		num2 = num1 % num2
		num1 = temp
	}
	return num1
}

/*
Output:
testcase: word1=[ABCABC], word2=[ABC], result:ABC 
testcase: word1=[A], word2=[A], result:A 
testcase: word1=[A], word2=[B], result: 
testcase: word1=[ABCDABCDABCD], word2=[ABCD], result:ABCD 
testcase: word1=[AAA], word2=[A], result:A 
testcase: word1=[ABABABABABABABABABABABABABABABABABABABAB], word2=[ABABABABABABABABAB], result:AB 
*/

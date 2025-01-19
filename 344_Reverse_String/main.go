/*
Write a function that reverses a string. The input string is given as an array of characters s.

You must do this by modifying the input array in-place with O(1) extra memory.



Example 1:

Input: s = ["h","e","l","l","o"]
Output: ["o","l","l","e","h"]
Example 2:

Input: s = ["H","a","n","n","a","h"]
Output: ["h","a","n","n","a","H"]


Constraints:

1 <= s.length <= 105
s[i] is a printable ascii character.
*/

package main

import "fmt"

type TestCase struct {
	GivenString string
	Expected    string
}

func main() {
	testCases := []TestCase{
		{"hello", "olleh"},
		{"Leetcode", "edocteeL"},
		{"l", "l"},
	}

	for _, testCase := range testCases {
		s := []byte(testCase.GivenString)
		reverseString(s)
		if string(s) == testCase.Expected {
			fmt.Printf("Test Case: [GivenString=%+v] PASSED: output: %s, expected %s\n", testCase.GivenString, string(s), testCase.Expected)
		} else {
			fmt.Printf("Test Case: [GivenString=%+v] FAILED: output: %s, expected %s\n", testCase.GivenString, string(s), testCase.Expected)
		}
	}

}

func reverseString(s []byte) {
	i, j := 0, len(s)-1
	for i < j {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
}

/*
Output:
Test Case: [GivenString=hello] PASSED: output: olleh, expected olleh
Test Case: [GivenString=Leetcode] PASSED: output: edocteeL, expected edocteeL
Test Case: [GivenString=l] PASSED: output: l, expected l
*/

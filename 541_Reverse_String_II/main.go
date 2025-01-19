/*
Given a string s and an integer k, reverse the first k characters for every 2k characters counting from the start of the string.

If there are fewer than k characters left, reverse all of them. If there are less than 2k but greater than or equal to k characters, then reverse the first k characters and leave the other as original.



Example 1:

Input: s = "abcdefg", k = 2
Output: "bacdfeg"
Example 2:

Input: s = "abcd", k = 2
Output: "bacd"


Constraints:

1 <= s.length <= 104
s consists of only lowercase English letters.
1 <= k <= 104
*/

package main

import "fmt"

type TestCase struct {
	GivenString string
	k           int
	Expected    string
}

func main() {
	testCases := []TestCase{
		{"abcdefg", 2, "bacdfeg"},
		{"abcd", 2, "bacd"},
		{"abcdefgh", 3, "cbadefhg"},
		{"abcdefghijklm", 4, "dcbaefghlkjim"},
		{"abcd", 5, "dcba"},
		{"krmyfshbspcgtesxnnljhfursyissjnsocgdhgfxubewllxzqhpasguvlrxtkgatzfybprfmmfithphckksnvjkcvnsqgsgosfxc", 20, "jlnnxsetgcpsbhsfymrkhfursyissjnsocgdhgfxtxrlvugsaphqzxllwebukgatzfybprfmmfithphccxfsogsgqsnvckjvnskk"},
	}

	for _, testCase := range testCases {
		result := reverseStr(testCase.GivenString, testCase.k)
		if result == testCase.Expected {
			fmt.Printf("Test Case: [GivenString=%+v] PASSED: output: %s, expected %s\n", testCase.GivenString, result, testCase.Expected)
		} else {
			fmt.Printf("Test Case: [GivenString=%+v] FAILED: output: %s, expected %s\n", testCase.GivenString, result, testCase.Expected)
		}
	}

}

func reverseStr(s string, k int) string {
	str := []byte(s)
	i := 0
	if k == 1 {
		return s
	}
	for i < len(str) {
		if len(str)-i < 2*k {
			if len(str)-i >= k {
				reverse(str[i : i+k])
				return string(str)
			}
			if len(str)-i < k {
				reverse(str[i:])
				return string(str)
			}
		}
		if k < len(str) {
			reverse(str[i : i+k])
			i += 2 * k
		}
	}
	return string(str)
}

func reverse(s []byte) {
	i, j := 0, len(s)-1
	for i < j {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
}

/*
Output:
Test Case: [GivenString=abcdefg] PASSED: output: bacdfeg, expected bacdfeg
Test Case: [GivenString=abcd] PASSED: output: bacd, expected bacd
Test Case: [GivenString=abcdefgh] PASSED: output: cbadefhg, expected cbadefhg
Test Case: [GivenString=abcdefghijklm] PASSED: output: dcbaefghlkjim, expected dcbaefghlkjim
Test Case: [GivenString=abcd] PASSED: output: dcba, expected dcba
Test Case: [GivenString=krmyfshbspcgtesxnnljhfursyissjnsocgdhgfxubewllxzqhpasguvlrxtkgatzfybprfmmfithphckksnvjkcvnsqgsgosfxc] PASSED: output: jlnnxsetgcpsbhsfymrkhfursyissjnsocgdhgfxtxrlvugsaphqzxllwebukgatzfybprfmmfithphccxfsogsgqsnvckjvnskk, expected jlnnxsetgcpsbhsfymrkhfursyissjnsocgdhgfxtxrlvugsaphqzxllwebukgatzfybprfmmfithphccxfsogsgqsnvckjvnskk
*/

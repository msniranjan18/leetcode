/*
Given a string s, reverse only all the vowels in the string and return it.

The vowels are 'a', 'e', 'i', 'o', and 'u', and they can appear in both lower and upper cases, more than once.



Example 1:

Input: s = "IceCreAm"

Output: "AceCreIm"

Explanation:

The vowels in s are ['I', 'e', 'e', 'A']. On reversing the vowels, s becomes "AceCreIm".

Example 2:

Input: s = "leetcode"

Output: "leotcede"

Constraints:

1 <= s.length <= 3 * 105
s consist of printable ASCII characters.
*/

package main

import "fmt"

type TestCase struct {
	GivenString string
	Expected    string
}

var vowels = map[byte]bool{
	'a': true, 'A': true,
	'e': true, 'E': true,
	'i': true, 'I': true,
	'o': true, 'O': true,
	'u': true, 'U': true,
}

func main() {
	testCases := []TestCase{
		{"IceCreAm", "AceCreIm"},
		{"leetcode", "leotcede"},
		{"l", "l"},
	}

	for _, testCase := range testCases {
		result := reverseVowels(testCase.GivenString)
		if result == testCase.Expected {
			fmt.Printf("Test Case: [GivenString=%+v] PASSED: output: %s, expected %s\n", testCase.GivenString, result, testCase.Expected)
		} else {
			fmt.Printf("Test Case: [GivenString=%+v] FAILED: output: %s, expected %s\n", testCase.GivenString, result, testCase.Expected)
		}
	}

}

func reverseVowels(s string) string {
	str := []byte(s)
	i, j := 0, len(str)-1

	for i < j {
		// Move i forward if not a vowel
		for i < j && !isVowel(str[i]) {
			i++
		}
		// Move j backward if not a vowel
		for i < j && !isVowel(str[j]) {
			j--
		}
		// Swap vowels
		if i < j {
			str[i], str[j] = str[j], str[i]
			i++
			j--
		}
	}

	return string(str)
}

// Helper function to check if a character is a vowel
func isVowel(c byte) bool {
	return vowels[c]
}

/*
Output:
Test Case: [GivenString=IceCreAm] PASSED: output: AceCreIm, expected AceCreIm
Test Case: [GivenString=leetcode] PASSED: output: leotcede, expected leotcede
Test Case: [GivenString=l] PASSED: output: l, expected l
*/

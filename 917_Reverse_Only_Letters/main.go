/*
Given a string s, reverse the string according to the following rules:

All the characters that are not English letters remain in the same position.
All the English letters (lowercase or uppercase) should be reversed.
Return s after reversing it.

 

Example 1:

Input: s = "ab-cd"
Output: "dc-ba"
Example 2:

Input: s = "a-bC-dEf-ghIj"
Output: "j-Ih-gfE-dCba"
Example 3:

Input: s = "Test1ng-Leet=code-Q!"
Output: "Qedo1ct-eeLg=ntse-T!"
 

Constraints:

1 <= s.length <= 100
s consists of characters with ASCII values in the range [33, 122].
s does not contain '\"' or '\\'.
*/

package main

import "fmt"

type TestCase struct {
	GivenString string
	Expected    string
}

func main() {
	testCases := []TestCase{
		{"ab-cd", "dc-ba"},
		{"a-bC-dEf-ghIj", "j-Ih-gfE-dCba"},
		{"Test1ng-Leet=code-Q!", "Qedo1ct-eeLg=ntse-T!"},
		{"7_28]", "7_28]"},
		{"-ab-", "-ba-"},
	}

	for _, testCase := range testCases {
		result := reverseOnlyLetters(testCase.GivenString)
		if result == testCase.Expected {
			fmt.Printf("Test Case: [GivenString=%+v] PASSED: output: %s, expected %s\n", testCase.GivenString, result, testCase.Expected)
		} else {
			fmt.Printf("Test Case: [GivenString=%+v] FAILED: output: %s, expected %s\n", testCase.GivenString, result, testCase.Expected)
		}
	}

}

// func reverseOnlyLetters(s string) string {
// 	i, j := 0, len(s)-1
// 	str := []byte(s)
// 	for i < j {
// 		for i < j {
// 			if isEnglishLetter(str[i]) {
// 				break
// 			} else {
// 				i++
// 			}
// 		}
// 		for i < j {
// 			if isEnglishLetter(str[j]) {
// 				break
// 			} else {
// 				j--
// 			}
// 		}
// 		str[i], str[j] = str[j], str[i]
// 		i++
// 		j--

// 	}
// 	return string(str)
// }

// func isEnglishLetter(c byte) bool {
// 	if (rune(c) >= 'a' && rune(c) <= 'z') || (rune(c) >= 'A' && rune(c) <= 'Z') {
// 		return true
// 	}
// 	return false
// }
func reverseOnlyLetters(s string) string {
    i, j := 0, len(s)-1
    str := []byte(s)
    
    for i < j {
        // Skip non-letters at the start
        if !isEnglishLetter(str[i]) {
            i++
            continue
        }
        // Skip non-letters at the end
        if !isEnglishLetter(str[j]) {
            j--
            continue
        }
        // Swap and move pointers
        str[i], str[j] = str[j], str[i]
        i++
        j--
    }
    return string(str)
}

func isEnglishLetter(c byte) bool {
    return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
}

/*
Output:
Test Case: [GivenString=ab-cd] PASSED: output: dc-ba, expected dc-ba
Test Case: [GivenString=a-bC-dEf-ghIj] PASSED: output: j-Ih-gfE-dCba, expected j-Ih-gfE-dCba
Test Case: [GivenString=Test1ng-Leet=code-Q!] PASSED: output: Qedo1ct-eeLg=ntse-T!, expected Qedo1ct-eeLg=ntse-T!
Test Case: [GivenString=7_28]] PASSED: output: 7_28], expected 7_28]
Test Case: [GivenString=-ab-] PASSED: output: -ba-, expected -ba-
*/

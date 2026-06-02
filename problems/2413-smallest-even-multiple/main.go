/*
Given a positive integer n, return the smallest positive integer that is a multiple of both 2 and n.
 

Example 1:

Input: n = 5
Output: 10
Explanation: The smallest multiple of both 5 and 2 is 10.
Example 2:

Input: n = 6
Output: 6
Explanation: The smallest multiple of both 6 and 2 is 6. Note that a number is a multiple of itself.
 

Constraints:

1 <= n <= 150
 */

package main

import "fmt"

type TestCase struct {
	GivenNumber int
	Expected    int
}

func main() {
	testCases := []TestCase{
		{100, 100},
		{17, 34},
	}

	for _, testCase := range testCases {
		result := smallestEvenMultiple(testCase.GivenNumber)
		if result == testCase.Expected {
			fmt.Printf("Test Case: [GivenNumber=%d] PASSED: output: %d, expected %d\n", testCase.GivenNumber, result, testCase.Expected)
		} else {
			fmt.Printf("Test Case: [GivenNumber=%d] FAILED: output: %d, expected %d\n", testCase.GivenNumber, result, testCase.Expected)
		}
	}

}

func smallestEvenMultiple(n int) int {
	if isEven(n) {
		return n
	} else {
		return 2 * n
	}
}

func isEven(n int) bool {
	if n%2 == 0 {
		return true
	}
	return false
}

/*
Output:
Test Case: [GivenNumber=100] PASSED: output: 100, expected 100
Test Case: [GivenNumber=17] PASSED: output: 34, expected 34
*/

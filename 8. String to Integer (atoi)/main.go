/*
Implement the myAtoi(string s) function, which converts a string to a 32-bit signed integer.

The algorithm for myAtoi(string s) is as follows:

Whitespace: Ignore any leading whitespace (" ").
Signedness: Determine the sign by checking if the next character is '-' or '+', assuming positivity if neither present.
Conversion: Read the integer by skipping leading zeros until a non-digit character is encountered or the end of the string is reached. If no digits were read, then the result is 0.
Rounding: If the integer is out of the 32-bit signed integer range [-231, 231 - 1], then round the integer to remain in the range. Specifically, integers less than -231 should be rounded to -231, and integers greater than 231 - 1 should be rounded to 231 - 1.
Return the integer as the final result.

 

Example 1:

Input: s = "42"

Output: 42

Explanation:

The underlined characters are what is read in and the caret is the current reader position.
Step 1: "42" (no characters read because there is no leading whitespace)
         ^
Step 2: "42" (no characters read because there is neither a '-' nor '+')
         ^
Step 3: "42" ("42" is read in)
           ^
Example 2:

Input: s = " -042"

Output: -42

Explanation:

Step 1: "   -042" (leading whitespace is read and ignored)
            ^
Step 2: "   -042" ('-' is read, so the result should be negative)
             ^
Step 3: "   -042" ("042" is read in, leading zeros ignored in the result)
               ^
Example 3:

Input: s = "1337c0d3"

Output: 1337

Explanation:

Step 1: "1337c0d3" (no characters read because there is no leading whitespace)
         ^
Step 2: "1337c0d3" (no characters read because there is neither a '-' nor '+')
         ^
Step 3: "1337c0d3" ("1337" is read in; reading stops because the next character is a non-digit)
             ^
Example 4:

Input: s = "0-1"

Output: 0

Explanation:

Step 1: "0-1" (no characters read because there is no leading whitespace)
         ^
Step 2: "0-1" (no characters read because there is neither a '-' nor '+')
         ^
Step 3: "0-1" ("0" is read in; reading stops because the next character is a non-digit)
          ^
Example 5:

Input: s = "words and 987"

Output: 0

Explanation:

Reading stops at the first non-digit character 'w'.

 

Constraints:

0 <= s.length <= 200
s consists of English letters (lower-case and upper-case), digits (0-9), ' ', '+', '-', and '.'.
*/
package main

import (
	"fmt"
	"math"
)

func main() {
	s := "9223372036854775808"
	fmt.Println(myAtoi(s))

}

func myAtoi(s string) int {
	var finalNumber int
	byteStr := []byte(s)
	byteStr = removeLeadingAndTrailingSpace(byteStr)
	if len(byteStr) == 0 {
		return 0
	}

	isNegative := false
	i := 0

	if byteStr[0] < '0' && byteStr[0] > '9' { // basic sanity checks
		if byteStr[0] != '+' && byteStr[0] != '-' {
			return 0
		}
	}

	if byteStr[i] == '-' { // check if negative or not.
		isNegative = true
		i++
	} else if byteStr[0] == '+' {
		i++
	}

	for i < len(byteStr) && byteStr[i] == '0' { // skiping the leading zero
		i++
	}

	for i < len(byteStr) {
		if byteStr[i] < '0' || byteStr[i] > '9' {
			break // stop the itreation at non interger char
		}
		// int32 overflow checkpoint before updating to the finalNumber
		if finalNumber > (math.MaxInt32-aciiToI(byteStr[i]))/10 {
			if isNegative {
				return math.MinInt32
			} else {
				return math.MaxInt32
			}
		}

		finalNumber = finalNumber*10 + aciiToI(byteStr[i])
		i++
	}
	if isNegative {
		finalNumber *= -1
	}
	return finalNumber
}

func aciiToI(char byte) int {
	return int(char - '0')
}

func removeLeadingAndTrailingSpace(str []byte) []byte {
	i, j := 0, len(str)-1
	for i < len(str) && str[i] == ' ' {
		i++
	}
	for j > i && str[j] == ' ' {
		j--
	}
	return str[i : j+1]
}

func removeLeadingAndTrailingSpaceV0(str []byte) []byte {
	if len(str) == 1 && str[0] == 32 {
		return []byte{}
	}
	k := 0
	i, j := 0, len(str)-1
	for k < len(str) && i < j && (str[i] == 32 || str[j] == 32) {
		if str[i] == 32 {
			i++
		}
		if str[j] == 32 {
			j--
		}
		k++
	}
	return str[i : j+1]
}

/*
Output:
2147483647
*/

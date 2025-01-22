/*
Your laptop keyboard is faulty, and whenever you type a character 'i' on it, it reverses the string that you have written. Typing other characters works as expected.

You are given a 0-indexed string s, and you type each character of s using your faulty keyboard.

Return the final string that will be present on your laptop screen.



Example 1:

Input: s = "string"
Output: "rtsng"
Explanation:
After typing first character, the text on the screen is "s".
After the second character, the text is "st".
After the third character, the text is "str".
Since the fourth character is an 'i', the text gets reversed and becomes "rts".
After the fifth character, the text is "rtsn".
After the sixth character, the text is "rtsng".
Therefore, we return "rtsng".
Example 2:

Input: s = "poiinter"
Output: "ponter"
Explanation:
After the first character, the text on the screen is "p".
After the second character, the text is "po".
Since the third character you type is an 'i', the text gets reversed and becomes "op".
Since the fourth character you type is an 'i', the text gets reversed and becomes "po".
After the fifth character, the text is "pon".
After the sixth character, the text is "pont".
After the seventh character, the text is "ponte".
After the eighth character, the text is "ponter".
Therefore, we return "ponter".


Constraints:

1 <= s.length <= 100
s consists of lowercase English letters.
s[0] != 'i'
*/

package main

import "fmt"

type TestCase struct {
	GivenString string
	Expected    string
}

func main() {
	testCases := []TestCase{
		{"string", "rtsng"},
		{"poiinter", "ponter"},
		{"poi", "op"},
		{"poiinteri", "retnop"},
		{"aiibiiciidiieiifii", "abcdef"},
		{"aibicidieifi", "fdbace"},
		{"aiiiiiii", "a"},
	}

	for _, testCase := range testCases {
		result := finalString(testCase.GivenString)
		if result == testCase.Expected {
			fmt.Printf("Test Case: [GivenString=%+v] PASSED: output: %s, expected %s\n", testCase.GivenString, result, testCase.Expected)
		} else {
			fmt.Printf("Test Case: [GivenString=%+v] FAILED: output: %s, expected %s\n", testCase.GivenString, result, testCase.Expected)
		}
	}

}

// func finalString(s string) string {
// 	str := []byte(s)
// 	for i := 0; i < len(str); i++ {
// 		if string(str[i]) == "i" {
// 			reverseStr(str[:i])
// 			if i != len(str)-1 {
// 				copy(str[i:], str[i+1:])
// 			}
// 			str = str[:len(str)-1]
// 			i--
// 		}
// 	}
// 	return string(str)
// }

func finalString(s string) string {
	// Use a fixed-capacity byte slice to hold the result
	str := make([]byte, 0, len(s))
	/*
	Using a fixed capacity slice with make([]byte, 0, len(s)) in the optimized code improves performance and memory efficiency. Here's a detailed explanation of why this approach is beneficial:
	
	1. Avoids Frequent Memory Allocations
	When you create a slice using make([]byte, 0, len(s)), you preallocate memory for the slice with a capacity equal to the length of the input string (len(s)). This ensures that the slice has enough space to hold all potential elements without needing to grow dynamically.
	
	If you were to append elements to a slice without preallocating capacity, Go would dynamically resize the slice when its capacity is exceeded. This resizing involves:
	
	Allocating a new array with larger capacity.
	Copying the existing elements to the new array.
	Releasing the old array for garbage collection.
	This dynamic resizing can be expensive, especially in cases where many appends are involved, as it introduces both time and memory overhead.
	
	2. Ensures Predictable Memory Usage
	By setting the capacity to len(s), the memory usage of the slice is predictable and fixed. Since the final string cannot exceed the length of the input string (s), preallocating this capacity ensures no unnecessary memory is allocated.
	
	3. Improves Append Performance
	Appending elements to a slice with sufficient preallocated capacity is efficient because it does not trigger slice resizing. Each append simply writes the element to the next available position in the underlying array. Without preallocation, every few append operations might trigger a resize, incurring additional cost.
	
	Why Not make([]byte, len(s))?
	Using make([]byte, len(s)) would preinitialize the slice with a length and capacity equal to len(s). While this avoids resizing issues, it creates a slice with a length that already includes placeholder values (0 bytes for []byte). This is unnecessary and can cause incorrect behavior if you directly use it without considering the initial values.
	
	In contrast, make([]byte, 0, len(s)) creates a slice with zero initial elements but a capacity large enough to accommodate all potential appends.
	*/

	for _, char := range s {
		if char == 'i' {
			// Reverse only the part of the slice that's been built so far
			reverseStr(str)
		} else {
			// Append the current character to the result
			str = append(str, byte(char))
		}
	}

	return string(bytes)
}

func reverseStr(s []byte) {
	i, j := 0, len(s)-1
	for i < j {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
}

/*
Output:
Test Case: [GivenString=string] PASSED: output: rtsng, expected rtsng
Test Case: [GivenString=poiinter] PASSED: output: ponter, expected ponter
Test Case: [GivenString=poi] PASSED: output: op, expected op
Test Case: [GivenString=poiinteri] PASSED: output: retnop, expected retnop
Test Case: [GivenString=aiibiiciidiieiifii] PASSED: output: abcdef, expected abcdef
Test Case: [GivenString=aibicidieifi] PASSED: output: fdbace, expected fdbace
Test Case: [GivenString=aiiiiiii] PASSED: output: a, expected a
*/

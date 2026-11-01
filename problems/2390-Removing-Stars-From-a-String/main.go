/*
2390. Removing Stars From a String

In one operation, you can:

Choose a star in s.
Remove the closest non-star character to its left, as well as remove the star itself.
Return the string after all stars have been removed.

Note:

The input will be generated such that the operation is always possible.
It can be shown that the resulting string will always be unique.
 

Example 1:

Input: s = "leet**cod*e"
Output: "lecoe"
Explanation: Performing the removals from left to right:
- The closest character to the 1st star is 't' in "leet**cod*e". s becomes "lee*cod*e".
- The closest character to the 2nd star is 'e' in "lee*cod*e". s becomes "lecod*e".
- The closest character to the 3rd star is 'd' in "lecod*e". s becomes "lecoe".
There are no more stars, so we return "lecoe".
Example 2:

Input: s = "erase*****"
Output: ""
Explanation: The entire string is removed, so we return an empty string.
 

Constraints:

1 <= s.length <= 105
s consists of lowercase English letters and stars *.
The operation above can be performed on s.
*/

package main

import (
  "fmt"
)

func main() {
	testcases := [][]string{
		{"leet**cod*e"}, {"lecoe"},
		{"erase*****"}, {""},
	}

	for i := 0; i < len(testcases); i += 2 {
	    fmt.Printf("testcase: input=%+v, expacted=%+v, result:%+v \n", testcases[i], testcases[i+1], removeStars(testcases[i][0]))
	}

}

func removeStars(s string) string {
    //return sol1(s)
    return sol2(s)
}

func sol2(s string) string { // inplace operation
    read, write := 0,0
    str := []byte(s)

    for ;read<len(str); read++ {
        if str[read] == '*' {
            write--
        } else {
            str[write]=str[read] // overwiting the nonstar letters.
            write++
        }
    }
    return string(str[:write])
}

func sol1(s string) string { // using extra space (O(n) array)
    var res []byte

    for i:=0; i<len(s); i++ {
        if s[i] == '*' && len(res) > 0 {
            res = res[:len(res)-1] // trimming the last element
        } else {
            res = append(res, s[i])
        }
    }
    return string(res)
}

/*
testcase: input=[leet**cod*e], expacted=[lecoe], result:lecoe 
testcase: input=[erase*****], expacted=[], result: 
*/

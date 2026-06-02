/*
You are given a string s, which contains stars *.
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

/*
1️⃣ Initial Stack-Based Approach (Very Slow)
func removeStars(s string) string {
    var stack []rune
    for _, c := range s {
        if c == '*' {
            if len(stack) != 0 {
                stack = stack[:len(stack)-1]
            } else {
                break
            }
        } else {
            stack = append(stack, c)
        }
    }
    result := ""
    for _, c := range stack {
        result += string(c)  // 🚨 Inefficient string concatenation
    }
    return result
}
❌ Problems:
result += string(c) creates a new string each iteration (O(n²) time).
Uses a slice of rune, which is unnecessary for ASCII.
Memory usage is very high (176MB) due to repeated allocations.
*/

/*
2️⃣ Improved strings.Builder Approach
func removeStars(s string) string {
    var stack []rune
    for _, c := range s {
        if c == '*' {
            if len(stack) != 0 {
                stack = stack[:len(stack)-1]
            } else {
                break
            }
        } else {
            stack = append(stack, c)
        }
    }
    var result strings.Builder
    for _, c := range stack {
        result.WriteRune(c)  // ✅ Avoids string concatenation overhead
    }
    return result.String()
}
✅ Improvements:

Uses strings.Builder, reducing memory overhead.
Execution time drops to 26ms (from 2186ms).
❌ Still Not the Best:

Still stores intermediate results in a slice.
Uses WriteRune(), which isn't necessary for single-byte ASCII.
*/


/*
3️⃣ In-Place Modification (Much Faster)
func removeStars(s string) string {
    str := []byte(s)
    write := 0
    for read := 0; read < len(s); read++ {
        if s[read] == '*' {
            write--  // Remove last character
        } else {
            str[write] = s[read]  // Overwrite
            write++
        }
    }
    return string(str[:write])
}
✅ Why is this faster?
Modifies the []byte array in-place (O(1) extra space).
Execution time improves to 2ms, 9.9MB.

❌ Still Has One Weakness:
string(str[:write]) creates a new string, causing a minor overhead
*/

/*
4️⃣ Best Approach: unsafe.String (0ms Execution Time)
*/

import "unsafe"

func removeStars(s string) string {
    stack := make([]byte, len(s))
    pos := 0
    for i := 0; i < len(s); i++ {
        if s[i] != '*' {
            stack[pos] = s[i]
            pos++
        } else {
            pos--  // Remove last added character
        }
    }
    return unsafe.String(&stack[0], pos)
}

/*
✅ Why is this the fastest?
1. Preallocates stack: Avoids dynamic memory allocation.
2. Only Uses Indexing (pos instead of append): No append() overhead.
3. unsafe.String(&stack[0], pos) Bypasses Copying:
   - Normally, string(stack[:pos]) creates a copy.
   - unsafe.String() directly constructs the string without extra allocation.

⛔ A Small Caveat with unsafe.String
The function unsafe.String(&stack[0], pos) creates a string from an existing memory address, bypassing Go’s internal safety checks. 
This is fine as long as stack remains in scope. However, if Go changes how memory is handled in future versions, this might break. 
For production code, sticking to string(stack[:pos]) might be safer.

That said, if you want the absolute fastest solution, unsafe.String is the best choice! 🚀
*/

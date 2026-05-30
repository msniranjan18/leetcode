package main

import (
	"fmt"
)

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
}

func lengthOfLongestSubstring(s string) int {
    
    set := make(map[byte]int) // cant use 26 length array as "s consists of English letters, digits, symbols and spaces."
    left := 0
    maxLen := 0

    for right:=0; right<len(s); right++ {
        ch :=s[right]
    
        if val, present := set[ch]; present && val >= left{ // val >= left we have to avoid the old value which is still there in the map
            left = val+1 // srinking the window to revome the duplicate 
        }
        set[ch]=right // update the new value to the map
        maxLen = max(maxLen, right-left+1)
    }
    return maxLen
}
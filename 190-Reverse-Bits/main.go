package main

import (
	"fmt"
	//"strconv"
)

func main() {
	fmt.Println(reverseBits(43261596))
}

func reverseBits(n int) int {
	var rev int
	
	for i:=0; i<32; i++ {
		rev <<= 1 // make room for comming bit by this way the bits will be reversed
		bit := n&1
		rev = rev | bit
		n >>= 1 // right sift by one so the n&1 oparate on n's next bit
	}

	return rev
}


/* 
The below approach introduces major bugs 
when dealing with the strict constraints of the 32-bit Reverse Bits problem.
There are three critical issues with this code that will cause it to fail:

1. Missing Leading Zeros (The Pad-Width Problem)
strconv.FormatInt converts a number to its shortest possible binary representation. 
It drops all leading zeros.
For example, if your input is n = 43261596:
    Its actual 32-bit binary representation is:
        00000010100101000001111010011100 (32 bits long)
strconv.FormatInt will output:10100101000001111010011100 (Only 26 bits long! The 6 leading zeros are gone).

When you reverse this shortened string, those 6 omitted zeros are completely missing. 
Instead of ending up as 6 trailing zeros in your final reversed answer, they vanish, corrupting the value.


2. Crash on Negative Numbers (-)
    If n is a negative number (or its 32-bit equivalent has a 1 as the sign bit),
    strconv.FormatInt(int64(n), 2) handles it by prepending a literal minus sign string ("-") to the front of the binary digits.
    If n = -4, it turns into the string: "-100"
    Your reverse function flips it to: "001-"
    strconv.ParseInt("001-", 2, 32) will throw an unhandled parsing error 
    because a minus sign at the end of a string is invalid formatting.
    
    
3. High Overhead / Performance Penalty
    In an interview setting, string allocations are costly.
    FormatInt allocates a string on the heap.
    Your manual reverse function creates another string allocation.
    ParseInt parses it out sequentially.
    This turns a simple hardware register level task ($O(1)$ constant bit shifts) 
    into an expensive $O(b)$ heap-allocation heavy routine that scales poorly.
    How to Fix It (If you must use strings)
    If you explicitly want to avoid bitwise operations and stick to your string approach, 
    you have to force Go to maintain a 32-character field width using
     fmt.Sprintf with a specific formatting verb, and parse it back as an unsigned integer 
     first to bypass sign bit parsing crashes:
     
    ```
    func reverseBits(n int) int {
	// 1. Force a 32-bit zero-padded binary string execution
	// %032b means: "format as binary, padded with leading zeros up to 32 characters"
	binaryStr := fmt.Sprintf("%032b", uint32(n))

	// 2. Reverse the string (using a simple rune/byte loop)
	runes := []rune(binaryStr)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	revBinaryStr := string(runes)

	// 3. Parse back as an UNSIGNED 64-bit value to avoid sign issues, then clamp
	revNum64, _ := strconv.ParseUint(revBinaryStr, 2, 32)
	
	return int(revNum64)
}
    ```


func reverseBits(n int) int {
	binaryStr:= strconv.FormatInt(int64(n), 2)
	revBinaryStr:=reverse(binaryStr)
	revNum64,_:=strconv.ParseInt(revBinaryStr, 2, 32)
	return int(revNum64)
}

func reverse(s string) string {
	n:=len(s)
	if n==0 {
		return ""
	}

	i:=0
	j:=n-1

	str:= []byte(s)

	for i<=j {
		str[i],str[j] = str[j], str[i]
		i++
		j--
	}
	return string(str)
}
*/
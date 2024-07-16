package main

import "fmt"

func main() {
	haystack := "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaao"
	needle := "aaao"
	fmt.Println(strStr(haystack, needle))
}

func strStr(haystack string, needle string) int {
	for i := 0; i <= len(haystack)-len(needle); i++ {
		if haystack[i:len(needle)+i] == needle {
			return i
		}
	}
	return -1
}

/*
Output:
32
*/

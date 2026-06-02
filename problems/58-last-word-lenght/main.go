package main

import (
	"fmt"
	"regexp"
)

func main() {
	fmt.Println("58. Easy")

	s := "      I     am   the                 greatfull                      "
	fmt.Println(lengthOfLastWord(s))
	fmt.Println(lengthOfLastWordv2(s))
	s = "Hello World "
	fmt.Println(lengthOfLastWord(s))
	fmt.Println(lengthOfLastWordv2(s))
	s = "      a      "
	fmt.Println(lengthOfLastWord(s))
	fmt.Println(lengthOfLastWordv2(s))
	s = "                   "
	fmt.Println(lengthOfLastWord(s))
	fmt.Println(lengthOfLastWordv2(s))
	s = ""
	fmt.Println(lengthOfLastWord(s))
	fmt.Println(lengthOfLastWordv2(s))
}

func lengthOfLastWord(s string) int {
	r := regexp.MustCompile(`(\s*)(?P<mid>[a-zA-Z]*)(\s*)$`)
	matches := r.FindStringSubmatch(s)
	lastIndex := r.SubexpIndex("mid")
	fmt.Println(matches[lastIndex])
	return len(matches[lastIndex])
}

func lengthOfLastWordv2(s string) int {
	l := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != 32 {
			l++
		} else if s[i] == 32 && l != 0 {
			return l
		}
	}
	return l
}

/*
Output:
58. Easy
greatfull
9
9
World
5
5
a
1
1

0
0

0
0
*/

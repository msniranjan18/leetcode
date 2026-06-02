package main

import (
	"fmt"
)

type stack []rune

func main() {
	fmt.Println("Valid-Parentheses")
	str := []string{
		"(",
		"()",
		"(){}[]",
		"({[",
		")(){}",
		"[])",
		"({[]})",
		"({[]}]",
		"",
	}
	for _, val := range str {
		fmt.Print(val, ": ")
		fmt.Println(isValid(val))
		fmt.Println()
	}
}

func isValid(s string) bool {
	stack := &stack{}
	var top rune
	var err error
	for _, char := range s {
		switch char {
		case '(', '{', '[':
			stack.push(char)
			continue
		case ')', '}', ']':
			top, err = stack.pop()
		}
		if err != nil || getClosingBracket(top) != char {
			return false
		}
	}
	if len(s) == 0 || len(*stack) != 0 {
		return false
	}
	return true
}

func (s *stack) push(data rune) {
	*s = append(*s, data)
}

func (s *stack) pop() (top rune, err error) {
	if len(*s) != 0 {
		top = (*s)[len(*s)-1]
		*s = (*s)[:len(*s)-1]
	} else {
		return top, fmt.Errorf("Empty")
	}
	return top, nil
}

func getClosingBracket(bracket rune) (closedBracket rune) {
	switch bracket {
	case '(':
		closedBracket = ')'
	case '{':
		closedBracket = '}'
	case '[':
		closedBracket = ']'
	}
	return closedBracket
}

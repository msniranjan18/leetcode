package main

import (
	"fmt"
)

func main(
	fmt.Println(binary2Int("101"))
)

func binary2Int(b string) int {
	n:=len(b)
	if n == 0 {
		return 0
	}

	i:=n-1
	num:=0
	for i>=0 {
		digit:=int(b[i]-'0')
		num = num*2 + digit
		i--
	}
	return num
}
package main

import (
	"fmt"
)

func main() {
	fmt.Println(int2Binary(5))
}



func int2Binary(n int) string {
	if n==0 {
		return "0"
	}
	b := make([]byte, 64)
	i:=63
	for i>=0 && n!=0 {
		r:=n%2
		if r == 1 {
			b[i]= '1'
		} else {
			b[i] = '0'
		}
		n=n/2
		i--
	}
	return string(b[i+1:])
}
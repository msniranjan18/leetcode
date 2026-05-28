package main

import (
	"fmt"
)

func main(){
	fmt.Println(countBits(16))
}

func countBits(n int) []int {
	/*
	Logic:
			3 => 3%2 = 1
	3/2 = 1 => 1%2 = 1
	1/2 = 0 

			5 => 5%2 = 1
	5/2 = 2 => 2%2 = 0
	2/2 = 1 => 1%2 = 1
	1/2 = 0 

			15 => 15%2 = 1
	15/2 = 7 => 7%2 = 1
	7/2 = 3 => 3%2 = 1
	3/2 = 1 => 1%2 = 1
	1/2 = 0 
	*/

   	res:=make([]int, n+1)
	res[0] = 0
	
	for i:=1; i<=n; i++ {
		num := i
		for num!=0 {
			res[i]+=num%2
			num/=2
		}
	}
	return res
}
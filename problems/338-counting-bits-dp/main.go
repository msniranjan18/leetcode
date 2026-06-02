package main

import (
	"fmt"
)

func main(){
	fmt.Println(countBits(16))
}

func countBits(n int) []int {
	// below code has O(nLogn) complexcity
   	// res:=make([]int, n+1)
	// res[0] = 0
	
	// for i:=1; i<=n; i++ {
	// 	num := i
	// 	for num!=0 {
	// 		res[i]+=num%2
	// 		num/=2
	// 	}
	// }
	// return res

	return countBitsUsingDP(n)
}

func countBitsUsingDP(n int) []int {
	/*
	Logic:
	n => Bits => 1bits 	=> Formula
	0 => 0000 => 0     	=> dp[0] = 0
	1 => 0001 => 1		=> dp[1] = 1
	2 => 0010 => 1		=> dp[2] = 1
	3 => 0011 => 2		=> dp[2]+dp[n-2] => 1 + 1 = 2
	4 => 0100 => 1     	=> dp[4] = 1
	5 => 0101 => 2		=> dp[4]+dp[n-4] => 1 + 1 => 2
	6 => 0110 => 2      => dp[4]+dp[n-4] => 1 + 1 => 2
	7 => 0111 => 3		=> dp[4]+dp[n-4] => 1 + 2 => 2
	8 => 1000 => 1      => dp[8] = 1

	formula is:
	dp[offset] + dp[n-offset]

	Offset will be: [0,1,2,4,8,16,32,...]
	
	offset formula:
	currentOffSet = previousOffset * 2 
	or
	if offset*2 == n {
		offset = n
		dp[n] = 1
	}
	*/

	if n==0{
		return []int{0}
	} else if n==1{
		return []int{0,1}
	} else if n==2{
		return []int{0,1,1}
	}

	dp := make([]int, n+1)

	dp[0] = 0
	dp[1] = 1
	dp[2] = 1

	offset := 2

	for i:=3; i<=n; i++ {
		if offset*2 == i {
			dp[i] = 1
			offset *= 2
		} else {
			dp[i] = dp[offset] + dp[i-offset]
		}
	}
	return dp
}




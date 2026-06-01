package main

import (
	"fmt"
)

func main() {
	fmt.Println(generateParenthesis(3))
	fmt.Println(generateParenthesis(4))
	fmt.Println("8",len(generateParenthesis(8)))
	fmt.Println("10",len(generateParenthesis(10)))
	fmt.Println("12",len(generateParenthesis(12)))
	fmt.Println("14",len(generateParenthesis(14)))
	fmt.Println("16",len(generateParenthesis(16)))
	fmt.Println("18",len(generateParenthesis(18)))
	fmt.Println("20",len(generateParenthesis(20)))
	fmt.Println("22",len(generateParenthesis(22)))

}

func generateParenthesis(n int) []string {
    open, close := 0,0
    res := make([]string,0) // later we'll optimize the if preallocation is posible basedd on given n.
    //ans:="" //later we'll optimize to make([]byte, 0) as string concatination are costly
	ans:=make([]byte,0)
    res= solve(n, open, close, ans, res)
    return res
}

func solve(n, open, close int, ans []byte, res []string ) []string {
    //base case
    if open == n && close == n {
        res = append(res, string(ans))
        return res
    }
    if open > n || close > n {
        return res
    }

    if open == close {
		ans = append(ans, '(')
        res = solve(n,open+1, close, ans, res)
		ans = ans[:len(ans)-1]

    } else {
		ans = append(ans, '(')
        res = solve(n,open+1, close, ans, res)
		ans = ans[:len(ans)-1]

		ans = append(ans, ')')
        res = solve(n,open, close+1, ans, res)
		ans = ans[:len(ans)-1]
    }
    return res
}
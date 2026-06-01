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
func generateParenthesisTest(n int) int {
	return len(generateParenthesis(n))
}

func generateParenthesis(n int) []string {
    /*
	open, close := 0,0
    res := make([]string,0) // later we'll optimize the if preallocation is posible basedd on given n.
    //ans:="" //later we'll optimize to make([]byte, 0) as string concatination are costly
	ans:=make([]byte,0)
    return solve(n, open, close, ans, res)
	*/
    //openCount, closeCount := 0,0
    //res:=make([]string,0)
    /*
    str:="" // -------------------------------------> string concation is very expensive and it not recconmonded as it creates new string as strings as immutable which puts the presure on GC so avoid it and use []byte or strings.Builder instead.
    return sol1(n, openCount, closeCount, str, res)
    */
    //str := []byte{}
    //return sol2(n, openCount, closeCount, str, res)
    
    return sol3(n)
}

func sol3(n int) []string {
    openCount, closeCount := 0,0
    res:=make([]string,0)
    str := []byte{}
    
    // Use a local function (closure). 
    // This keeps the main function signature clean and avoids passing 'res' and 'str' everywhere.
    var bracktrack func(openCount, closeCount int) []string
    bracktrack = func(openCount, closeCount int) []string {

        if openCount == n && closeCount == n {
            res = append(res, string(str))
            return res
        }
        
        if openCount < n {
            str = append(str, '(')
            res = bracktrack(openCount+1, closeCount)
            // Backtrack: Remove the character before trying the next path
            str = str[:len(str)-1]
        }
        if closeCount < openCount {
            str = append(str, ')')
            res = bracktrack(openCount, closeCount+1)
            // Backtrack: Remove the character before trying the next path
            str = str[:len(str)-1]
        }
        return res
    }

    bracktrack(openCount, closeCount) // string 0,0
    return res
}

func sol2(n, openCount, closeCount int, str []byte, res []string) []string {
    if openCount == n && closeCount == n {

        res = append(res, string(str))
        return res
    }
    
    if openCount < n {
        str = append(str, '(')
        res = sol2(n, openCount+1, closeCount, str, res)
        // Backtrack: Remove the character before trying the next path
        str = str[:len(str)-1]
    }
    if closeCount < openCount {
        str = append(str, ')')
        res = sol2(n, openCount, closeCount+1, str, res)
        // Backtrack: Remove the character before trying the next path
        str = str[:len(str)-1]
    }
    return res
}

func sol1 (n, openCount, closeCount int, str string, res []string) []string {
    if openCount == n && closeCount == n {
        res = append(res, str)
        return res
    }
    
    if openCount < n {
        res = sol1(n, openCount+1, closeCount, str+"(", res) 
        // Backtrack: we dont need to remove '(' as strings are immutable and this concatination is creating new string and passing to the next function so original str is unchanged
    }
    if closeCount < openCount {
        res = sol1(n, openCount, closeCount+1, str+")", res) 
        // Backtrack: we dont need to remove '(' as strings are immutable and this concatination is creating new string and passing to the next function so original str is unchanged
    }
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
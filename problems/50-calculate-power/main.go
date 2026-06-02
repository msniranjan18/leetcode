package main

import (
	"fmt"
)

func main() {
	fmt.Println(myPow(1.0, 2147483647))
	fmt.Println(myPow(1.0, -2147483647))
	fmt.Println(myPow(2.0, 10))
	fmt.Println(myPow(2.0, -2))
}

func myPow(x float64, n int) float64 {
    //return solution1(x,n)
    return solution2(x,n)
}

func solution2(x float64, n int) float64 {
    //res:= pow(x,n)
    res:=powv2(x,n)

    if n<0{
        return 1/res
    }
    return res
}

func powv2(x float64, n int) float64 {
    var res float64
    res = 1.0
    if x == 0 {
        return 0
    }
    if n == 1 {
        return x
    }
    if n == 0 {
        return 1
    }
    
    res = powv2(x,n/2) // caculate only once and use the result to calculate further.s
    res = res*res
    if n % 2 != 0 {
        res *= x
    }
    return res
}

/*
The reason your code is still hitting a TLE (Time Limit Exceeded) is because of 
how you structured the recursive calls.
You correctly identified that x^n can be broken down using n/2.
However, look closely at this line:Gores = pow(x, n/2) * pow(x, n/2)
By calling pow(x, n/2) twice separately, 
you accidentally created a massive tree of redundant calculations. 
Instead of halving the work, your code branches out exponentially,
executing O(n) total operations. 
For n = 2,147,483,647, it is still performing billions of steps!
*/
func pow(x float64, n int) float64 {
    var res float64
    res = 1.0
    if x == 0 {
        return 0
    }
    if n == 1 {
        return x
    }
    if n == 0 {
        return 1
    }

    if n % 2 == 0 {
        res = pow(x,n/2) * pow(x,n/2)
    } else {
        res = pow(x,n/2) * pow(x,n/2) * pow(x,1)
    }
    return res
}


/*
below code is hitting two major roadblocks: 
Time Limit Exceeded (TLE) and an Integer Overflow bug.
Here is exactly why your current logic breaks down, followed by the correct strategic approach to fix it.
Why Your Current Solution Fails

1. The Bottleneck: Linear Time Complexity O(n)
Your loop runs exactly pow times. According to the constraints, n can be as large as 2,147,483,647.
Running a loop over 2 billion times to multiply floats one by one takes way too long, forcing LeetCode to cut your program off with a "Time Limit Exceeded" error.

2. The Edge Case: Overflowing abs(n)
Look closely at the constraints: -2pow31 <= n <= 2pow31-1.
The minimum possible value for a 32-bit signed integer is -2147483648.
The maximum possible value is 2147483647.When n = -2147483648, 
your abs function tries to multiply it by -1. 
However, positive 2147483648 cannot fit into a signed 32-bit integer! 
It overflows right back into a negative number, creating an infinite loop or broken logic.
*/
func solution1(x float64, n int) float64 {
    var res float64
    res=1.0
    pow:= abs(n)
    for i:=0; i<pow; i++ {
        res *= x
    }
    
    if n < 0 {
        res = 1.0/res 
    }
    return res
}

func abs(n int) int {
    if n < 0 {
        return n*-1
    }
    return n
}
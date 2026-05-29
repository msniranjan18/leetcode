package  main

import (
	"fmt"
)

func main() {
	fmt.Println(isHappy(2))
	fmt.Println(isHappy(19))
}

func isHappy(n int) bool {
    fast:=nextNumber(n)
    slow:=n
    for {
        if fast == 1 {
            return true
        }
        if fast == slow {
            break
        }
        fast = nextNumber(nextNumber(fast))
        slow = nextNumber(slow)
    }
    return false
}

func nextNumber(n int) int {
    num:=0
    for n!=0 {
        digit := n%10
        num+= digit*digit
        n/=10
    }
    return num
}
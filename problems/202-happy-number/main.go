package  main

import (
	"fmt"
)

func main() {
	fmt.Println(isHappy(2))
	fmt.Println(isHappy(19))
}

func isHappy(n int) bool {
    //return bruteforceApproach(n)
    return floydCycleDetectionApproach(n)
}

func floydCycleDetectionApproach(n int) bool {
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

func bruteforceApproach(n int) bool {
    m := make(map[int]struct{}, 0)
    for {
        n = nextNumber(n)
        if n == 1 {
            return true
        } else if _, ok:= m[n]; !ok {
            m[n] = struct{}{}
        } else {
            return false
        } 
    }
    return true
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
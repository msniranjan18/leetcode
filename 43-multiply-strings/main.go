package main

import (
	"fmt"
	"strings"
	"strconv"
)

func main() {
	fmt.Println(multiply("123","456"))
}


func multiply(num1 string, num2 string) string {
    //return solution2(nums1, nums2)
    return solution2(num1, num2)
}

/*
The Strategy: Positional Grid Multiplication
Instead of maintaining intermediate numbers like sum and res as integers, we can simulate multiplication directly inside a fixed-size digit slice ([]int).

When we multiply a digit from num1 at index i by a digit from num2 at index j:

The maximum length of the resulting product string is at most len(num1) + len(num2).

The intermediate result of multiplying num1[i] by num2[j] will naturally settle into two specific indices in our results array: i + j and i + j + 1.

By using a pre-allocated slice of size len(num1) + len(num2), we can add products to our array columns dynamically and handle any lingering carries on the fly.
*/
func solution2(num1 string, num2 string) string {
    if num1 == "0" ||  num2 == "0" {
        return "0"
    }
    
    n,m := len(num1), len(num2)
    res := make([]int, n+m)

    // 123 x 456
    for i:=n-1; i>=0; i-- {
        d1:=int(num1[i]-'0')
        for j:=m-1; j>=0; j-- {
            d2:=int(num2[j]-'0')
            mul := d1*d2
            // writing right to left to res 
            pos1:=i+j
            pos2:=i+j+1 // if i=n-1 and j=m-1 

            res[pos2] += mul
            res[pos1] += res[pos2]/10 // add carry
            res[pos2] = res[pos2]%10
        }
    }

    // Removing leading zeros
    i:=0
    for ; i<len(res); i++ {
        if res[i] != 0 {
            break
        }
    }

    //fmt.Println(res, i)
    res = res[i:]
    //fmt.Println(res, i)
    var sb strings.Builder

    for i:=0; i<len(res); i++ {
        sb.WriteString(strconv.Itoa(res[i]))
    }
    return sb.String()
}

/*
not working for bigger values like:
num1="498828660196"
num2="840477629533"
code breaks due to a critical constraint violation hidden in the problem details:

Constraints: 1 <= num1.length, num2.length <= 200

Because the input strings can be up to 200 digits long, the final product could have up to 400 digits. Your variable res := 0 (and sum) utilizes standard 64-bit integers (int), which will completely overflow and crash into garbage values once a number exceeds roughly 19 digits. Since we aren't allowed to use BigInteger, we have to handle the arithmetic directly inside an array or slice.
*/

func solution1(num1 string, num2 string) string {
    n:=len(num1)
    m:=len(num2)
    small:=num1
    big:=num2
    if n > m {
        small = num2
        big = num1
    }

    res:=0
    shift:=1
    // loop on smaller str
    for i:=len(small)-1; i>=0; i-- {
        d1:=int(small[i]-'0')
        
        carry:=0
        curDigit:=0
        power := 1
        sum := 0

        for j:=len(big)-1; j>=0; j-- {
            d2:=int(big[j]-'0')
            total := d1 * d2 + carry
            curDigit = total % 10
            carry = total / 10

            sum += curDigit * power
            power *= 10
        }
        if carry > 0 {
            sum += carry * power
        }

        sum = sum * shift
        shift *= 10
        res += sum
        //fmt.Println(sum)
    }
    return strconv.Itoa(res)
}
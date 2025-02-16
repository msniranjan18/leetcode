/*
Given a positive integer n, return the punishment number of n.

The punishment number of n is defined as the sum of the squares of all integers i such that:

1 <= i <= n
The decimal representation of i * i can be partitioned into contiguous substrings such that the sum of the integer values of these substrings equals i.
 

Example 1:

Input: n = 10
Output: 182
Explanation: There are exactly 3 integers i in the range [1, 10] that satisfy the conditions in the statement:
- 1 since 1 * 1 = 1
- 9 since 9 * 9 = 81 and 81 can be partitioned into 8 and 1 with a sum equal to 8 + 1 == 9.
- 10 since 10 * 10 = 100 and 100 can be partitioned into 10 and 0 with a sum equal to 10 + 0 == 10.
Hence, the punishment number of 10 is 1 + 81 + 100 = 182
Example 2:

Input: n = 37
Output: 1478
Explanation: There are exactly 4 integers i in the range [1, 37] that satisfy the conditions in the statement:
- 1 since 1 * 1 = 1. 
- 9 since 9 * 9 = 81 and 81 can be partitioned into 8 + 1. 
- 10 since 10 * 10 = 100 and 100 can be partitioned into 10 + 0. 
- 36 since 36 * 36 = 1296 and 1296 can be partitioned into 1 + 29 + 6.
Hence, the punishment number of 37 is 1 + 81 + 100 + 1296 = 1478
 

Constraints:

1 <= n <= 1000
*/


/*
                  (1296, sum=0)
                /       |       \
          "1" (296, 1) "12"(96,12) "129"(6,129)
         /      \  
 "2"(96,3)   "29"(6,30)   
     /           \
 "9"(6,12)   "6"(,36 ✅)
   /  
"6"(,18 ❌)  

*/

/*
Explanation of Paths
1 → 2 → 9 → 6 → Sum = 18 ❌ (Backtrack)
1 → 2 → 96 → Sum = 99 ❌ (Backtrack)
1 → 29 → 6 → Sum = 36 ✅ (Valid partition!)
✅ Valid partition found: "1" + "29" + "6" = 36
*/

/*

*/

func punishmentNumber(n int) int {
    sum := 0
    for i := 1; i <= n; i++ {
        square := i * i
        if canPartition(square, i, 0) {  // ✅ Fix: Pass `square` as integer
            sum += square
        }
    }
    return sum
}

func canPartition(num, target, currSum int) bool {
    fmt.Println("canPartition func", num)
    if num == 0 {
        return currSum == target
    }

    temp := num
    part := 0
    factor := 1

    for temp > 0 {
        digit := temp % 10
        part = digit * factor + part
        factor *= 10
        temp /= 10
       fmt.Printf("num = %d, currSum = %d, temp = %d, part = %d, digit = %d, factor = %d\n",num, currSum, temp, part, digit,factor)
        if currSum+part <= target {
            if canPartition(temp, target, currSum+part) {
                return true
            }
        }
    }
    return false
}



// func punishmentNumber(n int) int {
//     sum:=0
//     square := 0
//     for i:=1;i<=n;i++{
//         square = i*i
//         if canPartition(strconv.Itoa(square), i, 0, 0) {
//             sum += square
//         }
//     }
//     return sum
// }

// func canPartition(s string, target, currSum, start int) bool {
//     if start == len(s) {
//         return currSum == target // return true for valid path
//     }

//     num:=0
//     for end:=start; end < len(s); end++ { 
//         num = num*10+int(s[end]-'0') // take the num
//         if currSum + num > target { // If sum exceeds target, stop exploring this path
//             break
//         }

//         // Recurse with the new sum but do NOT modify currSum directly
//         if canPartition(s, target, currSum + num, end+1) {
//             return true
//         }

//     }
//     return false
// }

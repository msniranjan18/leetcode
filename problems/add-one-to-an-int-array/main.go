package main

import (
	"fmt"
	"strconv"
)

/*
Description:
66. Plus One
You are given a large integer represented as an integer array digits, where each digits[i] is the ith digit of the integer. The digits are ordered from most significant to least significant in left-to-right order. The large integer does not contain any leading 0's.

Increment the large integer by one and return the resulting array of digits.



Example 1:

Input: digits = [1,2,3]
Output: [1,2,4]
Explanation: The array represents the integer 123.
Incrementing by one gives 123 + 1 = 124.
Thus, the result should be [1,2,4].
Example 2:

Input: digits = [4,3,2,1]
Output: [4,3,2,2]
Explanation: The array represents the integer 4321.
Incrementing by one gives 4321 + 1 = 4322.
Thus, the result should be [4,3,2,2].
Example 3:

Input: digits = [9]
Output: [1,0]
Explanation: The array represents the integer 9.
Incrementing by one gives 9 + 1 = 10.
Thus, the result should be [1,0].


Constraints:

1 <= digits.length <= 100
0 <= digits[i] <= 9
digits does not contain any leading 0's.
*/

func main() {
	digits := []int{9, 9, 9}
	fmt.Println("[9,9,9]", "->", plusOneV2(digits))

	digits = []int{8, 9}
	fmt.Println("[8,9]", "->", plusOneV2(digits))

	digits = []int{9}
	fmt.Println("[9]", "->", plusOneV2(digits))

	digits = []int{7, 2, 8, 5, 0, 9, 1, 2, 9, 5, 3, 6, 6, 7, 3, 2, 8, 4, 3, 7, 9, 5, 7, 7, 4, 7, 4, 9, 4, 7, 0, 1, 1, 1, 7, 4, 0, 0, 6}
	fmt.Println(digits, "->", plusOneV2(digits))

}

func plusOneV1(digits []int) []int {
	res := []int{}
	temp := ""
	for _, digit := range digits {
		temp += strconv.Itoa(digit)
	}
	d, _ := strconv.Atoi(temp)
	d += 1
	temp = strconv.Itoa(d)
	for _, digit := range temp {
		d, _ = strconv.Atoi(string(digit))
		res = append(res, d)
	}
	return res
}

func plusOneV2(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i]+1 < 10 {
			digits[i] = digits[i] + 1
			return digits
		} else {
			digits[i] = 0
		}
	}
	digits = append([]int{1}, digits...)
	return digits
}

/*
Output:
[9,9,9] -> [1 0 0 0]
[8,9] -> [9 0]
[9] -> [1 0]
[7 2 8 5 0 9 1 2 9 5 3 6 6 7 3 2 8 4 3 7 9 5 7 7 4 7 4 9 4 7 0 1 1 1 7 4 0 0 7] -> [7 2 8 5 0 9 1 2 9 5 3 6 6 7 3 2 8 4 3 7 9 5 7 7 4 7 4 9 4 7 0 1 1 1 7 4 0 0 7]
*/

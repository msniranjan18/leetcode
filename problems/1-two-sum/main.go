package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func numbers(input string) []int {
	var nums []int
	strNums := strings.Fields(input)
	for _, v := range strNums {
		n, _ := strconv.Atoi(v)
		nums = append(nums, n)
	}
	return nums
}

func readCommaSeperatedInputNumbers() string {
	var input string
	fmt.Println("Enter comma-separated integers: like 1,2,3")
	fmt.Scanln(&input)
	return strings.Join(strings.Split(input, ","), " ")
}

func readSpaceSeperatedInputNumbers() string {
	fmt.Println("Enter space-separated integers: Like 1 2 3 4")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()

}
func main() {
	var target int
	//input := readCommaSeperatedInputNumbers()
	input := readSpaceSeperatedInputNumbers()
	fmt.Println("Enter target for finding sum: ")
	fmt.Scanln(&target)
	numsSlice := numbers(input)
	//fmt.Println(twoSum(numsSlice, target))
	fmt.Println(twoSumv1(numsSlice, target))
	//fmt.Println(twoSumOptimal(numsSlice, target))
}

package main

import "fmt"

func twoSum(nums []int, target int) []int {
	var result []int
	i := 0
	j := i + 1
	itr := 1
	for {

		if i == j {
			j++
		}
		// loop break condition
		if nums[i]+nums[j] == target {
			result = append(result, i)
			result = append(result, j)
			fmt.Println("Itration:", itr)
			return result
		}
		if i >= len(nums) {
			return result
		}
		if j == len(nums)-1 {
			i++
			j = 0
		}
		j++
		itr++
	}
}

func twoSumv1(nums []int, target int) []int {
	arrLen := len(nums)

	itr := 0
	for i := 0; i < arrLen; i++ {
		for j := i + 1; j < arrLen; j++ {
			if nums[i]+nums[j] == target {
				fmt.Println("No Of itration: ", itr)
				return []int{i, j}
			}
			itr++
		}
	}

	return nil
}

func twoSumOptimal(nums []int, target int) []int {
	tempMap := make(map[int]int)
	itr := 1
	for i, v := range nums {
		if value, ok := tempMap[v]; ok {
			fmt.Println("NO of itrations: ", itr)
			return []int{value, i}
		}
		tempMap[target-v] = i
		itr++
	}
	return []int{}
}

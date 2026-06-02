package main

import "fmt"

func main() {
	nums := []int{3, 2, 2, 3}
	//l := removeElement(nums, 2)
	l := removeElementOptimize(nums, 3)

	fmt.Println(nums[:l])

}

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	i := 0
	j := len(nums) - 1
	for i < j {
		if nums[i] == val {
			if nums[j] == val {
				j--
				continue
			}
			//swap
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j--
		} else {
			i++
		}
	}
	if nums[i] != val {
		i++
	}
	return i
}

func removeElementOptimize(nums []int, val int) int {
	k := 0
	for i, _ := range nums {
		if nums[i] != val {
			nums[k] = nums[i]
			k++
		}
	}
	return k
}

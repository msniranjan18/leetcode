package main

import "fmt"

func main() {

	nums := []int{-100,-100, 1, 2, 2, 3,3,3,4,4}

	i := 0 // points to 1st index
	k := 0 // points to zeroth index
	for range nums {
		if nums[k] != nums[i] {
			k++
			nums[k] = nums[i]
		}
		i++
	}
	fmt.Println(nums)
	fmt.Println(k, nums[:k+1])
}

/*
Output:
go run main.go
[-100 1 2 3 4 3 3 3 4 4]
4 [-100 1 2 3 4]
*/

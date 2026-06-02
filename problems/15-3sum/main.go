package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println(threeSum([]int{0,0,0,0}))
	fmt.Println(threeSum([]int{-1,0,1,2,-1,-4}))
	
}

func threeSum(nums []int) [][]int {
    n := len(nums)
    slices.Sort(nums)
    // res:=make([][]int,0)
	/*
	Without preallocation:
	manav@MSNiranjan:~/workdir/leetcode/15-3sum$ go test -bench=. -benchmem
	goos: linux
	goarch: amd64
	pkg: github.com/msniranjan18/leetcode/15-3sum
	cpu: 13th Gen Intel(R) Core(TM) i5-13500H
	BenchmarkThreeSum-16                  68          16594867 ns/op         2128922 B/op      16278 allocs/op
	PASS
	ok      github.com/msniranjan18/leetcode/15-3sum        2.017s
	*/

	res:=make([][]int,0, len(nums)) // Dynamic guessing: Allocate space proportional to the input size
	/*
	With preallocation: 
	manav@MSNiranjan:~/workdir/leetcode/15-3sum$ go test -bench=. -benchmem
	goos: linux
	goarch: amd64
	pkg: github.com/msniranjan18/leetcode/15-3sum
	cpu: 13th Gen Intel(R) Core(TM) i5-13500H
	BenchmarkThreeSum-16                  61          16614521 ns/op         2208818 B/op      16266 allocs/op
	PASS
	ok      github.com/msniranjan18/leetcode/15-3sum        1.039s ----> 
	manav@MSNiranjan:~/workdir/leetcode/15-3sum$
	*/

    for i:=0; i<n; i++ {
        if i>0 && nums[i] == nums[i-1] { // i shall not be a duplicate
            //fmt.Println("Avoiding I",nums[i], nums[i-1])
            continue
        }  
        j:=i+1
        k:=n-1
        for j < k {
            total := nums[i]+nums[j]+nums[k]
            //fmt.Println("total", nums[i]+nums[j]+nums[k])
            if total == 0 {
                res = append(res, []int{nums[i], nums[j],nums[k]})
                //fmt.Println("Equal", res)
                j++
                k--
				// skip the same j and k to avoid record the same elements.
                // for j>i+1 && j<k && nums[j] == nums[j-1] { j++ }
                // for k<n-1 && j<k && nums[k] == nums[k+1] { k-- }
				/*
				For the j Loop (nums[j] == nums[j+1]):
					Why j > i+1 isn't needed: This loop only runs inside the total == 0 block. 
					Because j starts at i+1 and we only enter this block when j < k, j is already safely past i.

					Why it's safe: It looks forward (j+1). The safety condition j < k guarantees that j+1 can at most equal k. 
					Since k is bounded by n-1, j+1 can never overshoot the end of the array. It is physically impossible to index out of bounds.

				For the k Loop (nums[k] == nums[k-1]):
					Why k < n-1 isn't needed: This loop looks backward (k-1). 
					The safety condition j < k guarantees that k-1 can at lowest equal j. 
					Since j is at least i+1 (which is at least 1), k-1 can never drop below 0. 
					It is completely safe from crashing on a negative index.
				*/
				for j<k && nums[j] == nums[j-1] { j++ }
                for j<k && nums[k] == nums[k+1] { k-- }
            } else if total < 0 {
                j++
            } else {
                k--
            }           
        } 
    }
    return res
}
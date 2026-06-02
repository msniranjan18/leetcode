package main

import (
	"fmt"
	"slices"
	"math"
)

func main() {
	fmt.Println(threeSumClosest([]int{-1,2,1,-4}, 1))
}

func threeSumClosest(nums []int, target int) int {
	//return solution1(nums, target)
	//return solution2(nums, target)
    return optimalSolution(nums, target)
}


// solution1 + solution2
func optimalSolution(nums []int, target int) int {
	n := len(nums)
	slices.Sort(nums)

	// Initialize using actual data values to prevent default zero boundary bugs
	closestSum := nums[0] + nums[1] + nums[2]
	minDiff := abs(target - closestSum)

	for i := 0; i < n-2; i++ { // n-2 as we need three elements
		// Skip duplicate values for 'i' to eliminate redundant loops
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		j := i + 1
		k := n - 1

		for j < k {
			currentSum := nums[i] + nums[j] + nums[k]
			
			// Short-circuit: if we hit the exact target, no other combo can beat it
			if currentSum == target {
				return currentSum
			}

			diff := abs(target - currentSum)
			if diff < minDiff {
				minDiff = diff
				closestSum = currentSum
			}

			// Move pointers and aggressively fast-forward past duplicate elements
			if currentSum < target {
				j++
				for j < k && nums[j] == nums[j-1] { j++ }
			} else {
				k--
				for j < k && nums[k] == nums[k+1] { k-- }
			}
		}
	}

	return closestSum
}

/*
Solution 1 initializes closestSum = 0. What happens if your input array is [100, 100, 100] and your target is 100? 
The closest possible sum is 300. However, because your code initializes closestSum = 0, 
it returns a completely invalid placeholder sum if it never updates correctly, 
or relies on an uninitialized state. 
Solution 2 handles this perfectly by initializing closestSum to the sum of the first three elements: nums[0] + nums[1] + nums[2].
*/
func solution1(nums []int, target int) int {
    n := len(nums)
    slices.Sort(nums)
    closestSum := 0
    minDiff := math.MaxInt

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
            diff:=abs(target - total) // closer to target so -ve < targer < +ve
            if diff < minDiff {
                closestSum = total
                minDiff = diff
                //fmt.Println(closestSum, minDiff, total, i, j, k)
            }
            if total <= target {
                j++
				for j<k && nums[j] == nums[j-1] { j++ }
            } else {
                k--
                for j<k && nums[k] == nums[k+1] { k-- }
            }           
        } 
    }
    return closestSum 
}

/*
Loop Invariance Bound Check
Code 2 limits the outer loop to i < n-2. 
This is structurally correct since you need at least 3 elements. 
However, Code 2 completely misses duplicate skipping on j and k. 
If an input array contains nothing but duplicates, 
Code 2 will painfully recalculate the exact same window thresholds over and over again, wasting valuable CPU cycles.
*/
func solution2(nums []int, target int) int {
    n := len(nums)
    slices.Sort(nums)

    if n < 3 {
        return -1
    }

    closestSum := nums[0] + nums[1] + nums[2]
    
    for i := 0; i<n-2; i++ {
        if i>0 && nums[i] == nums[i-1] {
            continue
        }

        j := i+1
        k := n-1

        for j<k {
            currentSum := nums[i] + nums[j] + nums[k]
            if abs(target - currentSum) < abs(target - closestSum) {
                closestSum = currentSum
            }

            if currentSum == target {
                return currentSum
            } else if currentSum > target {
                k-=1
            } else {
                j+=1
            }
        }
    }
    return closestSum
}

func abs(n int) int {
    if n<0 {
        return n*-1
    }
    return n
}
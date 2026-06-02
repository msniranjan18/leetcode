/*
Given an array nums of n integers where nums[i] is in the range [1, n], return an array of all the integers in the range [1, n] that do not appear in nums.

 

Example 1:

Input: nums = [4,3,2,7,8,2,3,1]
Output: [5,6]
Example 2:

Input: nums = [1,1]
Output: [2]
 

Constraints:

n == nums.length
1 <= n <= 105
1 <= nums[i] <= n
 

Follow up: Could you do it without extra space and in O(n) runtime? You may assume the returned list does not count as extra space.
*/

// another array approach without extra space
func findDisappearedNumbers(nums []int) []int {
    arr := make([]int, len(nums))
    
    for _,v := range nums {
        arr[v-1] = v // [1 2 2] = [1,2,0]
    }

    writer := 0 
    for i,v := range arr {
        if v == 0 {
            arr[writer] = i+1
            writer++
        }
    }
    return arr[:writer]
}


// array approach without extra space
// func findDisappearedNumbers(nums []int) []int {
//     var missingNums []int
//     // mark the availble nums
//     for i:=0; i<len(nums);i++ {
//         idx := abs(nums[i]) - 1 // using abs as we are marking them -ve and using -1 to start index from zero
//         if nums[idx] > 0 { // checking for visited
//             nums[idx] = -nums[idx]
//         }
//     }

//     for i:=0; i<len(nums);i++ {
//         if nums[i] > 0 {
//             missingNums = append(missingNums, i+1)
//         }
//     }
//     return missingNums
// }

// // Helper function to get absolute value
// func abs(x int) int {
//     if x < 0 {
//         return -x
//     }
//     return x
// }
// array approach with extra space
// func findDisappearedNumbers(nums []int) []int {
//     visited := make([]bool, len(nums)+1) // as 1 <= nums[i] <= n so we'll start from 1
//     missingNums := []int{}

//     for i:=0; i<len(nums);i++ {
//        visited[nums[i]] = true 
//     }

//     for i:=1; i<=len(nums); i++ {
//         if !visited[i] {
//             missingNums = append(missingNums, i)
//         }
//     }
//     return missingNums
// }

// HashMap approach
// func findDisappearedNumbers(nums []int) []int {
//     hashMap := make(map[int]struct{})
//     missingNums := []int{}
//     // storing unique in the map
//     for i:=0; i<len(nums); i++ {
//         hashMap[nums[i]] = struct{}{}
//     }

//     // finding missing numbers
//     for i:=1; i<=len(nums); i++ {
//         if _,ok:=hashMap[i];!ok {
//             missingNums = append(missingNums, i)
//         }
//     }
//     return missingNums
// }

// brute force
// func findDisappearedNumbers(nums []int) []int {
//     if len(nums) == 0 {
//         return []int{}
//     }
//     missingNums := []int{}
//     // sorting the array
//     sort.Ints(nums)
//     //fmt.Println(nums)
//     // removing duplicate from the array
//     reader := 1
//     writer := 0
//     for reader < len(nums) {
//         //fmt.Println(nums[reader], nums[writer])
//         if nums[reader] != nums[writer] {
//             writer++
//             nums[writer] = nums[reader]
//         }
//         reader++
//     }
//     //fmt.Println(nums)
//     reader = 0
//     for i:=1; i<=len(nums);i++{
//         if nums[reader] != i {
//             missingNums = append(missingNums, i)
//         } else {reader++}
//     }
//     return missingNums
// }

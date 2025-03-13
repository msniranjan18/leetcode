/*
Given an integer array nums and an integer k, return true if there are two distinct indices i and j in the array such that nums[i] == nums[j] and abs(i - j) <= k.

 

Example 1:

Input: nums = [1,2,3,1], k = 3
Output: true
Example 2:

Input: nums = [1,0,1,1], k = 1
Output: true
Example 3:

Input: nums = [1,2,3,1,2,3], k = 2
Output: false
 

Constraints:

1 <= nums.length <= 105
-109 <= nums[i] <= 109
0 <= k <= 105
*/


func containsNearbyDuplicate(nums []int, k int) bool {
    map1 := make(map[int]int)
    for i, v := range nums {
        
        if lastIndex, exist := map1[v]; exist && i - lastIndex <= k {
            return true
        }
        map1[v] = i
    }
    return false
}

// func containsNearbyDuplicate(nums []int, k int) bool {
//     map1 := make(map[int]int)
//     for i, v := range nums {
//         val, ok := map1[v]
//         if !ok {
//             map1[v] = i
//         } else {
//             //fmt.Println("in else case", i,val, k)
//             if i-val <= k {
//                 //fmt.Println("in  i-val <= k case")
//                 return true
//             } else {
//                 map1[v] = i
//             }
//         }
//         //fmt.Println(v)
//         //fmt.Printf("%+v\n",map1)
//     }
//     return false
// }

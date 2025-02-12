/*
Given two integer arrays nums1 and nums2, return an array of their 
intersection
. Each element in the result must be unique and you may return the result in any order.

 

Example 1:

Input: nums1 = [1,2,2,1], nums2 = [2,2]
Output: [2]
Example 2:

Input: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
Output: [9,4]
Explanation: [4,9] is also accepted.
 

Constraints:

1 <= nums1.length, nums2.length <= 1000
0 <= nums1[i], nums2[i] <= 1000
*/


func intersection(nums1 []int, nums2 []int) []int {
    size := 1001 // as given constraint 0 <= nums1[i], nums2[i] <= 1000
    present1 := make([]bool,size)
    present2 := make([]bool,size)

    for _, n := range nums1 {
        present1[n] = true
    }
    for _, n := range nums2 {
        present2[n] = true
    }

    var result []int
    for i:=0;i<size;i++ {
        if present1[i] && present2[i] {
            result = append(result, i)
        }
    }
    return result
}


// func intersection(nums1 []int, nums2 []int) []int {
//     set := make(map[int]struct{})

//     for _, n := range nums1 {
//         if _, ok := set[n]; !ok {
//             set[n] = struct{}{}
//         }
//     }

//     var result []int
//     for _, n := range nums2 {
//         if _, ok := set[n]; ok {
//             result = append(result, n)
//             delete(set, n) // remove to avoid duplicacy
//         }
//     }
//     return result
// }

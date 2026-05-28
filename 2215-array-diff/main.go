package main

import "fmt"

func findDifference(nums1 []int, nums2 []int) [][]int {
    const offset = 1000      // Offset to convert negative values to non-negative indices.
    const size = 2001        // Since values are in [-1000, 1000].

    // Use slices of booleans for presence; their size is fixed.
    present1 := make([]bool, size)
    present2 := make([]bool, size)

    // Mark numbers present in nums1
    for _, n := range nums1 {
        present1[n+offset] = true
    }
    // Mark numbers present in nums2
    for _, n := range nums2 {
        present2[n+offset] = true
    }

    res1 := []int{}
    res2 := []int{}

    // Check for numbers present in one array but not in the other.
    for i := 0; i < size; i++ {
        if present1[i] && !present2[i] {
            res1 = append(res1, i-offset)
        }
        if present2[i] && !present1[i] {
            res2 = append(res2, i-offset)
        }
    }

    return [][]int{res1, res2}
}



// func findDifference(nums1 []int, nums2 []int) [][]int {
//     map1 := make(map[int]bool)
//     map2 := make(map[int]bool)

//     // Populate sets for nums1 and nums2
//     for _, n := range nums1 {
//         map1[n] = true
//     }
//     for _, n := range nums2 {
//         map2[n] = true
//     }

//     // Find elements in nums1 but not in nums2
//     result1 := make([]int, 0)
//     for num := range map1 {
//         if !map2[num] {
//             result1 = append(result1, num)
//         }
//     }

//     // Find elements in nums2 but not in nums1
//     result2 := make([]int, 0)
//     for num := range map2 {
//         if !map1[num] {
//             result2 = append(result2, num)
//         }
//     }

//     return [][]int{result1, result2}
// }


// func findDifference(nums1 []int, nums2 []int) [][]int {
//     map1 := make(map[int]struct{})
//     map2 := make(map[int]struct{})
//     var result [][]int
//     for _, n := range nums1 {
//         if _, ok := map1[n]; !ok {
//             map1[n] = struct{}{}
//         }
//     }
//     for _, n := range nums2 {
//         if _, ok := map2[n]; !ok {
//             map2[n] = struct{}{}
//         }
//     }

//     var temp []int
//     for key,_ := range map1 {
//         if _, ok := map2[key]; !ok {
//             temp = append(temp, key)
//         }        
//     }
//     result = append(result, temp)
//     temp = []int{}

//     for key,_ := range map2 {
//         if _, ok := map1[key]; !ok {
//             temp = append(temp, key)
//         }        
//     }
//     result = append(result, temp)

//     return result
// }
/*
Given two 0-indexed integer arrays nums1 and nums2, return a list answer of size 2 where:

answer[0] is a list of all distinct integers in nums1 which are not present in nums2.
answer[1] is a list of all distinct integers in nums2 which are not present in nums1.
Note that the integers in the lists may be returned in any order.

 

Example 1:

Input: nums1 = [1,2,3], nums2 = [2,4,6]
Output: [[1,3],[4,6]]
Explanation:
For nums1, nums1[1] = 2 is present at index 0 of nums2, whereas nums1[0] = 1 and nums1[2] = 3 are not present in nums2. Therefore, answer[0] = [1,3].
For nums2, nums2[0] = 2 is present at index 1 of nums1, whereas nums2[1] = 4 and nums2[2] = 6 are not present in nums2. Therefore, answer[1] = [4,6].
Example 2:

Input: nums1 = [1,2,3,3], nums2 = [1,1,2,2]
Output: [[3],[]]
Explanation:
For nums1, nums1[2] and nums1[3] are not present in nums2. Since nums1[2] == nums1[3], their value is only included once and answer[0] = [3].
Every integer in nums2 is present in nums1. Therefore, answer[1] = [].
 

Constraints:

1 <= nums1.length, nums2.length <= 1000
-1000 <= nums1[i], nums2[i] <= 1000
*/

package main

import "fmt"

func main() {
    nums1 := []int{1, 2, 3}
    nums2 := []int{2, 4, 6}
    fmt.Println(findDifference(nums1, nums2)) // Expected: [[1, 3], [4, 6]]
}

/*
Why Maps Might Be Slower
When you use maps in Go, even though their average time complexity for lookups and inserts is O(1), there are some constant‑factor overheads:

Hashing Overhead: Each lookup or insertion must compute a hash of the key.
Memory Allocation: Maps require dynamic memory allocation and may involve rehashing or resizing.
Indirection: Maps involve extra pointer dereferencing compared to direct array indexing.
Given the constraints of the problem:

Array Length: Up to 1000 elements.
Value Range: Each integer is between -1000 and 1000.
You can leverage the limited value range by using an array (or slice) of fixed size (2001 elements) to track the presence of numbers instead of using maps. 
This approach uses constant extra space (O(1) with respect to input size, since 2001 is a constant) and typically runs faster due to direct index access.


Why This is Optimal
Time Complexity:
Marking presence in each array takes O(n) time (n ≤ 1000).
Iterating over the fixed size (2001) array is O(1) relative to the constraints.

Space Complexity:
The extra space used is O(2001), which is a constant given the problem constraints.
Lower Constant Factors:

Array indexing is very fast and has minimal overhead compared to map operations.

Summary
Map-Based Approach:
Time: O(n + m)
Space: O(n + m)
Overhead: Hash computations, memory allocations, indirections.
Array-Based Approach (Optimal for This Problem):

Time: O(n + constant)
Space: O(1) (with a fixed size of 2001 booleans)
Lower overhead due to direct indexing.
This is why, under the given constraints, the array‑based solution is optimal and often faster than the map‑based approach. 
*/

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

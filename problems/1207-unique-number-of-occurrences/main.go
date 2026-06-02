/*
Given an array of integers arr, return true if the number of occurrences of each value in the array is unique or false otherwise.

 

Example 1:

Input: arr = [1,2,2,1,1,3]
Output: true
Explanation: The value 1 has 3 occurrences, 2 has 2 and 3 has 1. No two values have the same number of occurrences.
Example 2:

Input: arr = [1,2]
Output: false
Example 3:

Input: arr = [-3,0,1,-3,1,1,1,-3,10,0]
Output: true
 

Constraints:

1 <= arr.length <= 1000
-1000 <= arr[i] <= 1000
*/

func uniqueOccurrences(arr []int) bool {
    set1 := make(map[int]int)
    for _, n := range arr {
        set1[n]++
    }

    offset := 1000
    size := 2001 
    seen := make([]bool, size)
    for _, val := range set1 {
        if seen[offset + val] {
            return false
        } else {
            seen[offset + val] = true
        }
    }
    return true
}

/*
Time Complexity:
✅ O(n) (same as your solution)
✅ Space Complexity: O(n) → O(1)
(Since the seen array has a fixed size of 2001, it doesn't grow with input size.)

🔹 Pros: Faster lookup than a hash map.
🔹 Cons: Uses a fixed-size array instead of a dynamic map.
*/

// func uniqueOccurrences(arr []int) bool {
//     set1 := make(map[int]int)
//     for _, n := range arr {
//         set1[n]++
//     }

//     set2 := make(map[int]struct{})
//     for _, val := range set1 {
//         if _, ok:= set2[val]; !ok {
//             set2[val] = struct{}{}
//         } else {
//             return false
//         }
//     }
//     return true
// }

/*
Time Complexity Analysis
✅ Total Time Complexity: 𝑂(𝑛)+𝑂(𝑛)=𝑂(𝑛)
✅ Space Complexity: 𝑂(𝑛)
O(n) (since set1 and set2 store at most n elements)
*/


/*
🔹 For General Use: Stick with your hash map approach (best balance).
🔹 If Memory is a Concern: Use the sorting approach.
🔹 If Speed is Critical & Constraints Allow: Use the boolean array approach.
*/

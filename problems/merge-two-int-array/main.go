package main

import "fmt"

func main() {
	nums1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19}
	nums2 := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23, 25, 27, 30}
	fmt.Println(nums1, " + ", nums2, " = ", mergeGeneric(nums1, len(nums1), nums2, len(nums2)))

	nums1 = []int{1, 12, 13, 14, 15, 61}
	nums2 = []int{0, 1, 3, 5}
	fmt.Println(nums1, " + ", nums2, " = ", mergeGeneric(nums1, len(nums1), nums2, len(nums2)))

	nums1 = []int{1}
	nums2 = []int{0, 1, 3, 5}
	fmt.Println(nums1, " + ", nums2, " = ", mergeGeneric(nums1, len(nums1), nums2, len(nums2)))

	nums2 = []int{1}
	nums1 = []int{0, 1, 3, 5}
	fmt.Println(nums1, " + ", nums2, " = ", mergeGeneric(nums1, len(nums1), nums2, len(nums2)))

	nums1 = []int{0}
	nums2 = []int{1}
	fmt.Println(nums1, " + ", nums2, " = ", mergeGeneric(nums1, len(nums1), nums2, len(nums2)))

	nums1 = []int{1}
	nums2 = []int{0}
	fmt.Println(nums1, " + ", nums2, " = ", mergeGeneric(nums1, len(nums1), nums2, len(nums2)))

	nums1 = []int{}
	nums2 = []int{0, 2}
	fmt.Println(nums1, " + ", nums2, " = ", mergeGeneric(nums1, len(nums1), nums2, len(nums2)))

	nums1 = []int{1}
	nums2 = []int{}
	fmt.Println(nums1, " + ", nums2, " = ", mergeGeneric(nums1, len(nums1), nums2, len(nums2)))

}

func mergeGeneric(nums1 []int, m int, nums2 []int, n int) []int {

	var small []int
	var big []int
	res := []int{}
	if m <= n {
		small = nums1
		big = nums2
	} else {
		small = nums2
		big = nums1
	}

	i, j := 0, 0
	for i < len(small) && j < len(big) {
		if small[i] <= big[j] {
			res = append(res, small[i])
			i++
		} else {
			res = append(res, big[j])
			j++
		}
	}
	res = append(res, small[i:]...)
	res = append(res, big[j:]...)
	return res
}

/*
Output:
[1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19]  +  [1 3 5 7 9 11 13 15 17 19 21 23 25 27 30]  =  [1 1 2 3 3 4 5 5 6 7 7 8 9 9 10 11 11 12 13 13 14 15 15 16 17 17 18 19 19 21 23 25 27 30]
[1 12 13 14 15 61]  +  [0 1 3 5]  =  [0 1 1 3 5 12 13 14 15 61]
[1]  +  [0 1 3 5]  =  [0 1 1 3 5]
[0 1 3 5]  +  [1]  =  [0 1 1 3 5]
[0]  +  [1]  =  [0 1]
[1]  +  [0]  =  [0 1]
[]  +  [0 2]  =  [0 2]
[1]  +  []  =  [1]
*/

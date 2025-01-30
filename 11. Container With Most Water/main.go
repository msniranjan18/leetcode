/*
You are given an integer array height of length n. There are n vertical lines drawn such that the two endpoints of the ith line are (i, 0) and (i, height[i]).

Find two lines that together with the x-axis form a container, such that the container contains the most water.

Return the maximum amount of water a container can store.

Notice that you may not slant the container.



Example 1:


Input: height = [1,8,6,2,5,4,8,3,7]
Output: 49
Explanation: The above vertical lines are represented by array [1,8,6,2,5,4,8,3,7]. In this case, the max area of water (blue section) the container can contain is 49.
Example 2:

Input: height = [1,1]
Output: 1


Constraints:

n == height.length
2 <= n <= 105
0 <= height[i] <= 104
*/

package main

func main() {

	s := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}

	println(maxArea(s))
}

func maxArea(height []int) int {
	max := 0
	for i, j := 0, len(height)-1; i < j; {
		if height[i] > height[j] {
			if max < (j-i)*height[j] {
				max = (j - i) * height[j]
			}
			j--
		} else {
			if max < (j-i)*height[i] {
				max = (j - i) * height[i]
			}
			i++
		}

	}
	return max
}

/*
// More readable code; using built-in func
func maxArea(height []int) int {
	maximum := 0
        minmum := 0
	for i, j := 0, len(height)-1; i < j; {
		minmum = min(height[i], height[j])
		maximum = max(maximum, (j-i)*minmum)
		if height[i] > height[j] {
			j--
		} else {
			i++
		}

	}
	return maximum
}
*/

/*
Output:
49
*/

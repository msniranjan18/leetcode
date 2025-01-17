/*
Our hero Teemo is attacking an enemy Ashe with poison attacks! When Teemo attacks Ashe, Ashe gets poisoned for a exactly duration seconds. More formally, an attack at second t will mean Ashe is poisoned during the inclusive time interval [t, t + duration - 1]. If Teemo attacks again before the poison effect ends, the timer for it is reset, and the poison effect will end duration seconds after the new attack.

You are given a non-decreasing integer array timeSeries, where timeSeries[i] denotes that Teemo attacks Ashe at second timeSeries[i], and an integer duration.

Return the total number of seconds that Ashe is poisoned.

 

Example 1:

Input: timeSeries = [1,4], duration = 2
Output: 4
Explanation: Teemo's attacks on Ashe go as follows:
- At second 1, Teemo attacks, and Ashe is poisoned for seconds 1 and 2.
- At second 4, Teemo attacks, and Ashe is poisoned for seconds 4 and 5.
Ashe is poisoned for seconds 1, 2, 4, and 5, which is 4 seconds in total.
Example 2:

Input: timeSeries = [1,2], duration = 2
Output: 3
Explanation: Teemo's attacks on Ashe go as follows:
- At second 1, Teemo attacks, and Ashe is poisoned for seconds 1 and 2.
- At second 2 however, Teemo attacks again and resets the poison timer. Ashe is poisoned for seconds 2 and 3.
Ashe is poisoned for seconds 1, 2, and 3, which is 3 seconds in total.
 

Constraints:

1 <= timeSeries.length <= 104
0 <= timeSeries[i], duration <= 107
timeSeries is sorted in non-decreasing order.
*/

package main

import "fmt"

func main() {
	testcases := [][]int{
		{1, 4}, {2},
		{1, 2}, {4},
		{2, 4, 6}, {6},
		{3, 8, 13}, {6},
		{4, 10, 20}, {6},
		{10}, {6},
	}

	for i := 0; i < len(testcases); i += 2 {
		fmt.Printf("testcase: timeSeries=%+v, duration=%+v, totalPoisonedDuration:%+v \n", testcases[i], testcases[i+1], findPoisonedDuration(testcases[i], testcases[i+1][0]))
	}

}

func findPoisonedDuration(timeSeries []int, duration int) int {
	totalPoisonedDuration := 0
	if len(timeSeries) == 1 {
		return duration
	} else {
		totalPoisonedDuration += duration
	}
	for i := 1; i < len(timeSeries); i++ {
		if timeSeries[i-1]+duration-1 >= timeSeries[i] {
			totalPoisonedDuration += timeSeries[i] - timeSeries[i-1]
		} else {
			totalPoisonedDuration += duration
		}
	}
	return totalPoisonedDuration
}

/*
Output:
testcase: timeSeries=[1 4], duration=[2], totalPoisonedDuration:4 
testcase: timeSeries=[1 2], duration=[4], totalPoisonedDuration:5 
testcase: timeSeries=[2 4 6], duration=[6], totalPoisonedDuration:10 
testcase: timeSeries=[3 8 13], duration=[6], totalPoisonedDuration:16 
testcase: timeSeries=[4 10 20], duration=[6], totalPoisonedDuration:18 
testcase: timeSeries=[10], duration=[6], totalPoisonedDuration:6 
*/

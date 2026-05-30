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

type TestCase struct {
	TimeSeries []int
	Duration   int
	Expected   int
}

func main() {
	testCases := []TestCase{
		{TimeSeries: []int{1, 4}, Duration: 2, Expected: 4},
		{TimeSeries: []int{1, 2}, Duration: 4, Expected: 5},
		{TimeSeries: []int{2, 4, 6}, Duration: 6, Expected: 10},
		{TimeSeries: []int{3, 8, 13}, Duration: 6, Expected: 16},
		{TimeSeries: []int{4, 10, 20}, Duration: 6, Expected: 18},
		{TimeSeries: []int{10}, Duration: 6, Expected: 6},
	}

	for _, testCase := range testCases {
		result := findPoisonedDuration(testCase.TimeSeries, testCase.Duration)
		if result == testCase.Expected {
			fmt.Printf("Test Case: [TimeSeries=%+v, Duration=%d] PASSED: output: %d, expected %d\n", testCase.TimeSeries, testCase.Duration, result, testCase.Expected)
		} else {
			fmt.Printf("Test Case: [TimeSeries=%+v, Duration=%d] FAILED: output: %d, expected %d\n", testCase.TimeSeries, testCase.Duration, result, testCase.Expected)
		}
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
Test Case: [TimeSeries=[1 4], Duration=2] PASSED: output: 4, expected 4
Test Case: [TimeSeries=[1 2], Duration=4] PASSED: output: 5, expected 5
Test Case: [TimeSeries=[2 4 6], Duration=6] PASSED: output: 10, expected 10
Test Case: [TimeSeries=[3 8 13], Duration=6] PASSED: output: 16, expected 16
Test Case: [TimeSeries=[4 10 20], Duration=6] PASSED: output: 18, expected 18
Test Case: [TimeSeries=[10], Duration=6] PASSED: output: 6, expected 6
*/

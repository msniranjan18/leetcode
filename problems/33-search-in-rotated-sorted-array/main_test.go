package main

import (
	"testing"
	"github.com/msniranjan18/leetcode/util"
)

func TestSearchInARotatedSortedArray(t *testing.T) {
	testcases := []struct{
		input []int
		target int
		expected int
	}{
		{
			input: []int{4,5,6,7,0,1,2},
			target: 0,
			expected: 4,
		},
		{
			input: []int{3,1},
			target: 1,
			expected: 1,
		},
		{
			input: []int{1},
			target: 0,
			expected: -1,
		},
	}

	for _, tc := range testcases {
		t.Run("TestingSearch", func(t *testing.T){
			actual := search(tc.input, tc.target)
			if actual != tc.expected {
				t.Errorf("\nTest 'n:%d' Failed:\nExpected: %v\nActual:   %v",tc.input, tc.expected, actual)
			}
		})
	}
}

func BenchmarkSearch(b *testing.B) {
	//largeInput := []int{4,5,6,7,0,1,2}
	largeInput := util.BigIntArray
	target:= 4
	b.ResetTimer()
	for i:=0; i<b.N; i++ {
		search(largeInput, target)
	}
}
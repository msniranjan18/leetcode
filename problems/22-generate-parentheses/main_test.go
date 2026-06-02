package main

import (
	"testing"
)

func TestGenerateParentheses(t *testing.T) {
	testcases := []struct{
		n int
		expected int
	}{
		{
			n: 3,
			expected: 5,
		},
		{
			n: 8,
			expected: 1430,
		},
		{
			n: 10,
			expected: 16796,
		},
		{
			n: 12,
			expected: 208012,
		},
		{
			n: 14,
			expected: 2674440,
		},				
	}

	for _, tc := range testcases {
		t.Run("Testing-Generating-Parentheses", func(t *testing.T){
			actual:=generateParenthesisTest(tc.n)
			if actual != tc.expected {
				t.Errorf("\nTest 'n:%d' Failed:\nExpected: %v\nActual:   %v",tc.n, tc.expected, actual)
			}
		})
	}
}

func BenchmarkGenerateParentheses(b *testing.B) {
	largeInput := 16
	b.ResetTimer()

	for i:=0; i<b.N; i++ {
		generateParenthesisTest(largeInput)
	}
}
package main

import (
	"fmt"
)

func main() {
	fmt.Println(minCosts([]int{5,3,4,1,3,2}))
}

func minCosts(cost []int) []int {
    n:=len(cost)
    out := make([]int,n)
    minCostSoFar := 101 //as per given constaraint 1 <= cost[i] <= 100
    for i:=0; i<n; i++ {
        minCostSoFar = min(minCostSoFar, cost[i])
        out[i] = minCostSoFar
    }
    return out
}
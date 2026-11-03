package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minScore(4, [][]int{{1, 2, 9}, {2, 3, 6}, {2, 4, 5}, {1, 4, 7}}))
	fmt.Println(minScore(4, [][]int{{1, 2, 2}, {1, 3, 4}, {3, 4, 7}}))
}

func minScore(n int, roads [][]int) int {
	// create edjacent nodes
	type Edge struct {
		to   int
		dist int
	}
	graph := make([][]Edge, n+1)
	for _, r := range roads {
		src, dest, dist := r[0], r[1], r[2]
		graph[src] = append(graph[src], Edge{to: dest, dist: dist})
		graph[dest] = append(graph[dest], Edge{to: src, dist: dist})
	}

	visited := make([]bool, n+1)
	res := math.MaxInt
	var dfs func(i int)
	dfs = func(i int) {
		if visited[i] == true {
			return
		}

		visited[i] = true

		for _, neibour := range graph[i] {
			res = min(res, neibour.dist)
			dfs(neibour.to)
		}
	}

	dfs(1)
	return res

}

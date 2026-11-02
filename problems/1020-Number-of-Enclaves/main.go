package main

import (
	"fmt"
)

func main() {
	fmt.Println(numEnclaves([][]int{{0, 0, 0, 0}, {1, 0, 1, 0}, {0, 1, 1, 0}, {0, 0, 0, 0}}))
	fmt.Println(numEnclaves([][]int{{0, 1, 1, 0}, {0, 0, 1, 0}, {0, 0, 1, 0}, {0, 0, 0, 0}}))

}

func numEnclaves(grid [][]int) int {
	n := len(grid)
	m := len(grid[0])

	// visted flags
	visited := make([][]bool, n)
	for i := range visited {
		visited[i] = make([]bool, m)
	}

	var dfs func(r, c int) int
	dfs = func(r, c int) int {
		if r < 0 || c < 0 || r == n || c == m || grid[r][c] == 0 || visited[r][c] == true {
			return 0
		}

		res := 1
		visited[r][c] = true
		res += dfs(r+1, c)
		res += dfs(r-1, c)
		res += dfs(r, c+1)
		res += dfs(r, c-1)
		return res
	}

	boarderLand := 0
	totalLand := 0
	for r := 0; r < n; r++ {
		for c := 0; c < m; c++ {
			if grid[r][c] == 1 &&
				(r == 0 || r == n-1 ||
					c == 0 || c == m-1) {
				boarderLand += dfs(r, c)
			}
			if grid[r][c] == 1 {
				totalLand += 1
			}
		}
	}
	return totalLand - boarderLand
}

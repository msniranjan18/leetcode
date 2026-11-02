package main

import "fmt"

func main() {
	fmt.Println(closedIsland([][]int{{1, 1, 1, 1, 1, 1, 1, 0}, {1, 0, 0, 0, 0, 1, 1, 0}, {1, 0, 1, 0, 1, 1, 1, 0}, {1, 0, 0, 0, 0, 1, 0, 1}, {1, 1, 1, 1, 1, 1, 1, 0}}))
}

func closedIsland(grid [][]int) int {

	n := len(grid)
	m := len(grid[0])

	visited := make([][]bool, n)
	for i := range visited {
		visited[i] = make([]bool, m)
	}
	res := 0

	var dfs func(r, c int) int
	dfs = func(r, c int) int {
		if r < 0 || c < 0 || r == n || c == m {
			return 0 // false
		}

		if visited[r][c] == true || grid[r][c] == 1 {
			return 1 // true
		}
		visited[r][c] = true    // mark visited
		return min(dfs(r+1, c), // down
			dfs(r-1, c), // up
			dfs(r, c+1), // right
			dfs(r, c-1), // left
		)
	}

	for r := 0; r < n; r++ {
		for c := 0; c < m; c++ {
			if visited[r][c] == false && grid[r][c] == 0 {
				res += dfs(r, c)
			}

		}
	}
	return res
}

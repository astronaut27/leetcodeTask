# Intuition
The problem of finding the shortest path through a binary matrix where 1s represent obstacles and 0s represent free cells can be approached as a graph traversal problem. Each cell in the matrix can be seen as a node, and the goal is to reach the bottom-right corner from the top-left corner using the shortest path that only traverses cells containing 0. Immediate thoughts lead to using BFS, as it naturally finds the shortest path in an unweighted graph, which this grid can be thought of when considering each step from one cell to its neighbor has equal weight.

# Approach
1. **BFS Initialization**:
    - Start BFS from the top-left corner of the grid if it's not blocked.
    - Use a queue to facilitate the BFS process, starting with the initial cell (0, 0) if it's accessible (i.e., `grid[0][0] == 0`).

2. **BFS Expansion**:
    - For each cell, attempt to move in all 8 possible directions (including diagonals).
    - Check boundaries and whether the cell in the intended direction is blocked or has been visited.
    - If reachable and free, add the new cell to the queue for subsequent processing.

3. **Path Tracking**:
    - Maintain a path length counter that increments with each layer of BFS, representing the number of steps taken from the start.
    - Once the bottom-right corner is reached, return the current path length.

4. **Termination**:
    - If the queue is exhausted without reaching the bottom-right corner, return -1, indicating no available path.

# Complexity
- **Time complexity**: \(O(n^2)\), where \(n\) is the dimension of the matrix. In the worst case, every cell is processed once.
- **Space complexity**: \(O(n^2)\) for storing the visited matrix and the BFS queue, which in the worst case could contain all cells of the grid.

# Code
```golang
package main

import "fmt"

type Point struct {
	x, y int
}

func shortestPathBinaryMatrix(grid [][]int) int {
	n := len(grid)

	if grid[0][0] != 0 || grid[n-1][n-1] != 0 {
		return -1
	}

	directions := []Point{
		{0, 1}, {1, 0}, {0, -1}, {-1, 0},
		{1, 1}, {-1, -1}, {1, -1}, {-1, 1},
	}

	visited := make([][]bool, n)
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	queue := []Point{{0, 0}}
	visited[0][0] = true
	pathLength := 1

	for len(queue) > 0 {
		nextQueue := []Point{}

		for _, cell := range queue {
			x, y := cell.x, cell.y

			if x == n-1 && y == n-1 {
				return pathLength
			}

			for _, dir := range directions {
				nx, ny := x+dir.x, y+dir.y

				if nx < 0 || ny < 0 || nx >= n || ny >= n {
					continue
				}

				if grid[nx][ny] != 0 || visited[nx][ny] {
					continue
				}

				visited[nx][ny] = true
				nextQueue = append(nextQueue, Point{nx, ny})
			}
		}

		queue = nextQueue
		pathLength++
	}

	return -1
}

func main() {
	grid := [][]int{{0, 1}, {1, 0}}
	fmt.Println("Shortest path length:", shortestPathBinaryMatrix(grid)) // Output: 2
}
```

### Explanation
- **BFS Process**: The BFS explores the matrix layer by layer, ensuring the shortest path is found due to the nature of BFS where each node (cell) is processed level by level.
- **Edge Cases**: Handles cases where the start or end is blocked right at the beginning to avoid unnecessary processing.
- **Efficiency**: The algorithm is efficient within the constraints, given that it only processes each cell once and uses straightforward operations to check and enqueue neighboring cells.

This solution ensures a systematic exploration of the matrix, finding the shortest possible path while adhering to the constraints imposed by obstacles.

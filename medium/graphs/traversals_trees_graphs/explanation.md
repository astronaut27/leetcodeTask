# Intuition
The problem is about finding the shortest path from a given entrance in a maze to the nearest exit, where exits are defined as any empty cell on the border that is not the entrance. Since we need the **shortest** path and all moves have equal weight (one step), Breadth-First Search (BFS) is the natural choice.

# Approach
1. **Model the Maze**:
    - The maze is a 2D grid with `'.'` representing open paths and `'+'` representing walls.
    - We define valid exits as `'.'` cells on the border of the grid that are not the entrance itself.

2. **Initialize BFS**:
    - Start from the entrance cell using a queue for BFS.
    - Use a `visited` matrix to track cells we’ve already explored to prevent infinite loops or redundant checks.

3. **BFS Traversal**:
    - At each step, we process all nodes in the current level (layer of BFS), increment the step count, and explore the four cardinal directions (up, down, left, right).
    - If we reach a border cell (not the entrance) that is a `'.'`, we return the current step count.
    - If BFS completes without finding an exit, return `-1`.

# Complexity
- **Time complexity**: $$O(n \cdot m)$$  
  Every cell in the grid is visited at most once in the worst case, where `n` is the number of rows and `m` is the number of columns.

- **Space complexity**: $$O(n \cdot m)$$  
  For the `visited` matrix and the BFS queue in the worst case.

# Code
```golang
package main

import "fmt"

type Point struct {
	x, y int
}

var directions = []Point{
	{0, 1}, {1, 0}, {0, -1}, {-1, 0},
}

func nearestExit(maze [][]byte, entrance []int) int {
	n, m := len(maze), len(maze[0])
	visited := make([][]bool, n)
	for i := range visited {
		visited[i] = make([]bool, m)
	}
	visited[entrance[0]][entrance[1]] = true
	return bfs(maze, entrance[0], entrance[1], visited)
}

func bfs(maze [][]byte, startX, startY int, visited [][]bool) int {
	n, m := len(maze), len(maze[0])
	queue := []Point{{startX, startY}}
	steps := 0

	for len(queue) > 0 {
		for size := len(queue); size > 0; size-- {
			curr := queue[0]
			queue = queue[1:]

			x, y := curr.x, curr.y
			// Check if current cell is an exit
			if maze[x][y] == '.' &&
				(x == 0 || x == n-1 || y == 0 || y == m-1) &&
				!(x == startX && y == startY) {
				return steps
			}

			// Explore neighbors
			for _, d := range directions {
				nx, ny := x+d.x, y+d.y
				if nx >= 0 && ny >= 0 && nx < n && ny < m &&
					maze[nx][ny] == '.' && !visited[nx][ny] {
					visited[nx][ny] = true
					queue = append(queue, Point{nx, ny})
				}
			}
		}
		steps++
	}
	return -1
}

func main() {
	fmt.Println(nearestExit([][]byte{
		{'+', '+', '+'},
		{'.', '.', '.'},
		{'+', '+', '+'}}, []int{1, 0})) // Output: 2

	fmt.Println(nearestExit([][]byte{
		{'+', '+', '.', '+'},
		{'.', '.', '.', '+'},
		{'+', '+', '+', '.'}}, []int{1, 2})) // Output: 1
}
```

### Summary
- BFS is optimal for shortest path in unweighted grids like this maze.
- Exits are carefully filtered by checking border cells (excluding entrance).
- `visited` ensures that we do not reprocess or fall into loops.
- Each level of BFS increments the step counter, so the first time an exit is reached, it's guaranteed to be the shortest path.

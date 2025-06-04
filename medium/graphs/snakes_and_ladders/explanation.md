# Intuition
The Snakes and Ladders board is essentially a graph where each square represents a node, and a dice roll represents an edge to the next 1 to 6 positions. If a player lands on a square with a ladder or snake, they are immediately transported to another square, forming a jump in the graph. To find the minimum number of dice rolls to reach the final square, we can perform a Breadth-First Search (BFS) starting from square 1.

# Approach
1. **Flatten the 2D Board**:
    - Convert the 2D `board` into a 1D array to simulate moves linearly.
    - Account for the zig-zag pattern in rows (left-to-right and right-to-left alternation).

2. **Breadth-First Search (BFS)**:
    - Initialize a queue starting at position 1.
    - Use a `visited` array to prevent revisiting squares.
    - For each position, simulate all possible dice rolls (1 to 6).
    - If the result of a dice roll lands on a ladder or snake (`flatten[next] != -1`), jump to that position.
    - If the final square `n*n` is reached, return the number of moves taken.

3. **Termination**:
    - If all possible paths are exhausted and the final square was not reached, return -1.

# Complexity
- **Time complexity**: $$O(n^2)$$
    - We may visit each square once, and from each square, try up to 6 moves.
- **Space complexity**: $$O(n^2)$$
    - To store the `flatten` board and the `visited` array.

# Code
```golang
package main

func snakesAndLadders(board [][]int) int {
	n := len(board)

	// Flatten the board in a zig-zag manner
	flatten := make([]int, n*n+1)
	idx := 1
	leftToRight := true
	for i := n - 1; i >= 0; i-- {
		if leftToRight {
			for j := 0; j < n; j++ {
				flatten[idx] = board[i][j]
				idx++
			}
		} else {
			for j := n - 1; j >= 0; j-- {
				flatten[idx] = board[i][j]
				idx++
			}
		}
		leftToRight = !leftToRight
	}

	visited := make([]bool, n*n+1)
	queue := []int{1}
	visited[1] = true
	moves := 0

	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			curr := queue[0]
			queue = queue[1:]

			if curr == n*n {
				return moves
			}

			for dice := 1; dice <= 6; dice++ {
				next := curr + dice
				if next > n*n {
					continue
				}
				if flatten[next] != -1 {
					next = flatten[next]
				}
				if !visited[next] {
					visited[next] = true
					queue = append(queue, next)
				}
			}
		}
		moves++
	}
	return -1
}
```

### Summary
- This problem is a perfect fit for BFS because we need the shortest number of moves.
- Flattening the board simplifies the traversal logic.
- The presence of snakes and ladders is handled via direct jumps in the BFS.
- The algorithm ensures each position is visited only once to avoid cycles and redundant work.

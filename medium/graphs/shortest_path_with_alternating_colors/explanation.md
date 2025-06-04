# Intuition
The task is to find the shortest path from node 0 to every other node in a directed graph with red and blue edges, under the constraint that the edge colors must alternate along the path. The alternating edge condition adds a twist to a typical shortest path problem, hinting at the need for tracking additional state during traversal — specifically, the color of the last edge used.

# Approach
1. **Graph Construction**:
    - Create two separate adjacency lists: one for red edges and one for blue edges.
    - This separation makes it easy to switch between red and blue edges when traversing.

2. **BFS Traversal**:
    - Use Breadth-First Search (BFS) because we are looking for the shortest path in an unweighted graph.
    - Each state in the queue holds the current node, the color of the last edge used, and the distance from the start node.

3. **Alternating Edges**:
    - From a node reached by a red edge, only explore blue edges next, and vice versa.
    - Use a visited matrix to track if a node has already been visited via a specific edge color to prevent cycles and redundant processing.

4. **Tracking Results**:
    - For each node, update the result array with the shortest distance when first reached or when a shorter alternating path is found.

# Complexity
- **Time complexity**: $$O(n + r + b)$$  
  Where \(n\) is the number of nodes, \(r\) is the number of red edges, and \(b\) is the number of blue edges. Each edge is visited at most once for each color transition.
- **Space complexity**: $$O(n + r + b)$$  
  For the adjacency lists, visited matrix, and BFS queue.

# Code
```golang
const (
	RED  = 0
	BLUE = 1
)

type State struct {
	node  int
	color int // RED=0, BLUE=1
	dist  int
}

func shortestAlternatingPaths(n int, redEdges [][]int, blueEdges [][]int) []int {
	// Build separate graphs for red and blue edges
	redGraph := make(map[int][]int)
	blueGraph := make(map[int][]int)

	for _, edge := range redEdges {
		redGraph[edge[0]] = append(redGraph[edge[0]], edge[1])
	}
	for _, edge := range blueEdges {
		blueGraph[edge[0]] = append(blueGraph[edge[0]], edge[1])
	}

	// visited[node][color] = whether the node was visited using a specific color
	visited := make([][2]bool, n)
	ans := make([]int, n)
	for i := range ans {
		ans[i] = -1
	}

	// Start BFS with both red and blue edges from node 0
	queue := []State{
		{0, RED, 0},
		{0, BLUE, 0},
	}
	visited[0][RED] = true
	visited[0][BLUE] = true
	ans[0] = 0

	for len(queue) > 0 {
		state := queue[0]
		queue = queue[1:]

		node := state.node
		color := state.color
		dist := state.dist

		if ans[node] == -1 || dist < ans[node] {
			ans[node] = dist
		}

		// Determine next graph and color to use
		var nextGraph map[int][]int
		var nextColor int
		if color == RED {
			nextGraph = blueGraph
			nextColor = BLUE
		} else {
			nextGraph = redGraph
			nextColor = RED
		}

		for _, neighbor := range nextGraph[node] {
			if !visited[neighbor][nextColor] {
				visited[neighbor][nextColor] = true
				queue = append(queue, State{neighbor, nextColor, dist + 1})
			}
		}
	}

	return ans
}
```

### Summary
- This solution uses a BFS to guarantee shortest path computation and adds an edge color constraint by keeping track of the last color used.
- It avoids revisiting the same node in the same context (color) to ensure efficiency.
- The alternating path requirement is elegantly handled by maintaining separate graphs and toggling between them during traversal.

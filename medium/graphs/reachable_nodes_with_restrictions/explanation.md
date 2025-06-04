# Intuition
Given a graph represented by nodes and edges along with a list of restricted nodes, the problem involves determining how many nodes can be reached from the starting node (node 0), while avoiding the restricted nodes. This is a typical graph traversal problem but with the added twist of restrictions that need to be considered.

# Approach
1. **Graph Representation**:
    - First, construct the adjacency list for the graph from the given edges, allowing easy traversal from any given node to its neighbors.

2. **Handling Restrictions**:
    - Convert the list of restricted nodes into a set for O(1) lookup times, ensuring that during traversal, we can efficiently determine if a node is restricted.

3. **Depth-First Search (DFS)**:
    - Use DFS starting from node 0 to explore all reachable nodes.
    - For each node encountered, check if it is not visited and not restricted. If so, mark it as visited and recurse into its neighbors.
    - Keep a count of all nodes visited during this process, which are not restricted.

# Complexity
- **Time complexity**: O(V + E), where `V` is the number of vertices and `E` is the number of edges. This covers building the graph and the DFS traversal.
- **Space complexity**: O(V), for storing the adjacency list, visited array, and the restricted set. Additionally, the space for the recursion stack in DFS should also be considered, which in the worst case could be O(V) if the graph is structured like a line.

# Code
```golang
package main

import "fmt"

func reachableNodes(n int, edges [][]int, restricted []int) int {
    // Build the graph using an adjacency list
    graph := make([][]int, n)
    for _, edge := range edges {
        a, b := edge[0], edge[1]
        graph[a] = append(graph[a], b)
        graph[b] = append(graph[b], a)
    }

    // Create a set for O(1) look-up time for restricted nodes
    restrictedSet := make(map[int]bool, n)
    for _, node := range restricted {
        restrictedSet[node] = true
    }

    // Array to keep track of visited nodes
    visited := make([]bool, n)

    // Start DFS from node 0
    return dfs(0, graph, visited, restrictedSet)
}

func dfs(node int, graph [][]int, visited []bool, restrictedSet map[int]bool) int {
    // Check if the node is restricted or already visited
    if restrictedSet[node] || visited[node] {
        return 0
    }

    // Mark this node as visited
    visited[node] = true

    // Start count from 1 (this node itself)
    count := 1

    // Recur for all adjacent nodes
    for _, neighbor := range graph[node] {
        count += dfs(neighbor, graph, visited, restrictedSet)
    }

    return count
}

func main() {
    n := 7
```

### Explanation
- **DFS Functionality**: The DFS checks for restrictions and whether a node is visited. If clear, it proceeds to recursively count all reachable nodes from that point, ensuring not to revisit nodes


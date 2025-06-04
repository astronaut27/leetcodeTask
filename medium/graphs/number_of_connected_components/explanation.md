# Intuition
When presented with the problem of finding connected components in a graph represented by nodes (`n`) and edges, it's akin to identifying isolated groups where each group represents a network of directly or indirectly connected nodes. This can be approached by using graph traversal methods like Depth-First Search (DFS) to explore each node and its connected nodes thoroughly before moving to another unvisited node, marking components as we go.

# Approach
1. **Graph Construction**:
    - Construct an adjacency list from the given `edges` where each node points to its directly connected neighbors. This allows easy traversal of connected nodes.

2. **Depth-First Search (DFS) to Mark Components**:
    - Use a boolean array `visited` to keep track of which nodes have been visited.
    - Iterate over each node. If it hasn't been visited, start a DFS from that node to mark all nodes reachable from it, indicating they all belong to the same component.
    - Each time a DFS is initiated, increment the count of connected components.

3. **DFS Implementation**:
    - Recursively visit each node's neighbors. If a neighbor has not been visited, recurse into it, marking all reachable nodes.

# Complexity
- **Time complexity**: O(V + E), where `V` is the number of vertices (nodes) and `E` is the number of edges. Each node and edge is visited exactly once in the worst case when constructing the adjacency list and during the DFS.
- **Space complexity**: O(V + E), to store the adjacency list and the visited array. The recursion stack during the DFS also contributes, but its depth will not exceed `V`.

# Code
```golang
package main

import "fmt"

func countComponents(n int, edges [][]int) int {
    graph := make([][]int, n)
    for _, edge := range edges {
        a, b := edge[0], edge[1]
        graph[a] = append(graph[a], b)
        graph[b] = append(graph[b], a)
    }

    visited := make([]bool, n)
    result := 0

    for i := 0; i < n; i++ {
        if !visited[i] {
            dfs(i, graph, visited)
            result++ // Each DFS call that starts a new traversal means a new component found
        }
    }

    return result
}

func dfs(node int, graph [][]int, visited []bool) {
    visited[node] = true
    for _, neighbor := range graph[node] {
        if !visited[neighbor] {
            dfs(neighbor, graph, visited)
        }
    }
}

func main() {
    edges := [][]int{{0, 1}, {1, 2}, {3, 4}}
    fmt.Println("Number of components:", countComponents(5, edges)) // Output: 2
}
```


### Explanation
- **Initialization**: The adjacency list is built by iterating through the given edges, ensuring each connection is bidirectional since it's an undirected graph.
- **Component Counting**: By initiating a DFS from each unvisited node, each unique DFS call corresponds to a new component because it explores all nodes connected to the starting node.
- **DFS Functionality**: The recursive nature of DFS ensures that all nodes that are part of the same connected component are marked visited during the same DFS call.

This solution is efficient and straightforward, effectively categorizing nodes into connected components using a classic graph traversal technique.

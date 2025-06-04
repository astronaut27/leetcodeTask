# Intuition
The problem asks for the maximum number of bombs that can be detonated by triggering a single bomb, where each bomb may cause other bombs to explode if they are within its blast radius. This naturally forms a **directed graph**, where each bomb is a node and an edge exists from bomb A to bomb B if B lies within the explosion radius of A.

# Approach
We build a graph where each bomb is a node, and directed edges represent the reachability (explosion) from one bomb to another. We use **depth-first search (DFS)** to simulate the chain reaction that starts from each bomb, and count how many bombs are detonated in total. By doing this for each bomb and taking the maximum result, we find the optimal start point for detonation.

Steps:
1. Build a graph: For each pair of bombs, add an edge from `i` to `j` if bomb `j` lies within the explosion radius of bomb `i`.
2. For every bomb, simulate the chain detonation using DFS.
3. Keep track of the maximum number of bombs that can be detonated from any single starting bomb.

# Complexity
- Time complexity:  
  $$O(n^2)$$ for building the graph (checking all pairs of bombs)  
  and $$O(n^2)$$ for running DFS from each node in the worst case  
  → Overall: $$O(n^2)$$

- Space complexity:  
  $$O(n^2)$$ for the adjacency list in the worst case  
  $$O(n)$$ for the visited map during each DFS

# Code
```golang
package main

import (
	"fmt"
)

func maximumDetonation(bombs [][]int) int {
	n := len(bombs)

	// Построение графа: b[i] -> b[j] если j в радиусе действия i
	graph := make(map[int][]int)
	for i := 0; i < n; i++ {
		x1, y1, r := bombs[i][0], bombs[i][1], bombs[i][2]
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}
			x2, y2 := bombs[j][0], bombs[j][1]
			dx, dy := x2-x1, y2-y1
			if dx*dx+dy*dy <= r*r {
				graph[i] = append(graph[i], j)
			}
		}
	}

	maxCount := 0
	for i := 0; i < n; i++ {
		visited := make(map[int]bool)
		count := dfs(i, graph, visited)
		if count > maxCount {
			maxCount = count
		}
	}
	return maxCount
}

// 📦 Отдельная функция DFS
func dfs(node int, graph map[int][]int, visited map[int]bool) int {
	if visited[node] {
		return 0
	}
	visited[node] = true
	count := 1
	for _, nei := range graph[node] {
		count += dfs(nei, graph, visited)
	}
	return count
}

func main() {
	fmt.Println(maximumDetonation([][]int{{2, 1, 3}, {6, 1, 4}})) // 2
	fmt.Println(maximumDetonation([][]int{{1, 1, 5}, {10, 10, 5}})) // 1
	fmt.Println(maximumDetonation([][]int{{1, 2, 3}, {2, 3, 1}, {3, 4, 2}, {4, 5, 3}, {5, 6, 4}})) // 5
}

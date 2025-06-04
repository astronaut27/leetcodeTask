# Intuition
The problem is essentially a shortest path search through a state space of 4-digit lock combinations. Each digit can be rotated forward or backward (wraps around 0–9), and certain states are "deadends" that cannot be visited. The goal is to find the fewest number of moves from `"0000"` to a given `target` while avoiding the deadends. BFS is a natural choice here as it finds the shortest path in an unweighted graph-like structure.

# Approach
1. **Deadend and Visited Sets**:
    - Store the deadends in a hash map for O(1) lookup.
    - Use a `visited` set to prevent cycles and redundant processing.

2. **Breadth-First Search (BFS)**:
    - Initialize BFS from `"0000"`.
    - At each step, generate all possible next combinations by rotating each of the 4 digits one step forward and one step backward.
    - Add valid, non-deadend and unvisited combinations to the queue.

3. **Termination Conditions**:
    - If `"0000"` is the target, return `0` immediately.
    - If the queue is exhausted without finding the target, return `-1`.

# Complexity
- **Time complexity**: $$O(10^4)$$  
  Since there are only 10,000 possible 4-digit combinations, and each generates up to 8 neighbors, the BFS is bounded by a constant number of states.

- **Space complexity**: $$O(10^4)$$  
  For storing the visited map and the BFS queue.

# Code
```golang
package main

import (
	"fmt"
)

func openLock(deadends []string, target string) int {
	// Set of deadend states
	dead := make(map[string]bool)
	for _, d := range deadends {
		dead[d] = true
	}

	if dead["0000"] {
		return -1
	}
	if target == "0000" {
		return 0
	}

	// BFS initialization
	visited := make(map[string]bool)
	queue := []string{"0000"}
	visited["0000"] = true
	steps := 0

	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			curr := queue[0]
			queue = queue[1:]

			if curr == target {
				return steps
			}

			// Generate next states
			for _, next := range getNextStates(curr) {
				if !visited[next] && !dead[next] {
					visited[next] = true
					queue = append(queue, next)
				}
			}
		}
		steps++
	}

	return -1
}

// Helper to generate all valid next combinations
func getNextStates(state string) []string {
	result := []string{}
	bytes := []byte(state)

	for i := 0; i < 4; i++ {
		original := bytes[i]

		// Rotate forward
		bytes[i] = (original-'0'+1)%10 + '0'
		result = append(result, string(bytes))

		// Rotate backward
		bytes[i] = (original-'0'+9)%10 + '0'
		result = append(result, string(bytes))

		// Restore original digit
		bytes[i] = original
	}

	return result
}

func main() {
	fmt.Println(openLock([]string{"0201", "0101", "0102", "1212", "2002"}, "0202"))                         // 6
	fmt.Println(openLock([]string{"8888"}, "0009"))                                                         // 1
	fmt.Println(openLock([]string{"8887", "8889", "8878", "8898", "8788", "8988", "7888", "9888"}, "8888")) // -1
}
```

### Summary
- This problem is a classic BFS use case for shortest-path search in an implicit graph.
- Edge cases are handled early: if `"0000"` is the target or in the deadend list.
- The use of string manipulation and byte operations ensures that state generation is efficient and correct.

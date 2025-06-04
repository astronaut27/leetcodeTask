# Intuition
Given a list of directed roads connecting various cities, the goal is to determine the minimum number of changes needed to ensure that all roads lead directly or indirectly to the capital city (city 0). This problem can be visualized as a directed graph where nodes represent cities and edges represent directed roads. The challenge is to find out how many edges (roads) need their direction reversed to ensure all paths can lead to city 0.

# Approach
1. **Graph Representation**:
    - Represent the city connections using an adjacency list where each city points to its direct connections. For easy identification of direction, store outgoing roads as positive and incoming roads as negative values.

2. **Breadth-First Search (BFS)**:
    - Perform a BFS starting from city 0. The BFS will traverse the graph and keep track of cities visited.
    - While traversing, count the number of outgoing roads (positive connections) from each visited city that lead to unvisited cities. These roads count towards the changes needed since they represent roads that need to be reversed to direct towards city 0.

3. **Direction Reversal Counting**:
    - For each visited node, increment the count of changes when an outgoing road leads to an unvisited city. This operation effectively reverses the road's direction.

# Complexity
- **Time complexity**: O(n + m), where `n` is the number of cities and `m` is the number of connections. Each city and connection is processed exactly once.
- **Space complexity**: O(n + m), which accounts for the storage of the graph and the BFS queue. The graph stores both incoming and outgoing connections, and the queue might store all cities in the worst-case scenario of traversing the entire graph.


### Explanation
- **Graph Setup**: The graph uses an adjacency list where outgoing roads are stored as positive numbers and incoming as negative. This distinction helps determine which roads contribute to the count of changes.
- **BFS Operation**: BFS is used for its level-order traversal, ensuring all nodes reachable from city 0 are processed. It effectively counts how many roads must be reversed to ensure all can lead to city 0.
- **Change Counting**: Only outgoing roads leading to unvisited cities increase the change count, aligning with the goal of redirecting all roads towards the capital.

This approach ensures a systematic traversal of the graph and an accurate count of the minimum reversals needed, leveraging the BFS's ability to exhaustively explore each city's connections.

# Code
```golang
package main

import "fmt"

func minReorder(n int, connections [][]int) int {
    // Создаем список смежности
    graph := make([][]int, n)
    for _, conn := range connections {
        a, b := conn[0], conn[1]
        graph[a] = append(graph[a], b)  // Оригинальное направление
        graph[b] = append(graph[b], -a) // Обратное направление
    }

    // Используем BFS для обхода графа
    visited := make([]bool, n)
    queue := []int{0}
    visited[0] = true
    changes := 0

    for len(queue) > 0 {
        node := queue[0]
        queue = queue[1:]

        for _, neighbor := range graph[node] {
            absNeighbor := abs(neighbor)
            if !visited[absNeighbor] {
                visited[absNeighbor] = true
                queue = append(queue, absNeighbor)
                if neighbor > 0 {  // Проверяем, является ли дорога исходящей
                    changes++
                }
            }
        }
    }

    return changes
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}

func main() {
    connections := [][]int{{0, 1}, {1, 3}, {2, 3}, {4, 0}, {4, 5}}
    n := 6
    fmt.Println("Minimum reorder required:", minReorder(n, connections))
}
```
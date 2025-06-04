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
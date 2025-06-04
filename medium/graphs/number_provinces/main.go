package main

import "fmt"

func findCircleNum(isConnected [][]int) int {
	n := len(isConnected)
	visited := make([]bool, n)
	var provinces int
	for i := 0; i < n; i++ {
		if !visited[i] {
			provinces++
			stack := []int{i}
			for len(stack) > 0 {
				node := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				visited[node] = true // Mark the node as visited when popped from the stack
				for neighbor, connect := range isConnected[node] {
					if connect == 1 && !visited[neighbor] {
						stack = append(stack, neighbor)
					}
				}
			}
		}
	}
	return provinces
}

func main() {
	isConnected := [][]int{
		{1, 1, 0},
		{1, 1, 0},
		{0, 0, 1},
	}
	fmt.Println("Number of provinces:", findCircleNum(isConnected))
}

package main

import "fmt"

func canVisitAllRooms(rooms [][]int) bool {
	visit := map[int]bool{0: true} // Start with room 0 as visited
	stack := []int{0}              // Stack for DFS, start with room 0

	for len(stack) > 0 {
		node := stack[len(stack)-1]            // Current node/room
		stack = stack[:len(stack)-1]           // Pop from stack
		for _, neighbor := range rooms[node] { // Explore each key/neighbor
			if !visit[neighbor] { // If this room hasn't been visited
				visit[neighbor] = true          // Mark it as visited
				stack = append(stack, neighbor) // Add it to the stack for further exploration
			}
		}
	}
	return len(visit) == len(rooms) // Check if all rooms were visited
}

func main() {
	rooms := [][]int{{1}, {2}, {3}, {}}
	fmt.Println("Can visit all rooms:", canVisitAllRooms(rooms))
}

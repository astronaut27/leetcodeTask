
func reachableNodes(n int, edges [][]int, restricted []int) int {
    graph := make([][]int, n)
    for _, edge := range edges {
        a, b := edge[0], edge[1]
        graph[a] = append(graph[a], b)
        graph[b] = append(graph[b], a)
    }
    restrictedSet:=make(map[int]bool,n)
    for _,node:=range restricted{
        restrictedSet[node]=true
    }
    visited := make([]bool, n)
    result := 0
    return dfs(i, graph, visited,restrictedSet)
}

func dfs(node int, graph [][]int, visited []bool,restrictedSet map[int]bool)int{
if restrictedSet[node] || visited[node]{
        return 0
    }
    visited[node] = true
	count := 1
    	for _, neighbor := range graph[node] {
		count += dfs(neighbor, graph, visited, restrictedSet)
	}

	return count
}
package main

func snakesAndLadders(board [][]int) int {
	n := len(board)

	// Преобразуем доску в одномерный массив с учетом змейки
	flatten := make([]int, n*n+1)
	idx := 1
	leftToRight := true
	for i := n - 1; i >= 0; i-- {
		if leftToRight {
			for j := 0; j < n; j++ {
				flatten[idx] = board[i][j]
				idx++
			}
		} else {
			for j := n - 1; j >= 0; j-- {
				flatten[idx] = board[i][j]
				idx++
			}
		}
		leftToRight = !leftToRight
	}

	visited := make([]bool, n*n+1)
	queue := []int{1}
	visited[1] = true
	moves := 0

	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			curr := queue[0]
			queue = queue[1:]

			if curr == n*n {
				return moves
			}

			for dice := 1; dice <= 6; dice++ {
				next := curr + dice
				if next > n*n {
					continue
				}
				if flatten[next] != -1 {
					next = flatten[next]
				}
				if !visited[next] {
					visited[next] = true
					queue = append(queue, next)
				}
			}
		}
		moves++
	}
	return -1
}

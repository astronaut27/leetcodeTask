package main

import "fmt"

func generateParenthesis(n int) []string {
	var ans []string
	backtrack(n, n, "", &ans)
	return ans
}

func backtrack(open, close int, temp string, ans *[]string) {
	// Базовый случай: если длина текущей комбинации равна 2*n
	if open == 0 && close == 0 {
		*ans = append(*ans, temp)
		return
	}

	// Если можем добавить открывающую скобку
	if open > 0 {
		backtrack(open-1, close, temp+"(", ans)
	}

	// Если можем добавить закрывающую скобку
	if close > open {
		backtrack(open, close-1, temp+")", ans)
	}
}

func main() {
	fmt.Println(generateParenthesis(3)) // ["((()))","(()())","(())()","()(())","()()()"]
}

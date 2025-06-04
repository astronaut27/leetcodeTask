package main

import (
	"fmt"
)

func ladderLength(beginWord string, endWord string, wordList []string) int {
	// Шаг 1: поместим все слова из wordList в множество
	wordSet := make(map[string]bool)
	for _, word := range wordList {
		wordSet[word] = true
	}

	// Если конечного слова нет в словаре — сразу возвращаем 0
	if !wordSet[endWord] {
		return 0
	}

	// Шаг 2: запускаем BFS от beginWord
	queue := []string{beginWord}
	visited := make(map[string]bool)
	visited[beginWord] = true
	steps := 1

	for len(queue) > 0 {
		levelSize := len(queue)
		for i := 0; i < levelSize; i++ {
			word := queue[0]
			queue = queue[1:]

			// Генерируем все возможные слова с заменой 1 буквы
			wordBytes := []byte(word)
			for j := 0; j < len(wordBytes); j++ {
				original := wordBytes[j]
				for c := byte('a'); c <= 'z'; c++ {
					if c == original {
						continue
					}
					wordBytes[j] = c
					newWord := string(wordBytes)

					// Если дошли до цели
					if newWord == endWord {
						return steps + 1
					}

					if wordSet[newWord] && !visited[newWord] {
						visited[newWord] = true
						queue = append(queue, newWord)
					}
				}
				wordBytes[j] = original // восстанавливаем букву
			}
		}
		steps++
	}

	return 0
}

func main() {
	fmt.Println(ladderLength("hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"})) // 5
	fmt.Println(ladderLength("hit", "cog", []string{"hot", "dot", "dog", "lot", "log"}))        // 0
}

# Intuition
This problem is a classic **shortest path** problem in disguise. Each word can be thought of as a node in a graph, and there’s an edge between two nodes if their words differ by exactly one character. We’re tasked with finding the shortest path from `beginWord` to `endWord` following such edges.

# Approach
We perform **Breadth-First Search (BFS)** starting from `beginWord`. At each step, we generate all possible one-letter mutations of the current word. If a mutation exists in the word list and hasn't been visited yet, we enqueue it. BFS ensures we find the shortest transformation sequence first.

The algorithm uses:
- A `map` to store all words for fast lookup.
- A `visited` map to prevent revisiting words.
- A queue for BFS traversal.
- For each word, we attempt to change each character to every letter from `'a'` to `'z'`.

We terminate early if we find the `endWord`.

# Complexity
- Time complexity:  
  $$O(N \cdot M^2)$$  
  where \(N\) is the number of words in the list and \(M\) is the word length. We try 26 letters at each position in every word.

- Space complexity:  
  $$O(N \cdot M)$$  
  due to storing the word list, visited map, and BFS queue.

# Code
```golang
package main

import (
	"fmt"
)

func ladderLength(beginWord string, endWord string, wordList []string) int {
	wordSet := make(map[string]bool)
	for _, word := range wordList {
		wordSet[word] = true
	}

	if !wordSet[endWord] {
		return 0
	}

	queue := []string{beginWord}
	visited := make(map[string]bool)
	visited[beginWord] = true
	steps := 1

	for len(queue) > 0 {
		levelSize := len(queue)
		for i := 0; i < levelSize; i++ {
			word := queue[0]
			queue = queue[1:]

			wordBytes := []byte(word)
			for j := 0; j < len(wordBytes); j++ {
				original := wordBytes[j]
				for c := byte('a'); c <= 'z'; c++ {
					if c == original {
						continue
					}
					wordBytes[j] = c
					newWord := string(wordBytes)

					if newWord == endWord {
						return steps + 1
					}

					if wordSet[newWord] && !visited[newWord] {
						visited[newWord] = true
						queue = append(queue, newWord)
					}
				}
				wordBytes[j] = original
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

# Intuition
The problem is similar to finding the shortest path in a graph, where each valid gene string in the `bank` can be a node, and an edge exists between two nodes if they differ by exactly one character. The goal is to transform the `start` gene into the `end` gene with the fewest mutations, using only valid mutations from the bank. This immediately suggests a **Breadth-First Search (BFS)** approach.

# Approach
We use BFS to explore all valid mutations level by level. Each mutation step generates all possible one-character changes of the current gene. If the resulting gene exists in the `bank` and hasn’t been visited before, we add it to the BFS queue. We continue until we reach the `end` gene or exhaust all options. A set is used for constant-time lookup of valid genes, and a visited map ensures we don’t revisit the same gene multiple times.

Steps:
1. Add all genes from `bank` into a set for quick lookup.
2. Start BFS from the `start` gene.
3. At each step, for each gene in the current level, try all possible mutations.
4. If a mutation is valid and hasn't been visited, add it to the queue.
5. Stop when the `end` gene is found.

# Complexity
- Time complexity:  
  $$O(n \cdot m \cdot 4)$$, where  
  - \( n \) is the number of genes in the bank,  
  - \( m \) is the length of each gene (always 8 in this problem),  
  - we try 4 possible replacements (A, C, G, T) at each position.

- Space complexity:  
  $$O(n)$$ for storing the gene bank and visited genes.

# Code
```golang []
func minMutation(start string, end string, bank []string) int {
    bankSet := make(map[string]bool)
    for _, gene := range bank {
        bankSet[gene] = true
    }
    if !bankSet[end] {
        return -1 // Если конечного гена нет в банке — смысла нет
    }

    queue := []string{start}
    visited := map[string]bool{start: true}
    steps := 0
    genes := []byte{'A', 'C', 'G', 'T'}

    for len(queue) > 0 {
        for size := len(queue); size > 0; size-- {
            curr := queue[0]
            queue = queue[1:]

            if curr == end {
                return steps
            }

            // Пробуем менять по одной букве
            currBytes := []byte(curr)
            for i := 0; i < len(currBytes); i++ {
                original := currBytes[i]
                for _, g := range genes {
                    currBytes[i] = g
                    next := string(currBytes)
                    if bankSet[next] && !visited[next] {
                        visited[next] = true
                        queue = append(queue, next)
                    }
                }
                currBytes[i] = original // откат
            }
        }
        steps++
    }

    return -1
}

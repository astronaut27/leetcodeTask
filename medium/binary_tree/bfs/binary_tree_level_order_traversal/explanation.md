# Intuition
To traverse a binary tree level by level, a Breadth-First Search (BFS) approach naturally fits the requirement. Using a queue allows us to process each level's nodes in the order they appear, and collect their values accordingly.

# Approach
We use a queue to perform BFS starting from the root. For each level, we:
1. Determine how many nodes are in the current level (`nodesInCurrentLevel`).
2. Iterate through these nodes, dequeueing them and collecting their values into a slice.
3. Enqueue their left and right children if they exist.
4. After finishing a level, append the collected values to the final result.

The outer loop continues until the queue is empty, which means all tree levels have been processed.

# Complexity
- Time complexity:  
  $$O(n)$$  
  where \( n \) is the number of nodes in the tree, since each node is visited exactly once.

- Space complexity:  
  $$O(n)$$  
  for the queue and the result, in the worst case when the tree is completely balanced and the last level has \( n/2 \) nodes.

# Code
```golang
func levelOrder(root *TreeNode) [][]int {
    var result [][]int

    if root == nil {
        return result
    }

    queue := []*TreeNode{root}

    for len(queue) > 0 {
        nodesInCurrentLevel := len(queue)
        level := make([]int, 0, nodesInCurrentLevel)

        for i := 0; i < nodesInCurrentLevel; i++ {
            node := queue[0]
            queue = queue[1:]
            level = append(level, node.Val)

            if node.Left != nil {
                queue = append(queue, node.Left)
            }
            if node.Right != nil {
                queue = append(queue, node.Right)
            }
        }

        result = append(result, level)
    }

    return result
}
```
# Intuition
The goal is to find all nodes in a binary tree that are exactly `k` edges away from a given target node. At first glance, we might consider performing a BFS or DFS from the target node. However, in a typical binary tree, nodes do not have references to their parent nodes, making it hard to traverse "upward." To overcome this, we can perform a two-step DFS approach: the first to locate the target and track distance from the root, and the second to explore all nodes at the required distance.

# Approach
1. **DFS for Locating Target and Tracking Distances**:
    - Start with a DFS from the root to locate the target node.
    - As we recurse back up, we track the distance from each ancestor node to the target.
    - At each node along the path, we check if the node itself or its subtree contains nodes at distance `k`.

2. **DFS for Collecting Nodes at Distance K**:
    - Once the target is located, we perform another DFS to collect all nodes at distance `k` from it.
    - If a node is an ancestor of the target, we must consider its opposite subtree (not the one containing the target) to find additional nodes at the remaining distance.

3. **Recursive Returns and Distance Propagation**:
    - The first DFS returns the distance from the current node to the target if the target exists in its subtree.
    - This distance is used to evaluate whether the current node or its other subtree contains nodes at the required distance.

# Complexity
- **Time complexity**: $$O(n)$$  
  Every node is visited at most once in both the initial DFS (to find the target) and the secondary DFS (to collect nodes at distance `k`).
- **Space complexity**: $$O(h)$$  
  The space is used for the recursion stack, where `h` is the height of the tree. In the worst case, the height is `n` for a skewed tree.

# Code
```golang
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// Helper function to perform DFS traversal and find nodes at distance K
func dfs(root *TreeNode, target *TreeNode, k int, parent *TreeNode, result *[]int) int {
    if root == nil {
        return -1
    }

    if root == target {
        // If the current node is the target, start the second DFS to find nodes at distance K
        findNodesAtDistanceK(root, k, result)
        return 0
    }

    // Recursively search the left and right children
    leftDist := dfs(root.Left, target, k, root, result)
    rightDist := dfs(root.Right, target, k, root, result)

    // If the target is found on the left subtree
    if leftDist != -1 {
        // If the distance from the target node to the current node is exactly K
        if leftDist+1 == k {
            *result = append(*result, root.Val)
        }
        // Search the right subtree (opposite direction) for remaining distance
        findNodesAtDistanceK(root.Right, k-leftDist-2, result)
        return leftDist + 1
    }

    // If the target is found on the right subtree
    if rightDist != -1 {
        // If the distance from the target node to the current node is exactly K
        if rightDist+1 == k {
            *result = append(*result, root.Val)
        }
        // Search the left subtree (opposite direction) for remaining distance
        findNodesAtDistanceK(root.Left, k-rightDist-2, result)
        return rightDist + 1
    }

    return -1
}

// Helper function to find nodes at distance K from the current node
func findNodesAtDistanceK(root *TreeNode, k int, result *[]int) {
    if root == nil || k < 0 {
        return
    }

    if k == 0 {
        *result = append(*result, root.Val)
        return
    }

    findNodesAtDistanceK(root.Left, k-1, result)
    findNodesAtDistanceK(root.Right, k-1, result)
}

// Main function to get nodes at distance K from the target
func distanceK(root *TreeNode, target *TreeNode, k int) []int {
    var result []int
    dfs(root, target, k, nil, &result)
    return result
}
```

### Summary
- This solution uses a clever recursive DFS approach to handle both downward and upward traversal from the target node without explicitly storing parent pointers.
- It efficiently collects nodes at distance `k` from the target by navigating both toward and away from the target.
- A two-pass strategy ensures accurate distance tracking and optimal traversal.

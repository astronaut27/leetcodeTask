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
        // Search the right subtree (parent node in the DFS)
        findNodesAtDistanceK(root.Right, k-leftDist-2, result)
        return leftDist + 1
    }

    // If the target is found on the right subtree
    if rightDist != -1 {
        // If the distance from the target node to the current node is exactly K
        if rightDist+1 == k {
            *result = append(*result, root.Val)
        }
        // Search the left subtree (parent node in the DFS)
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

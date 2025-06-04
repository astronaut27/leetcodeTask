package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Res struct {
	Res []int
}

func inorderTraversal(root *TreeNode) []int {
	r := Res{}
	r.dfs(root)
	return r.Res
}

func (r *Res) dfs(node *TreeNode) {
	if node != nil {
		r.dfs(node.Left)
		r.Res = append(r.Res, node.Val)
		r.dfs(node.Right)
	}
}

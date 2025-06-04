package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func New(rootValue int) *TreeNode {
	return &TreeNode{
		Val: rootValue,
	}
}

func (b *TreeNode) Add(val int) {
	add(b, val)
}

func add(leaf *TreeNode, val int) *TreeNode {
	if leaf == nil {
		return &TreeNode{
			Left:  nil,
			Right: nil,
			Val:   val,
		}
	}
	switch {
	case leaf.Val > val:
		leaf.Left = add(leaf.Left, val)
	case leaf.Val < val:
		leaf.Right = add(leaf.Right, val)
	case leaf.Val == val:
		return leaf
	}
	return leaf
}

func main() {
	myTree := New(6)
	myTree.Add(2)
	myTree.Add(8)
	myTree.Add(0)
	myTree.Add(4)
	myTree.Add(7)
	myTree.Add(9)
	myTree.Add(3)
	myTree.Add(5)
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	if p.Val < root.Val && q.Val < root.Val {
		return lowestCommonAncestor(root.Left, p, q)
	}
	if p.Val > root.Val && q.Val > root.Val {
		return lowestCommonAncestor(root.Right, p, q)
	}
	if p.Val <= root.Val && q.Val >= root.Val {
		return root
	}
	return root
}

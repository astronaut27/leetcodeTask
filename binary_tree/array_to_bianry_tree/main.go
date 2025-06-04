package main

import "fmt"

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Val   int
}

func (b *TreeNode) Print() {
	print(b)
}

func print(leaf *TreeNode) {
	fmt.Println(leaf.Val)
	if leaf.Right != nil {
		print(leaf.Right)
	}
	if leaf.Left != nil {
		print(leaf.Left)
	}
	return
}

func main() {
	mytree := sortedArrayToBST([]int{-10, -3, 0, 5, 9})
	mytree.Print()
}

func sortedArrayToBST(nums []int) *TreeNode {
	return ToBST(nums, 0, len(nums)-1)
}

func ToBST(nums []int, start, finish int) *TreeNode {
	if start > finish {
		return nil
	}
	mid := start + (finish-start)/2
	root := &TreeNode{Val: nums[mid]}

	root.Left = ToBST(nums, start, mid-1)
	root.Right = ToBST(nums, mid+1, finish)

	return root
}

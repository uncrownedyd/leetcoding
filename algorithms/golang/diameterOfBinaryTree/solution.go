package diameterOfBinaryTree

/**
* Definition for a binary tree node.
* type TreeNode struct {
*     Val int
*     Left *TreeNode
*     Right *TreeNode
* }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res int

func diameterOfBinaryTree(root *TreeNode) int {
	res = 0
	helper(root)
	return res
}

func helper(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l := helper(root.Left)
	r := helper(root.Right)
	if l + r > res {
		res = l + r
	}
	return max(r, l) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
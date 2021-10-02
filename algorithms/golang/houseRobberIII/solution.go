package houseRobberIII

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rob(root *TreeNode) int {
	l1, l2 := helper(root.Left)
	r1, r2 := helper(root.Right)
	return max(l2 + r2 + root.Val, max(l1, l2) + max(r1, r2))
}

// n1: 以cur为终点的最大金额数
// n2: 不以cur为终点的最大金额数
func helper(cur *TreeNode) (n1, n2 int) {
	if cur == nil {
		return 0, 0
	}

	l1, l2 := helper(cur.Left)
	r1, r2 := helper(cur.Right)
	return l2 + r2 + cur.Val, max(l1, l2) + max(r1, r2)
	//return max(l2, r2) + cur.Val, max(l1, r1)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
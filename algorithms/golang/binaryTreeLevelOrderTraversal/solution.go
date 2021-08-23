package binaryTreeLevelOrderTraversal

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	queue := make([]*TreeNode, 0)
	// Push to the queue
	queue = append(queue, &TreeNode{Val: 1})
	queue = append(queue, &TreeNode{Val: 2})
	// Top (just get next element, don't remove it)

	for len(queue) > 0 {
		l := len(queue)
		for i := 0; i < l; i++ {
			node := queue[i]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			fmt.Printf(" %d ", node.Val)
		}
		fmt.Print("\n")
		queue = queue[l:]
	}
	return [][]int{}
}

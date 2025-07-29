/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func isBalanced(root *TreeNode) bool {
	var (
		dfs     func(node *TreeNode, depth int) int
		MAX_INT int = 5005
	)

	dfs = func(node *TreeNode, depth int) int {
		if node == nil {
			return 0
		}
		left := dfs(node.Left, depth+1)
		right := dfs(node.Right, depth+1)

		if abs(left-right) > 1 {
			return MAX_INT
		}
		return max(left, right) + 1
	}
	res := dfs(root, 0)

	return res < MAX_INT
}
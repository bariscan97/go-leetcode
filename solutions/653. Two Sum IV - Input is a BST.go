/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findTarget(root *TreeNode, k int) bool {
	var dfs func(node *TreeNode) bool
	dic := make(map[int]bool)

	dfs = func(node *TreeNode) bool {
		if node == nil {
			return false
		}
		if dic[k-node.Val] {
			return true
		}
		dic[node.Val] = true
		left := dfs(node.Left)
		right := dfs(node.Right)
		return left || right
	}

	return dfs(root)
}
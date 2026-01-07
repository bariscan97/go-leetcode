/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

const MOD int64 = 1_000_000_007


func maxProduct(root *TreeNode) int {
	var sumDfs func(*TreeNode) int64
	sumDfs = func(node *TreeNode) int64 {
		if node == nil {
			return 0
		}
		return sumDfs(node.Left) + sumDfs(node.Right) + int64(node.Val)
	}
	total := sumDfs(root)

	
	var best int64 = 0
	var prodDfs func(*TreeNode) int64
	prodDfs = func(node *TreeNode) int64 {
		if node == nil {
			return 0
		}
		left := prodDfs(node.Left)
		right := prodDfs(node.Right)
		sub := left + right + int64(node.Val)

		
		product := sub * (total - sub)
		if product > best {
			best = product
		}
		return sub
	}
	prodDfs(root)

	return int(best % MOD)
}

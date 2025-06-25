/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func inorderTraversal(root *TreeNode) []int {
    var (
        arr []int
        dfs func(node *TreeNode)
    )
    dfs = func(node *TreeNode) {
        if node == nil {
            return 
        }
        dfs(node.Left)
        arr = append(arr, node.Val)
        dfs(node.Right)
    }
    dfs(root)
    return arr
}
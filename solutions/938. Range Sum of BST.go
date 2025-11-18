/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rangeSumBST(root *TreeNode, low int, high int) int {
    q := []*TreeNode{root}
    var res int
    for len(q) > 0 {
        size := len(q)
        for i := 0; i < size; i++ {
            node := q[0]
            if low <= node.Val && node.Val <= high {
                res += node.Val
            }
            if node.Left != nil {
                q = append(q, node.Left)
            }
            if node.Right != nil {
                q = append(q, node.Right)
            }
            q = q[1:]
        }
    }
    return res
}
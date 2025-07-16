/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func deepestLeavesSum(root *TreeNode) int {
    q := []*TreeNode{root}
    var res, curr int
    for len(q) > 0 {
        size := len(q)
        for size > 0 {
            node := q[0]
            q = q[1:]
            if node.Left != nil { q = append(q, node.Left) }
            if node.Right != nil { q = append(q, node.Right) }
            curr += node.Val
            size--
        }
        res = curr 
        curr = 0
    }
    return res
}
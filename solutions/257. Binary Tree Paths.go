/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func binaryTreePaths(root *TreeNode) []string {
   var (
        res []string
        dfs func(node *TreeNode, str string)
    )

    dfs = func(node *TreeNode, str string) {
        if node == nil {
            return
        }
        strNum := strconv.Itoa(node.Val)
        
        if node.Left == nil && node.Right == nil {
            res = append(res, str + strNum)
        }

        dfs(node.Left, str  + strNum + "->")
        dfs(node.Right, str  + strNum + "->")
    }

    dfs(root, "")

    return res
}
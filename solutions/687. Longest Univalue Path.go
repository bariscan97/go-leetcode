/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func longestUnivaluePath(root *TreeNode) int {
    res := 0
    var dfs func(node *TreeNode) int
    dfs = func(node *TreeNode) int {
        if node == nil {
            return 0
        }

        left := dfs(node.Left)
        right := dfs(node.Right)

        if node.Left != nil && node.Right != nil {
            if node.Left.Val == node.Val && node.Right.Val == node.Val {
                if left+right+2 > res {
                    res = left + right + 2
                }
                if left > right {
                    return left + 1
                }
                return right + 1
            }
            if node.Left.Val == node.Val {
                if left+1 > res {
                    res = left + 1
                }
                return left + 1
            }
            if node.Right.Val == node.Val {
                if right+1 > res {
                    res = right + 1
                }
                return right + 1
            }
            return 0
        } else if node.Left != nil {
            if node.Left.Val == node.Val {
                if left+1 > res {
                    res = left + 1
                }
                return left + 1
            }
            return 0
        } else if node.Right != nil {
            if node.Right.Val == node.Val {
                if right+1 > res {
                    res = right + 1
                }
                return right + 1
            }
            return 0
        } else {
            return 0
        }
    }

    dfs(root)
    return res
}

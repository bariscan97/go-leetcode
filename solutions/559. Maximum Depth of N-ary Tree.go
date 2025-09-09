/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func maxDepth(root *Node) int {
    if root == nil {
        return 0
    }
    q := []*Node{}
    q = append(q, root)
    var res int
    for len(q) > 0 {
        size := len(q)
        for i := 0; i < size; i++ {
            node := q[0]
            q = q[1:]
            for _, j := range node.Children {
                q = append(q,j)
            }
        }
        res++
    }
    return res
}
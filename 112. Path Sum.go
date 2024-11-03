/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 type Queue struct {
	items *list.List
}

func NewQueue() *Queue {
	return &Queue{items: list.New()}
}

func (q *Queue) Enqueue(value *Tuple) {
	q.items.PushBack(value)
}

func (q *Queue) IsEmpty() bool {
	return q.items.Len() == 0
}

func (q *Queue) Dequeue() (*TreeNode, int) {
	if q.items.Len() == 0 {
		return nil ,0
	}
	front := q.items.Front()
	q.items.Remove(front)
    elem, _ := front.Value.(*Tuple)
	return elem.Node , elem.Total
}

type Tuple struct {
    Node *TreeNode
    Total int
}

func hasPathSum(root *TreeNode, targetSum int) bool {       
    q := NewQueue()
    q.Enqueue(&Tuple{
        Node : root,
        Total : 0,
    })
    for !q.IsEmpty() {
        node, total := q.Dequeue()
        if node == nil {
            continue
        }
        if node.Left == nil && node.Right == nil && node.Val + total == targetSum{
            return true
        }
        q.Enqueue(&Tuple{
            Node : node.Left,
            Total : total + node.Val,
        })
        q.Enqueue(&Tuple{
            Node : node.Right,
            Total : total + node.Val,
        })
    }

    return false
  
}


/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type IntHeap []int

func (h IntHeap) Len() int {
    return len(h)
}

func (h IntHeap) Less(i, j int) bool {
    return h[i] < h[j]
}

func (h IntHeap) Swap(i, j int) {
    h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x interface{}) {
    *h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}

func (h IntHeap) Empty() bool {
    return h.Len() == 0
}

func mergeKLists(lists []*ListNode) *ListNode {
    if len(lists) == 0 { return nil }
    
    h := &IntHeap{}
    heap.Init(h)
    
    for _, list := range lists {
        for list != nil {
            val := list.Val
            list = list.Next
            heap.Push(h, val)
        }
    }

    if h.Len() == 0 { return nil }

    res := &ListNode{}
    var node *ListNode = res
    
    for !h.Empty() {
        val := heap.Pop(h).(int)
        node.Val = val
        if !h.Empty() {
            node.Next = &ListNode{}
            node = node.Next 
        }
            
    }
    return res
}
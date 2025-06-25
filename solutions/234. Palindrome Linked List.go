/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
    var stack []int
    for head != nil {
        stack = append(stack, head.Val)
        head = head.Next
    }
    var N int = len(stack)
    for i := 0 ; i < N / 2; i++ {
        if stack[i] != stack[N - 1 - i] {
            return false
        }
    }
    return true
}
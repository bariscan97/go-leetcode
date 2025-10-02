/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type Solution struct {
    Nums []int
}

func Constructor(head *ListNode) Solution {
    var nums []int
    for head != nil {
        nums = append(nums, head.Val)
        head = head.Next
    }
    return Solution{
        Nums: nums,
    }
}


func (this *Solution) GetRandom() int {
    return this.Nums[rand.Intn(len(this.Nums))]
}


/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(head);
 * param_1 := obj.GetRandom();
 */
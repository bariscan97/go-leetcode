type MyQueue struct {
	arr []int
}

func Constructor() MyQueue {
	my_arr := make([]int, 0)
	return MyQueue{
		arr: my_arr,
	}
}

func (this *MyQueue) Push(x int) {
	this.arr = append([]int{x}, this.arr...)
}

func (this *MyQueue) Pop() int {
	last_elem := this.arr[len(this.arr)-1]
	this.arr = this.arr[:len(this.arr)-1]
	return last_elem
}

func (this *MyQueue) Peek() int {
	return this.arr[len(this.arr)-1]
}

func (this *MyQueue) Empty() bool {
	return len(this.arr) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
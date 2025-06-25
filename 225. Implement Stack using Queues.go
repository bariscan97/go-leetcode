type MyStack struct {
    Stack []int
}


func Constructor() MyStack {
    return MyStack{
        Stack : []int{},
    }
}


func (this *MyStack) Push(x int)  {
    this.Stack = append(this.Stack, x)
}


func (this *MyStack) Pop() int {
    val := this.Stack[len(this.Stack) - 1]
    this.Stack = this.Stack[:len(this.Stack) - 1] 
    return val
}


func (this *MyStack) Top() int {
    return this.Stack[len(this.Stack) -1]
}   


func (this *MyStack) Empty() bool {
    return len(this.Stack) == 0
}


/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
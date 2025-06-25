type MyCircularDeque struct {
	q     []int
	front int
	rear  int
	size  int // == len(q)
}

func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{
		q:     make([]int, k),
		front: -1,
		rear:  -1,
		size:  0,
	}
}

func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.IsFull() {
		return false
	}

	if this.front == -1 {
		this.front, this.rear = 0, 0
	} else {
		this.front = (this.front - 1 + cap(this.q)) % cap(this.q)
	}

	this.q[this.front] = value
	this.size++

	return true
}

func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.IsFull() {
		return false
	}

	if this.rear == -1 {
		this.front, this.rear = 0, 0
	} else {
		this.rear = (this.rear + 1) % cap(this.q)
	}

	this.q[this.rear] = value
	this.size++

	return true
}

func (this *MyCircularDeque) DeleteFront() bool {
	if this.IsEmpty() || this.front == -1 {
		return false
	}

	this.front = (this.front + 1) % cap(this.q)
	this.size--

	return true
}

func (this *MyCircularDeque) DeleteLast() bool {
	if this.IsEmpty() || this.rear == -1 {
		return false
	}

	this.rear = (this.rear - 1 + cap(this.q)) % cap(this.q)
	this.size--

	return true
}

func (this *MyCircularDeque) GetFront() int {
	if this.IsEmpty() || this.front == -1 {
		return -1
	}

	return this.q[this.front]
}

func (this *MyCircularDeque) GetRear() int {
	if this.IsEmpty() || this.rear == -1 {
		return -1
	}

	return this.q[this.rear]
}

func (this *MyCircularDeque) IsEmpty() bool {
	return this.size == 0
}

func (this *MyCircularDeque) IsFull() bool {
	return this.size == cap(this.q)
}
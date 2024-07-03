type DataStream struct {
	value int
	k     int
	count int
}

func Constructor(value int, k int) DataStream {
	return DataStream{
		value: value,
		k:     k,
		count: 0,
	}
}

func (this *DataStream) Consec(num int) bool {
	if this.value == num {
		this.count++
	} else {
		this.count = 0
	}

	return this.count >= this.k
}

/**
 * Your DataStream object will be instantiated and called as such:
 * obj := Constructor(value, k);
 * param_1 := obj.Consec(num);
 */
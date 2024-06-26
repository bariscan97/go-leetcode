type MyHashMap struct {
	_map map[int]int
}

func Constructor() MyHashMap {
	_map := make(map[int]int)
	return MyHashMap{
		_map: _map,
	}
}

func (this *MyHashMap) Put(key int, value int) {
	this._map[key] = value
}

func (this *MyHashMap) Get(key int) int {
	if _, exists := this._map[key]; exists {
		return this._map[key]
	}
	return -1
}

func (this *MyHashMap) Remove(key int) {
	delete(this._map, key)
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
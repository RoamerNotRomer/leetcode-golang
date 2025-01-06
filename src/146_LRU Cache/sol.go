package sol

import "container/list"

type LRUCache struct {
	cache    map[int]*list.Element
	list     *list.List
	capacity int
}

type Pair struct {
	key int
	val int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		cache:    make(map[int]*list.Element),
		capacity: capacity,
		list:     list.New(),
	}
}

func (this *LRUCache) Get(key int) int {
	elem, ok := this.cache[key]
	if !ok {
		return -1
	}
	this.list.MoveToFront(elem)
	return elem.Value.(Pair).val
}

func (this *LRUCache) Put(key int, value int) {
	if elem, ok := this.cache[key]; ok {
		//hit&update cache
		this.list.MoveToFront(elem)
		elem.Value = Pair{key, value}
	} else {
		if len(this.cache) >= this.capacity {
			//out of memory
			delete(this.cache, this.list.Back().Value.(Pair).key)
			this.list.Remove(this.list.Back())
		}
		newElem := this.list.PushFront(Pair{key, value})
		this.cache[key] = newElem
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

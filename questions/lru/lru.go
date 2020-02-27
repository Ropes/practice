package lru

type LRUCache struct {
	n     int
	keys  map[int]int
	stack []int
}

func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		n:     capacity,
		keys:  make(map[int]int),
		stack: make([]int, capacity),
	}
}

func (l *LRUCache) Get(key int) int {
	v, ok := l.keys[key]
	if !ok {
		return -1
	}
	// TODO: Should the get push this key onto the stack upgrading its priority?
	return v
}

func (l *LRUCache) Put(key int, value int) {
	if len(l.stack) >= l.n {
		// pop one key from LRU
		var popped int
		//popped, l.stack = l.stack[len(l.stack)-1], l.stack[:len(l.stack)-1]
		popped, l.stack = l.stack[0], l.stack[1:]
		delete(l.keys, popped)
	}

	// Add new key
	l.keys[key] = value
	l.stack = append(l.stack, key)
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

package lc

// Benchmark: 96ms 7.4mb | 98%

type elem struct {
	key, val   int
	prev, next *elem
}

type MyHashMap struct {
	m [16][16]*elem
}

/** Initialize your data structure here. */
func Constructor() MyHashMap {
	return MyHashMap{[16][16]*elem{}}
}

/** value will always be non-negative. */
func (this *MyHashMap) Put(key int, value int) {
	e := &elem{key, value, nil, nil}
	bIdx := key % 16
	rIdx := (key / 16) % 16
	elem := this.m[bIdx][rIdx]
	if elem == nil {
		this.m[bIdx][rIdx] = e
	} else {
		for elem != nil {
			if elem.key == key {
				elem.val = value
				break
			}

			// we insert
			if elem.next == nil || elem.next.val > key {
				e.next = elem.next
				e.prev = elem
				elem.next = e
				if e.next != nil {
					e.next.prev = e
				}
				break
			}
			elem = elem.next
		}
	}
}

/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyHashMap) Get(key int) int {
	e := this.m[key%16][(key/16)%16]
	for e != nil {
		if e.key == key {
			return e.val
		}
		e = e.next
	}
	return -1
}

/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMap) Remove(key int) {
	e := this.m[key%16][(key/16)%16]
	for e != nil {
		if e.key == key {
			if e.prev == nil {
				if e.next != nil {
					e.next.prev = nil
				}
				this.m[key%16][(key/16)%16] = e.next
			} else {
				e.prev.next = e.next
				if e.next != nil {
					e.next.prev = e.prev
				}
			}
		}
		e = e.next
	}
}

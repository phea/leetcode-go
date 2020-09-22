package lc

// Benchmark: 72ms 11mb | 97% 35%

type elem struct {
	key  int
	next *elem
}

type MyHashSet struct {
	m [1024]*elem
}

/** Initialize your data structure here. */
func Constructor() MyHashSet {
	return MyHashSet{[1024]*elem{}}
}

func (this *MyHashSet) Add(key int) {
	if this.Contains(key) {
		return
	}

	head := this.m[key%1024]
	if head == nil {
		this.m[key%1024] = &elem{key, nil}
	} else {
		this.m[key%1024] = &elem{key, head}
	}
}

func (this *MyHashSet) Remove(key int) {
	head := this.m[key%1024]
	if head == nil {
		return
	}

	// if key is at the front
	if head.key == key {
		this.m[key%1024] = head.next
		return
	}

	var prev *elem
	for head != nil {
		if head.key == key {
			prev.next = head.next
		}

		prev = head
		head = head.next
	}
}

/** Returns true if this set contains the specified element */
func (this *MyHashSet) Contains(key int) bool {
	head := this.m[key%1024]
	for head != nil {
		if head.key == key {
			return true
		}
		head = head.next
	}
	return false
}

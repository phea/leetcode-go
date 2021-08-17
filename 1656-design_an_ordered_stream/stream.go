package lc

// Time: O(n)
// Benchmark: 76ms 7.4mb | 67% 83%

type OrderedStream struct {
	stream []string
	ptr    int
}

func Constructor(n int) OrderedStream {
	return OrderedStream{make([]string, n), 0}
}

func (this *OrderedStream) Insert(id int, value string) []string {
	this.stream[id-1] = value

	orig := this.ptr
	var n int
	for i := this.ptr; i < len(this.stream); i++ {
		if this.stream[i] == "" {
			break
		}
		n++
	}

	this.ptr = orig + n
	if this.ptr >= len(this.stream) {
		return this.stream[orig:]
	}

	return this.stream[orig:this.ptr]
}

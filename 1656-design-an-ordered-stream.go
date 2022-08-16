package main

type OrderedStream struct {
	prt    int
	stream []string
}

func Constructor(n int) OrderedStream {
	return OrderedStream{stream: make([]string, n)}
}

func (this *OrderedStream) Insert(idKey int, value string) (result []string) {
	this.stream[idKey-1] = value
	if idKey-1 == this.prt {
		for ; this.prt < len(this.stream); this.prt++ {
			if this.stream[this.prt] == "" {
				break
			}
			result = append(result, this.stream[this.prt])
		}
	}
	return result
}

/**
 * Your OrderedStream object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Insert(idKey,value);
 */

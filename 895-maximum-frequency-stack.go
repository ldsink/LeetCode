package main

type FreqStack struct {
	valFreq   map[int]int
	freqCount map[int]int
	maxFreq   int
	stack     []int
}

func Constructor() FreqStack {
	return FreqStack{valFreq: make(map[int]int), freqCount: make(map[int]int)}
}

func (this *FreqStack) Push(val int) {
	this.stack = append(this.stack, val)
	if this.valFreq[val] > 0 {
		this.freqCount[this.valFreq[val]]--
		if this.freqCount[this.valFreq[val]] == 0 {
			delete(this.freqCount, this.valFreq[val])
		}
	}
	this.valFreq[val]++
	this.freqCount[this.valFreq[val]]++
	if this.maxFreq < this.valFreq[val] {
		this.maxFreq = this.valFreq[val]
	}
}

func (this *FreqStack) Pop() int {
	// 取出当前的数值
	var idx int
	for i := len(this.stack) - 1; i >= 0; i-- {
		if this.valFreq[this.stack[i]] == this.maxFreq {
			idx = i
			break
		}
	}
	val := this.stack[idx]
	this.stack = append(this.stack[:idx], this.stack[idx+1:]...)
	// 更新频率数据
	this.freqCount[this.valFreq[val]]--
	if this.freqCount[this.valFreq[val]] == 0 {
		delete(this.freqCount, this.valFreq[val])
		this.maxFreq = 0
	}
	this.valFreq[val]--
	this.freqCount[this.valFreq[val]]++
	if this.maxFreq == 0 {
		this.maxFreq = this.valFreq[val]
	}
	return val
}

/**
 * Your FreqStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * param_2 := obj.Pop();
 */

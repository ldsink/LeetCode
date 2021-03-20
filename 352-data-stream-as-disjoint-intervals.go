package main

type SummaryRanges struct {
	intervals [][]int
}

/** Initialize your data structure here. */
func Constructor() SummaryRanges {
	return SummaryRanges{}
}

func (this *SummaryRanges) AddNum(val int) {
	if len(this.intervals) == 0 {
		this.intervals = append(this.intervals, []int{val, val})
		return
	}
	start := 0
	end := len(this.intervals) - 1
	for start < end {
		middle := (start + end) >> 1
		if val <= this.intervals[middle][1] {
			end = middle
		} else {
			start = middle + 1
		}
	}
	idx := start
	if this.intervals[idx][0] <= val && val <= this.intervals[idx][1] {
		return
	} else if val < this.intervals[idx][0] {
		intervals := make([][]int, idx)
		copy(intervals, this.intervals[:idx])
		intervals = append(intervals, []int{val, val})
		intervals = append(intervals, this.intervals[idx:]...)
		this.intervals = intervals
	} else {
		intervals := make([][]int, idx+1)
		copy(intervals, this.intervals[:idx+1])
		intervals = append(intervals, []int{val, val})
		intervals = append(intervals, this.intervals[idx+1:]...)
		this.intervals = intervals
		idx += 1
	}
	// try combine
	if idx-1 >= 0 && this.intervals[idx-1][1]+1 == this.intervals[idx][0] {
		this.intervals[idx-1][1] = this.intervals[idx][1]
		this.intervals = append(this.intervals[:idx], this.intervals[idx+1:]...)
		idx = idx - 1
	}
	if idx+1 < len(this.intervals) && this.intervals[idx][1]+1 == this.intervals[idx+1][0] {
		this.intervals[idx][1] = this.intervals[idx+1][1]
		this.intervals = append(this.intervals[:idx+1], this.intervals[idx+2:]...)
	}
}

func (this *SummaryRanges) GetIntervals() [][]int {
	return this.intervals
}

/**
 * Your SummaryRanges object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(val);
 * param_2 := obj.GetIntervals();
 */

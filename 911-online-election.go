package main

type TopVotedCandidate struct {
	top   []int
	times []int
}

func Constructor(persons []int, times []int) TopVotedCandidate {
	top := make([]int, len(times))
	var max, person int
	count := make([]int, len(persons))
	for i := 0; i < len(times); i++ {
		count[persons[i]]++
		if max <= count[persons[i]] {
			max = count[persons[i]]
			person = persons[i]
		}
		top[i] = person
	}
	return TopVotedCandidate{top: top, times: times}
}

func (this *TopVotedCandidate) Q(t int) int {
	if t >= this.times[len(this.times)-1] {
		return this.top[len(this.top)-1]
	}
	for start, end := 0, len(this.times)-1; start <= end; {
		if start == end {
			return this.top[start]
		}
		middle := (start + end) >> 1
		if t < this.times[middle+1] {
			end = middle
		} else {
			start = middle + 1
		}
	}
	return 0
}

/**
 * Your TopVotedCandidate object will be instantiated and called as such:
 * obj := Constructor(persons, times);
 * param_1 := obj.Q(t);
 */

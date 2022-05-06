package main

type RecentCounter struct {
	start int
	pings []int
}

func Constructor() RecentCounter {
	return RecentCounter{}
}

func (this *RecentCounter) Ping(t int) int {
	this.pings = append(this.pings, t)
	for this.pings[this.start] < t-3000 {
		this.start += 1
	}
	return len(this.pings) - this.start
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */

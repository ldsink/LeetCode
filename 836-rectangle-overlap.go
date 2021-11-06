package main

func isRectangleOverlap(rec1 []int, rec2 []int) bool {
	// rec1 horizon / rec2 vertical
	// (rec1[0], rec1[1]) (rec1[2], rec1[1])
	// (rec1[0], rec1[3]) (rec1[2], rec1[3])
	// (rec2[0], rec2[1]) (rec2[0], rec2[3])
	// (rec2[2], rec2[1]) (rec2[2], rec2[3])
	if rec1[0] < rec2[0] && rec2[0] < rec1[2] && rec2[1] < rec1[1] && rec1[1] < rec2[3] {
		return true
	}
	if rec1[0] < rec2[0] && rec2[0] < rec1[2] && rec2[1] < rec1[1] && rec1[1] < rec2[3] {
		return true
	}
	if rec1[0] < rec2[0] && rec2[0] < rec1[2] && rec2[1] < rec1[3] && rec1[3] < rec2[3] {
		return true
	}
	if rec1[0] < rec2[2] && rec2[2] < rec1[2] && rec2[1] < rec1[3] && rec1[3] < rec2[3] {
		return true
	}
	// rec1 vertical / rec2 horizon
	// (rec1[0], rec1[1]) (rec1[0], rec1[3])
	// (rec1[2], rec1[1]) (rec1[2], rec1[3])
	// (rec2[0], rec2[1]) (rec2[2], rec2[1])
	// (rec2[0], rec2[3]) (rec2[2], rec2[3])
	if rec2[0] < rec1[0] && rec1[0] < rec2[2] && rec1[1] < rec2[1] && rec2[1] < rec1[3] {
		return true
	}
	if rec2[0] < rec1[0] && rec1[0] < rec2[2] && rec1[1] < rec2[1] && rec2[1] < rec1[3] {
		return true
	}
	if rec2[0] < rec1[0] && rec1[0] < rec2[2] && rec1[1] < rec2[3] && rec2[3] < rec1[3] {
		return true
	}
	if rec2[0] < rec1[2] && rec1[2] < rec2[2] && rec1[1] < rec2[3] && rec2[3] < rec1[3] {
		return true
	}
	// full contain
	if rec1[0] <= rec2[0] && rec1[1] <= rec2[1] && rec2[2] <= rec1[2] && rec2[3] <= rec1[3] {
		return true
	}
	if rec2[0] <= rec1[0] && rec2[1] <= rec1[1] && rec1[2] <= rec2[2] && rec1[3] <= rec2[3] {
		return true
	}
	// partial contain
	if rec1[0] == rec2[0] && rec1[2] == rec2[2] && ((rec1[1] < rec2[1] && rec2[1] < rec1[3]) || (rec1[1] < rec2[3] && rec2[3] < rec1[3])) {
		return true
	}
	if rec1[1] == rec2[1] && rec1[3] == rec2[3] && ((rec1[0] < rec2[0] && rec2[0] < rec1[2]) || (rec1[0] < rec2[2] && rec2[2] < rec1[2])) {
		return true
	}
	return false
}

package main

func categorizeBox(length int, width int, height int, mass int) string {
	var isBulky, isHeavy bool
	const p4 = 10000
	const p9 = 1000000000
	if length >= p4 || width >= p4 || height >= p4 || mass >= p4 || length*width*height >= p9 {
		isBulky = true
	}
	if mass >= 100 {
		isHeavy = true
	}
	if isBulky && isHeavy {
		return "Both"
	} else if !isBulky && !isHeavy {
		return "Neither"
	} else if isBulky {
		return "Bulky"
	} else {
		return "Heavy"
	}
}

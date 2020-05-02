package main

func computeArea(A int, B int, C int, D int, E int, F int, G int, H int) int {
	var minTop, maxBottom, maxLeft, minRight int

	if D < H {
		minTop = D
	} else {
		minTop = H
	}

	if B < F {
		maxBottom = F
	} else {
		maxBottom = B
	}

	if A < E {
		maxLeft = E
	} else {
		maxLeft = A
	}

	if C < G {
		minRight = C
	} else {
		minRight = G
	}

	area := (C-A)*(D-B) + (G-E)*(H-F)
	if maxLeft <= minRight && maxBottom <= minTop {
		area -= (minRight - maxLeft) * (minTop - maxBottom)
	}
	return area
}

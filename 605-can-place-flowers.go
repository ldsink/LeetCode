package main

func canPlaceFlowers(flowerbed []int, n int) bool {
	var i, j, k int
	const Flower = 1
	for i < len(flowerbed) && n > 0 {
		if flowerbed[i] == Flower {
			i++
			continue
		}
		for j = i + 1; j < len(flowerbed); j++ {
			if flowerbed[j] == Flower {
				break
			}
		}
		k = j - i
		if i == 0 {
			k++
		}
		if j == len(flowerbed) {
			k++
		}
		n -= (k - 1) >> 1
		i = j + 1 // skip flower
	}
	return n <= 0
}

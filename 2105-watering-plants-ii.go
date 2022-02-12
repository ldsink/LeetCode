package main

func fillWater(water, capacity, needed int) (int, int) {
	var count int
	for needed > 0 {
		if water >= needed { // 直接浇水
			water -= needed
			needed = 0
		} else if water != capacity { // 水量不足以浇水，所以先重新灌满水
			count++
			water = capacity
		} else { // 当前的水应该是满的，但是依旧不足以浇水
			if needed%capacity == 0 {
				count += needed/capacity - 1
				water = 0
			} else {
				count += needed / capacity
				water = capacity - (needed % capacity)
			}
			needed = 0
		}
	}
	return water, count
}

func minimumRefill(plants []int, capacityA int, capacityB int) int {
	var count int
	middle := len(plants) >> 1
	alice := capacityA
	bob := capacityB
	for i := 0; i < middle; i++ {
		var c int
		// Alice
		alice, c = fillWater(alice, capacityA, plants[i])
		count += c
		// Bob
		bob, c = fillWater(bob, capacityB, plants[len(plants)-1-i])
		count += c
	}
	// 相遇的情况
	if len(plants)%2 == 1 {
		if alice >= bob {
			_, c := fillWater(alice, capacityA, plants[middle])
			count += c
		} else {
			_, c := fillWater(bob, capacityB, plants[middle])
			count += c
		}
	}
	return count
}

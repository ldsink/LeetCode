package main

const trieBitLen = 14

type trieNode struct {
	children [2]*trieNode
	count    int
}

type trieTree struct {
	root *trieNode
}

func (t *trieTree) add(val int) {
	for i, current := trieBitLen, t.root; i >= 0; i-- {
		bit := (val >> i) & 1
		if current.children[bit] == nil {
			current.children[bit] = &trieNode{}
		}
		current = current.children[bit]
		current.count++
	}
}

func (t *trieTree) countLt(val, max int) (count int) {
	for i, current := trieBitLen, t.root; i >= 0; i-- {
		bit := (val >> i) & 1
		if (max>>i)&1 > 0 { // max 目标值此位是 1
			if current.children[bit] != nil {
				count += current.children[bit].count // 前缀树中 i 位与 val 相同，xor 值为 0，一定小于 max，直接统计数量
			}
			bit ^= 1
		}
		if current.children[bit] == nil {
			return count
		}
		current = current.children[bit]
	}
	return count
}

func countPairs(nums []int, low, high int) int {
	t := &trieTree{&trieNode{}}
	var result int
	for i := 0; i < len(nums); i++ {
		result += t.countLt(nums[i], high+1) - t.countLt(nums[i], low)
		t.add(nums[i])
	}
	return result
}

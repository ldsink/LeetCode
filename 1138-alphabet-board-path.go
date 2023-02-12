package main

func alphabetBoardPath(target string) string {
	const MoveLeft = 'L'
	const MoveRight = 'R'
	const MoveUp = 'U'
	const MoveDown = 'D'
	var result []rune
	var current int
	for _, r := range target {
		idx := int(r - 'a')
		if current != idx {
			if idx == 25 { // z
				for col := current % 5; col > 0; col-- {
					result = append(result, MoveLeft)
				}
				for row := current / 5; row < 5; row++ {
					result = append(result, MoveDown)
				}
			} else {
				curRow, curCol := current/5, current%5
				dstRow, dstCol := idx/5, idx%5
				for i := curRow; i < dstRow; i++ {
					result = append(result, MoveDown)
				}
				for i := curRow; i > dstRow; i-- {
					result = append(result, MoveUp)
				}
				for i := curCol; i < dstCol; i++ {
					result = append(result, MoveRight)
				}
				for i := curCol; i > dstCol; i-- {
					result = append(result, MoveLeft)
				}
			}
			current = idx
		}
		result = append(result, '!')
	}
	return string(result)
}

package main

func originalDigits(s string) string {
	result := make(map[int]int)
	count := make([]rune, 26)
	for _, r := range []rune(s) {
		count[r-'a']++
	}

	// process special rune: 0 Zero, 2 tWo, 4 foUr, 6 siX, 8 eiGht
	for i := int('z' - 'a'); count[i] > 0; {
		result[0]++
		count['z'-'a']--
		count['e'-'a']--
		count['r'-'a']--
		count['o'-'a']--
	}
	for i := int('w' - 'a'); count[i] > 0; {
		result[2]++
		count['t'-'a']--
		count['w'-'a']--
		count['o'-'a']--
	}
	for i := int('u' - 'a'); count[i] > 0; {
		result[4]++
		count['f'-'a']--
		count['o'-'a']--
		count['u'-'a']--
		count['r'-'a']--
	}
	for i := int('x' - 'a'); count[i] > 0; {
		result[6]++
		count['s'-'a']--
		count['i'-'a']--
		count['x'-'a']--
	}
	for i := int('g' - 'a'); count[i] > 0; {
		result[8]++
		count['e'-'a']--
		count['i'-'a']--
		count['g'-'a']--
		count['h'-'a']--
		count['t'-'a']--
	}

	// 1 One, 3 thRee
	for i := int('o' - 'a'); count[i] > 0; {
		result[1]++
		count['o'-'a']--
		count['n'-'a']--
		count['e'-'a']--
	}
	for i := int('r' - 'a'); count[i] > 0; {
		result[3]++
		count['t'-'a']--
		count['h'-'a']--
		count['r'-'a']--
		count['e'-'a'] -= 2
	}

	// 5 Five
	for i := int('f' - 'a'); count[i] > 0; {
		result[5]++
		count['f'-'a']--
		count['i'-'a']--
		count['v'-'a']--
		count['e'-'a']--
	}

	// 7 Seven / 9 nIne
	for i := int('s' - 'a'); count[i] > 0; {
		result[7]++
		count['s'-'a']--
		count['e'-'a']--
		count['v'-'a']--
		count['e'-'a']--
		count['n'-'a']--
	}
	for i := int('i' - 'a'); count[i] > 0; {
		result[9]++
		count['n'-'a']--
		count['i'-'a']--
		count['n'-'a']--
		count['e'-'a']--
	}

	var rs []rune
	for i := 0; i < 10; i++ {
		for ; result[i] > 0; result[i]-- {
			rs = append(rs, rune('0'+i))
		}
	}
	return string(rs)
}

package game_19_9_15

func maxNumberOfBalloons(text string) int {
	charMap := make(map[string]int, 5)
	charMap["a"] = 0
	charMap["b"] = 0
	charMap["l"] = 0
	charMap["o"] = 0
	charMap["n"] = 0
	for _, v := range text {
		switch v {
		case 'a':
			charMap["a"]++
		case 'b':
			charMap["b"]++
		case 'l':
			charMap["l"]++
		case 'o':
			charMap["o"]++
		case 'n':
			charMap["n"]++
		default:
			continue
		}
	}
	count := 0
	for {
		if charMap["a"] >= 1 && charMap["b"] >= 1 && charMap["l"] >= 2 && charMap["o"] >= 2 && charMap["n"] >= 1 {
			count++
			charMap["a"]--
			charMap["b"]--
			charMap["l"] -= 2
			charMap["o"] -= 2
			charMap["n"]--
		} else {
			break
		}
	}
	return count
}

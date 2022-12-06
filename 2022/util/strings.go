package util

func CountRuneOccurences(r rune, rs []rune) int {
	var count int
	for _, rc := range rs {
		if r == rc {
			count++
		}
	}
	return count
}

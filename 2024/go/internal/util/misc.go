package util

// IntAbs returns the absolute value of an integer
func IntAbs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

package util

func IntAbs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func IntMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func IntMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

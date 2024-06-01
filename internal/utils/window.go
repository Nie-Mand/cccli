package utils

func GetWindow(l int, c int, s int) (int, int) {
	if c < 0 {
		return 0, 0
	}

	// If in the first page
	if c < s {
		return 0, s % l
	}

	// If in the middle pages
	return c - s + 1, (c + 1) % l
}

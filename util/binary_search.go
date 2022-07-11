package util

func BinarySearch(needle int, haystack []int) bool {
	l := 0
	h := len(haystack) - 1

	for l <= h {
		m := (h + l) / 2

		if haystack[m] < needle {
			l = m + 1
		} else {
			h = m - 1
		}
	}

	if l == len(haystack) || haystack[l] != needle {
		return false
	}

	return true
}

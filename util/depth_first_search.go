package util

var items []int
var length int
var want int

func DepthFirstSerach(w int, s []int) bool {
	items = s
	length = len(items)
	want = w

	return depthFirstSearch(0, 0)
}

func depthFirstSearch(i, sum int) bool {
	if i == length {
		return sum == want
	}

	if depthFirstSearch(i+1, sum) {
		return true
	}

	if depthFirstSearch(i+1, sum+items[i]) {
		return true
	}

	return false
}

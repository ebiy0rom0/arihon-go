package triangle

import "sort"

func Calc(cnt int, lines []int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(lines)))

	max := len(lines) - 1 // スライスのインデックス考慮
	for i, l := range lines {
		if i+2 > max {
			break
		}
		if l < lines[i+1]+lines[i+2] {
			return l + lines[i+1] + lines[i+2]
		}
	}

	return 0
}

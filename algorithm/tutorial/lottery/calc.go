package lottery

import (
	"arihon-go/util"
)

var stack []int

func Calc(n, m int, k []int) bool {
	stack = make([]int, n*(n+1)/2-1)

	for i := 0; i < n; i++ {
		stack[i] = k[0] + k[i]
		if i > 0 {
			stack[n+i-1] = k[i] + k[n-1]
		}
	}

	for _, kk := range stack {
		if m-kk > m {
			return false
		}
		if util.BinarySearch(m-kk, stack) {
			return true
		}
	}

	return false
}

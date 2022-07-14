package partial_sum

import "arihon-go/util"

func Calc(n int, a []int, k int) bool {
	return util.DepthFirstSerach(k, a)
}

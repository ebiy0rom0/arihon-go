package util

import (
	"math"
	"math/rand"
	"time"
)

// ランダムの生成にコストがかかっては本末転倒なので、
// 必要数に到達していなくても試行回数の限界値を設定しておく
var limit = math.Pow10(6)

func MakeRandomSlice(min, max, needs int) []int {
	nRange := max - min

	// 最終系が確定してるので、ランダムで生成しない
	if nRange < needs {
		slice := make([]int, nRange)
		for i := 0; i < nRange; i++ {
			slice[i] = min + i
		}
		return slice
	}

	rand.Seed(time.Now().UnixMicro())

	picked := make(map[int]bool, needs)
	for i, limiter := 0, 0; i < needs; {
		n := rand.Intn(nRange) + min
		if !picked[n] {
			picked[n] = true
			i++
		}

		limiter++
		if limiter >= int(limit) {
			break
		}
	}

	res := make([]int, len(picked))
	k := 0
	for n := range picked {
		res[k] = n
		k++
	}

	return res
}

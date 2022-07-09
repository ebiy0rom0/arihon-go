package util

import (
	"math"
	"math/rand"
	"time"
)

// ランダムの生成にコストがかかっては本末転倒なので、
// 必要数に到達していなくても試行回数の限界値を設定しておく
var limit = math.Pow10(6)

func MakeRandom(min, max, needs int) []int {
	nRange := max - min

	rand.Seed(time.Now().UnixMicro())

	picked := make(map[int]bool)
	for i, limiter := 0, 0; i < needs; i++ {
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
	for n := range picked {
		res = append(res, n)
	}

	return res
}

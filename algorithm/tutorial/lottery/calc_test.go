package lottery_test

import (
	"arihon-go/algorithm/tutorial/lottery"
	"arihon-go/util"
	"math"
	"testing"
)

func TestCalc(t *testing.T) {
	type args struct {
		n int
		m int
		k []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{name: "simple", args: args{n: 3, m: 10, k: []int{1, 3, 5}}, want: true},
		{name: "no answer", args: args{n: 3, m: 9, k: []int{1, 3, 5}}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lottery.Calc(tt.args.n, tt.args.m, tt.args.k); got != tt.want {
				t.Errorf("Calc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkCalc(b *testing.B) {
	nums := util.MakeRandomSlice(1, int(math.Pow10(8)), 1000)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// this logic answer is `false` = not fast return
		lottery.Calc(len(nums), 1, nums)
	}
}

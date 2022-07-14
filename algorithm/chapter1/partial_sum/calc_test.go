package partial_sum_test

import (
	"arihon-go/algorithm/chapter1/partial_sum"
	"arihon-go/util"
	"math"
	"testing"
)

func TestCalc(t *testing.T) {
	type args struct {
		n int
		a []int
		k int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{name: "simple", args: args{n: 4, a: []int{1, 2, 4, 7}, k: 13}, want: true},
		{name: "no answer", args: args{n: 4, a: []int{1, 2, 4, 7}, k: 15}, want: false},
		{name: "include minus", args: args{n: 4, a: []int{-5, -1, 3, 9}, k: 7}, want: true},
		{name: "no answer minus", args: args{n: 4, a: []int{-5, -1, 3, 9}, k: 5}, want: false},
		{name: "multi answers", args: args{n: 4, a: []int{-5, -1, 3, 9}, k: 3}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := partial_sum.Calc(tt.args.n, tt.args.a, tt.args.k); got != tt.want {
				t.Errorf("Calc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkCalc(b *testing.B) {
	stack := util.MakeRandomSlice(int(math.Pow10(-8)), int(math.Pow10(8)), 20)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		partial_sum.Calc(20, stack, 10000)
	}
}

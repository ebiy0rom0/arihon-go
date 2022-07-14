package util_test

import (
	"arihon-go/util"
	"math"
	"testing"
)

func TestDepthFirstSerach(t *testing.T) {
	type args struct {
		w int
		s []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{name: "simple", args: args{w: 13, s: []int{1, 2, 4, 7}}, want: true},
		{name: "no answer", args: args{w: 15, s: []int{1, 2, 4, 7}}, want: false},
		{name: "include minus", args: args{w: 7, s: []int{-5, -1, 3, 9}}, want: true},
		{name: "no answer minus", args: args{w: 5, s: []int{-5, -1, 3, 9}}, want: false},
		{name: "multi answers", args: args{w: 3, s: []int{-5, -1, 3, 9}}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := util.DepthFirstSerach(tt.args.w, tt.args.s); got != tt.want {
				t.Errorf("DepthFirstSerach() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkDepthFirstSearch(b *testing.B) {
	stack := util.MakeRandomSlice(int(math.Pow10(-8)), int(math.Pow10(8)), 20)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		util.DepthFirstSerach(100, stack)
	}
}

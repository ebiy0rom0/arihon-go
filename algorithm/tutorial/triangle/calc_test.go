package triangle_test

import (
	"arihon-go/algorithm/tutorial/triangle"
	"arihon-go/util"
	"math"
	"testing"
)

func TestCalc(t *testing.T) {
	type args struct {
		cnt   int
		lines []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{name: "simple", args: args{cnt: 5, lines: []int{2, 3, 4, 5, 10}}, want: 12},
		{name: "no answer", args: args{cnt: 4, lines: []int{4, 5, 10, 20}}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := triangle.Calc(tt.args.cnt, tt.args.lines); got != tt.want {
				t.Errorf("Calc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkCalc(b *testing.B) {
	lines := util.MakeRandomSlice(1, int(math.Pow10(6)), 100)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		triangle.Calc(len(lines), lines)
	}
}

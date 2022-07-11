package ants_test

import (
	"arihon-go/algorithm/tutorial/ants"
	"arihon-go/util"
	"math"
	"testing"
)

func TestCalc(t *testing.T) {
	type args struct {
		l   int
		n   int
		pos []int
	}
	tests := []struct {
		name    string
		args    args
		wantMin int
		wantMax int
	}{
		// TODO: Add test cases.
		{name: "simple", args: args{l: 10, n: 3, pos: []int{2, 6, 7}}, wantMin: 4, wantMax: 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotMin, gotMax := ants.Calc(tt.args.l, tt.args.n, tt.args.pos)
			if gotMin != tt.wantMin {
				t.Errorf("Calc() gotMin = %v, want %v", gotMin, tt.wantMin)
			}
			if gotMax != tt.wantMax {
				t.Errorf("Calc() gotMax = %v, want %v", gotMax, tt.wantMax)
			}
		})
	}
}

func BenchmarkCalc(b *testing.B) {
	pos := util.MakeRandomSlice(1, int(math.Pow10(6)), int(math.Pow10(6)))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ants.Calc(int(math.Pow10(6)), len(pos), pos)
	}
}

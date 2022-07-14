package lake_counting_test

import (
	"arihon-go/algorithm/chapter1/lake_counting"
	"testing"
)

func TestCalc(t *testing.T) {
	type args struct {
		n int
		m int
		f [][]string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{name: "simple", args: args{n: 10, m: 12, f: [][]string{
			{"w", ".", ".", ".", ".", ".", ".", ".", ".", "w", "w", "."},
			{".", "w", "w", "w", ".", ".", ".", ".", ".", "w", "w", "w"},
			{".", ".", ".", ".", "w", "w", ".", ".", ".", "w", "w", "."},
			{".", ".", ".", ".", ".", ".", ".", ".", ".", "w", "w", "."},
			{".", ".", ".", ".", ".", ".", ".", ".", ".", "w", ".", "."},
			{".", ".", "w", ".", ".", ".", ".", ".", ".", "w", ".", "."},
			{".", "w", ".", "w", ".", ".", ".", ".", ".", "w", "w", "."},
			{"w", ".", "w", ".", "w", ".", ".", ".", ".", ".", "w", "."},
			{".", "w", ".", "w", ".", ".", ".", ".", ".", ".", "w", "."},
			{".", ".", "w", ".", ".", ".", ".", ".", ".", ".", "w", "."},
		}}, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lake_counting.Calc(tt.args.n, tt.args.m, tt.args.f); got != tt.want {
				t.Errorf("Calc() = %v, want %v", got, tt.want)
			}
		})
	}
}

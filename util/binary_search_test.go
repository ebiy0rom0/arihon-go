package util_test

import (
	"arihon-go/util"
	"math"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	type args struct {
		needle   int
		haystack []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{name: "simple", args: args{needle: 5, haystack: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}}, want: true},
		{name: "no number", args: args{needle: 5, haystack: []int{1, 2, 3, 4, 6, 7, 8, 9, 10}}, want: false},
		{name: "first position", args: args{needle: 1, haystack: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}}, want: true},
		{name: "last position", args: args{needle: 10, haystack: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}}, want: true},
		{name: "over position", args: args{needle: 100, haystack: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}}, want: false},
		{name: "hole exist", args: args{needle: 13, haystack: []int{1, 3, 4, 6, 7, 8, 13, 15, 18, 20}}, want: true},
		{name: "hole no number", args: args{needle: 9, haystack: []int{1, 3, 4, 6, 7, 8, 13, 15, 18, 20}}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := util.BinarySearch(tt.args.needle, tt.args.haystack); got != tt.want {
				t.Errorf("BinarySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkCalc(b *testing.B) {
	stack := util.MakeRandomSlice(1, int(math.Pow10(6)), int(math.Pow10(6)))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		util.BinarySearch(len(stack)/2, stack)
	}
}

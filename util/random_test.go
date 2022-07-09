package util_test

import (
	"arihon-go/util"
	"testing"
)

func TestMakeRandomSlice(t *testing.T) {
	type args struct {
		min   int
		max   int
		needs int
	}
	tests := []struct {
		name string
		args args
		want int // 期待するスライスのレンジ
	}{
		// min > max のようなイレギュラー系はサポートしてないので今回省略
		{name: "simple", args: args{min: 1, max: 10, needs: 5}, want: 5},
		{name: "size to range", args: args{min: 1, max: 10, needs: 100}, want: 10},
		{name: "big", args: args{min: 1, max: 100000, needs: 1000}, want: 1000},
		{name: "max", args: args{min: 1, max: 10000000, needs: 10000000}, want: 10000000},
		{name: "range not start 1", args: args{min: 1000, max: 1000000, needs: 1000}, want: 1000},
	}
	// ランダム生成のため、同テストケースでも結果が変わる
	// 各ケース100回試行し、エラーがなければ機能として十分と判断
	// テストパフォーマンスに影響し出したら試行回数を減らす
	for i := 0; i < 100; i++ {
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got := util.MakeRandomSlice(tt.args.min, tt.args.max, tt.args.needs)
				// レンジチェック(リミット到達時にサイズ ≠ 期待値となるので、超えてなければOK)
				if len(got) > tt.want {
					t.Errorf("range error = %d, want %d", len(got), tt.want)
				}

				// 返却されたスライスの最大値・最小値が適切か
				min := got[0]
				max := got[len(got)-1]
				if min < tt.args.min {
					t.Errorf("slice min out of range = %d, min limit %d", min, tt.args.min)
				}
				if max > tt.args.max {
					t.Errorf("slice max out of range = %d, max limit %d", max, tt.args.max)
				}

				// ソート済みか(仕組み上、ソートされている)
				before := -1
				for _, n := range got {
					if n <= before {
						t.Errorf("return slice is not sorted = %v", got)
					}
				}
			})
		}
	}
}

package max_sub_sum

import "testing"

func TestMaxSubSum(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "[34, -50, 42, 14, -5, 86]",
			args: args{nums: []int{34, -50, 42, 14, -5, 86}},
			want: 137,
		},
		{
			name: "[-5, -1, -8, -9] ",
			args: args{nums: []int{-5, -1, -8, -9}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxSubSum(tt.args.nums); got != tt.want {
				t.Errorf("MaxSubSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

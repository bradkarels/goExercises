package main

import "testing"

func Test_findMaxAverage(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "test case - 0",
			args: args{
				nums: []int{1, 12, -5, -6, 50, 3},
				k:    4,
			},
			want: 12.75,
		},
		{
			name: "test case - 1",
			args: args{
				nums: []int{5},
				k:    1,
			},
			want: 5.0000,
		},
		{
			name: "test case - 2",
			args: args{
				nums: []int{-1},
				k:    1,
			},
			want: -1.0000,
		},
		{
			name: "test case - 3",
			args: args{
				nums: []int{0, 1, 1, 3, 3},
				k:    4,
			},
			want: 2.0000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMaxAverage(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("findMaxAverage() = %v, want %v", got, tt.want)
			}
		})
	}
}

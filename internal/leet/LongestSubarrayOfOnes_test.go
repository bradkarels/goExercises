package main

import "testing"

func Test_longestSubarray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test case one",
			args: args{
				nums: []int{1, 1, 0, 1},
			},
			want: 3,
		},
		{
			name: "test case two",
			args: args{
				nums: []int{0, 1, 1, 1, 0, 1, 1, 0, 1},
			},
			want: 5,
		},
		{
			name: "test case three",
			args: args{
				nums: []int{1, 1, 1},
			},
			want: 2,
		},
		{
			name: "test case four",
			args: args{
				nums: []int{1},
			},
			want: 0,
		},
		{
			name: "test case five",
			args: args{
				nums: []int{1, 0, 0, 0, 0, 1, 1, 0, 1},
			},
			want: 3,
		},
		{
			name: "test case six",
			args: args{
				nums: []int{1, 1, 0},
			},
			want: 2,
		},
		{
			name: "test case four",
			args: args{
				nums: []int{0},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestSubarray(tt.args.nums); got != tt.want {
				t.Errorf("longestSubarray() = %v, want %v", got, tt.want)
			}
		})
	}
}

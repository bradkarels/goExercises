package main

import "testing"

func Test_maxVowels(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test case - 0",
			args: args{
				"abciiidef",
				3,
			},
			want: 3,
		},
		{
			name: "test case - 1",
			args: args{
				"aeiou",
				2,
			},
			want: 2,
		},
		{
			name: "test case - 0",
			args: args{
				"leetcode",
				3,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxVowels(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("maxVowels() = %v, want %v", got, tt.want)
			}
		})
	}
}

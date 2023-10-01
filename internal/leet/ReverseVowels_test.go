package main

import "testing"

func Test_reverseVowels(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test case 0",
			args: args{
				"hello",
			},
			want: "holle",
		},
		{
			name: "test case 1",
			args: args{
				"leetcode",
			},
			want: "leotcede",
		},
		{
			name: "test case 2",
			args: args{
				"Aa",
			},
			want: "aA",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseVowels(tt.args.s); got != tt.want {
				t.Errorf("reverseVowels() = %v, want %v", got, tt.want)
			}
		})
	}
}

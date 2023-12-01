package cracked

import "testing"

func Test_isPalindromePermutation(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: `test case 1`,
			args: args{
				s: `man aaa plan Canal panama`,
			},
			want: true,
		},
		{
			name: `test case 0`,
			args: args{
				s: `Taco cat`,
			},
			want: true,
		},
		{
			name: `test case 2`,
			args: args{
				s: `Holy Shit!`,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindromePermutation(tt.args.s); got != tt.want {
				t.Errorf("isPalindromePermutation() = %v, want %v", got, tt.want)
			}
		})
	}
}

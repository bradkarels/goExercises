package main

import "testing"

func Test_gcdOfStrings(t *testing.T) {
	type args struct {
		str1 string
		str2 string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test case 0",
			args: args{
				str1: "ABCABC",
				str2: "ABC",
			},
			want: "ABC",
		},
		{
			name: "test case 1",
			args: args{
				str1: "ABABAB",
				str2: "ABAB",
			},
			want: "AB",
		},
		{
			name: "test case 2",
			args: args{
				str1: "LEET",
				str2: "CODE",
			},
			want: "",
		},
		{
			name: "test case 3",
			args: args{
				str1: "ABC",
				str2: "ABCABC",
			},
			want: "ABC",
		},
		{
			name: "test case 3",
			args: args{
				str1: "TAUXXTAUXXTAUXXTAUXXTAUXX",
				str2: "TAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXX",
			},
			want: "TAUXX",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := gcdOfStrings(tt.args.str1, tt.args.str2); got != tt.want {
				t.Errorf("gcdOfStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}

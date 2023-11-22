package main

import "testing"

func Test_decodeString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test case -1",
			args: args{
				s: "xyz12[abc]",
			},
			want: "xyzabcabcabcabcabcabcabcabcabcabcabcabc",
		},
		{
			name: "test case -2",
			args: args{
				s: "x3[a]x",
			},
			want: "xaaax",
		},
		{
			name: "test case -42",
			args: args{
				s: "3[3[3[a]]]",
			},
			want: "aaaaaaaaaaaaaaaaaaaaaaaaaaa", // (3 * 3 * 3) * a -> 27 a's
		},
		{
			name: "test case 0",
			args: args{
				s: "3[a]2[bc]",
			},
			want: "aaabcbc",
		},
		{
			name: "test case 1",
			args: args{
				s: "3[a2[c]]",
			},
			want: "accaccacc",
		},
		{
			name: "test case 2",
			args: args{
				s: "2[abc]3[cd]ef",
			},
			want: "abcabccdcdcdef",
		},
		{
			name: "test case 3",
			args: args{
				s: "abcdefg",
			},
			want: "abcdefg",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := decodeString(tt.args.s); got != tt.want {
				t.Errorf("decodeString() = %v, want %v", got, tt.want)
			}
		})
	}
}

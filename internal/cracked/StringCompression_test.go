package cracked

import "testing"

func Test_compress(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test 0.2",
			args: args{
				s: "aaaaaaaaaaaaabbbccc",
			},
			want: `a13b3c3`,
		},
		{
			name: "test 0",
			args: args{
				s: "aaabbbccc",
			},
			want: `a3b3c3`,
		},
		{
			name: "test 0.5",
			args: args{
				s: "aaabbbc",
			},
			want: `a3b3c1`,
		},
		{
			name: "test 1",
			args: args{
				s: "aabbcc",
			},
			want: `aabbcc`, // No compression value -> `a2b2c2`
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compress(tt.args.s); got != tt.want {
				t.Errorf("compress() = %v, want %v", got, tt.want)
			}
		})
	}
}

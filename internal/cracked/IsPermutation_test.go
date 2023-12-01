package cracked

import "testing"

func Test_isPermutation(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: `tst case 0`,
			args: args{
				a: "abc",
				b: "abc",
			},
			want: true,
		},
		{
			name: `tst case 1`,
			args: args{
				a: "abc",
				b: "cBa",
			},
			want: true,
		},
		{
			name: `tst case panama`,
			args: args{
				a: "abcxxxxxxxxxxcba",
				b: "XxxXxxxxXxabCbcA",
			},
			want: true,
		},
		{
			name: `tst case panama2`,
			args: args{
				a: "abcxxxxxxxxxxcbX",
				b: "XxxXxxxxXxabCbcA",
			},
			want: false,
		},
		{
			name: `tst case empty`,
			args: args{
				a: "",
				b: "",
			},
			want: true,
		},
		{
			name: `tst case xxx`,
			args: args{
				a: "aa",
				b: "ZZ",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPermutation(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("isPermutation() = %v, want %v", got, tt.want)
			}
		})
	}
}

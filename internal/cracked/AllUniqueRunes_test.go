package cracked

import "testing"

func Test_rAllUniq(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: `test case 0`,
			args: args{
				s: "abcdefghijklmopqrstuvwxyz",
			},
			want: true,
		},
		{
			name: `test case 1`,
			args: args{
				s: "abcdefghijklmopqrstuvwxxxxyz",
			},
			want: false,
		},
		{
			name: `test case empty`,
			args: args{
				s: "",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rAllUniq(tt.args.s); got != tt.want {
				t.Errorf("rAllUniq() = %v, want %v", got, tt.want)
			}
		})
	}
}

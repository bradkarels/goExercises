package cracked

import "testing"

func Test_isRotation(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: `test 0.1`,
			args: args{
				s1: `ababababababa`,
				s2: `babababababaa`,
			},
			want: true,
		},
		{
			name: `test 0.2`,
			args: args{
				s1: `ababababababa`,
				s2: `bababaXababaa`,
			},
			want: false,
		},
		{
			name: `test 0`,
			args: args{
				s1: `wasserflasche`,
				s2: `erflaschewass`,
			},
			want: true,
		},
		{
			name: `test 1`,
			args: args{
				s1: `wasserflasche`,
				s2: `erflaXchewass`,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isRotation(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("isRotation() = %v, want %v", got, tt.want)
			}
		})
	}
}

package main

import "testing"

func Test_addBinary(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: `Test case -4`,
			args: args{
				a: `101010101`,
				b: `1`,
			},
			want: `101010110`,
		},
		{
			name: `Test case -3`,
			args: args{
				a: `1`,
				b: `1`,
			},
			want: `10`,
		},
		{
			name: `Test case -2`,
			args: args{
				a: `111`,
				b: `11`,
			},
			want: `1010`,
		},
		{
			name: `Test case -1`,
			args: args{
				a: `111`,
				b: `111`,
			},
			want: `1110`,
		},
		{
			name: `Test case 0`,
			args: args{
				a: `11`,
				b: `1`,
			},
			want: `100`,
		},
		{
			name: `Test case 1`,
			args: args{
				a: `1010`,
				b: `1011`,
			},
			want: `10101`,
		},
		{
			name: `Test case 1`,
			args: args{
				a: `0`,
				b: `1010`,
			},
			want: `1010`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addBinary(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("addBinary() = %v, want %v", got, tt.want)
			}
		})
	}
}

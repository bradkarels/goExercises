package grok

import "testing"

func Test_search(t *testing.T) {
	type args struct {
		ints []int
		tgt  int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: `test case 0 - found`,
			args: args{
				ints: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				tgt:  9,
			},
			want: true,
		},
		{
			name: `test case 1 - not found`,
			args: args{
				ints: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				tgt:  42,
			},
			want: false,
		},
		{
			name: `test case 2 - found`,
			args: args{
				ints: []int{0, 2, 4, 6, 8, 10},
				tgt:  0,
			},
			want: true,
		},
		{
			name: `test case 3 - not found`,
			args: args{
				ints: []int{1, 3, 5, 7, 9},
				tgt:  9,
			},
			want: true,
		},
		{
			name: `test case 4 - found`,
			args: args{
				ints: []int{0, 2, 4, 6, 8, 10},
				tgt:  8,
			},
			want: true,
		},
		{
			name: `test case 5 - not found`,
			args: args{
				ints: []int{1, 3, 5, 7, 9},
				tgt:  3,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := search(tt.args.ints, tt.args.tgt)
			if got != tt.want {
				t.Errorf("search() got = %v, want %v", got, tt.want)
			}
		})
	}
}

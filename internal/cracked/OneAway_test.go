package cracked

import "testing"

func Test_oneAway(t *testing.T) {
	type args struct {
		valid   string
		attempt string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: `test case len0`,
			args: args{
				valid:   `abcde`,
				attempt: `abc`, // too short
			},
			want: false,
		},
		{
			name: `test case len1`,
			args: args{
				valid:   `abcde`,
				attempt: `abcdefg`, // too long
			},
			want: false,
		},
		{
			name: `test case len2`,
			args: args{
				valid:   `abcde`,
				attempt: `abcd`, // short but good
			},
			want: true,
		},
		{
			name: `test case len3`,
			args: args{
				valid:   `abcde`,
				attempt: `abcdef`, // long but good
			},
			want: true,
		},
		{
			name: `test case len4`,
			args: args{
				valid:   `abcde`,
				attempt: `edcba`, // just right
			},
			want: false,
		},
		{
			name: `test case equal`,
			args: args{
				valid:   `abcde`,
				attempt: `abcde`, // not one off - match!
			},
			want: false,
		},
		{
			name: `test case one off good`,
			args: args{
				valid:   `abcde`,
				attempt: `abXde`, // exactly one non-match
			},
			want: true,
		},
		{
			name: `test case one off bad`,
			args: args{
				valid:   `abcde`,
				attempt: `XbXdX`, // more than 1 non-match
			},
			want: false,
		},
		{
			name: `test case min good simple`,
			args: args{
				valid:   `abcde`,
				attempt: `abde`,
			},
			want: true,
		},
		{
			name: `test case min good simple`,
			args: args{
				valid:   `abcde`,
				attempt: `bcde`,
			},
			want: true,
		},
		{
			name: `test case min bad simple`,
			args: args{
				valid:   `abcde`,
				attempt: `abee`,
			},
			want: false,
		},
		{
			name: `test case min bad simple`,
			args: args{
				valid:   `abcde`,
				attempt: `bcdx`,
			},
			want: false,
		},
		{
			name: `test case one too many good`,
			args: args{
				valid:   `abcde`,
				attempt: `abcdef`,
			},
			want: true,
		},
		{
			name: `test case one too many bad2`,
			args: args{
				valid:   `abcde`,
				attempt: `aacdef`,
			},
			want: false,
		},
		{
			name: `test case one too many bad`,
			args: args{
				valid:   `abcde`,
				attempt: `abcdff`,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := oneAway(tt.args.valid, tt.args.attempt); got != tt.want {
				t.Errorf("oneAway() = %v, want %v", got, tt.want)
			}
		})
	}
}

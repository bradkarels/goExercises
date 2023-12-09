package cracked

import (
	"reflect"
	"testing"
)

var (
	ln7  = ListNode{Val: 3, Next: nil}
	ln6  = ListNode{Val: 5, Next: &ln7}
	ln5  = ListNode{Val: 4, Next: &ln6}
	ln4  = ListNode{Val: 1, Next: &ln5}
	ln3  = ListNode{Val: 3, Next: &ln4}
	ln2  = ListNode{Val: 2, Next: &ln3}
	head = ListNode{Val: 1, Next: &ln2}
)

func Test_remDup(t *testing.T) {
	type args struct {
		h *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// TODO: Add test cases.
		{
			name: `test 0`,
			args: args{
				h: &head,
			},
			want: &head,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := remDup(tt.args.h); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("remDup() = %v, want %v", got, tt.want)
			}
		})
	}
}

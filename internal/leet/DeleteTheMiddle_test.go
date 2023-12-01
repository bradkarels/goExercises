package main

import (
	"reflect"
	"testing"
)

var (
	ln6  = ListNode{Val: 6, Next: nil}
	ln2  = ListNode{Val: 2, Next: &ln6}
	ln1  = ListNode{Val: 1, Next: &ln2}
	ln7  = ListNode{Val: 7, Next: &ln1}
	ln4  = ListNode{Val: 4, Next: &ln7}
	ln3  = ListNode{Val: 3, Next: &ln4}
	head = ListNode{Val: 1, Next: &ln3}
)

func Test_deleteMiddle(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// TODO: Add test cases.
		{
			name: `aa`,
			args: args{
				head: &head,
			},
			want: &head,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteMiddle(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deleteMiddle() = %v, want %v", got, tt.want)
			}
		})
	}
}

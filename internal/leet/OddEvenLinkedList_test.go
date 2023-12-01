package main

import (
	"reflect"
	"testing"
)

var (
	lnx5 = ListNode{Val: 5, Next: nil}
	lnx4 = ListNode{Val: 4, Next: &lnx5}
	lnx3 = ListNode{Val: 3, Next: &lnx4}
	lnx2 = ListNode{Val: 2, Next: &lnx3}
	lnx1 = ListNode{Val: 1, Next: &lnx2}
)

func Test_oddEvenList(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: `asdf`,
			args: args{
				head: &lnx1,
			},
			want: []int{1, 3, 5, 2, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := oddEvenList(tt.args.head)
			var nodeVals = []int{got.Val}
			current := got
			for {
				if current.Next != nil {
					nodeVals = append(nodeVals, current.Next.Val)
					current = current.Next
					continue
				} else {
					break
				}
			}
			if !reflect.DeepEqual(nodeVals, tt.want) {
				t.Errorf("oddEvenList() = %v, want %v", got, tt.want)
			}
		})
	}
}

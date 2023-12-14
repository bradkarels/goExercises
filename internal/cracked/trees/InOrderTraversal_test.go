package trees

import (
	"testing"
)

var (
	nine     = Node{value: 9, left: nil, right: nil}
	eighteen = Node{value: 18, left: nil, right: nil}
	three    = Node{value: 3, left: nil, right: nil}
	seven    = Node{value: 7, left: nil, right: nil}
	five     = Node{value: 5, left: &nine, right: &eighteen}
	twenty   = Node{value: 20, left: &three, right: &seven}
	ten      = Node{value: 10, left: &five, right: &twenty}
)

func TestInOrderTraversal(t *testing.T) {
	type args struct {
		n *Node
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: `can we kick it - perhaps, but can we test it?`,
			args: args{
				n: &ten,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			InOrderTraversal(tt.args.n)
		})
	}
}

func TestPreOrderTraversal(t *testing.T) {
	type args struct {
		n *Node
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: `silly test name for non-test`,
			args: args{
				n: &ten,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			PreOrderTraversal(tt.args.n)
		})
	}
}

func TestPostOrderTraversal(t *testing.T) {
	type args struct {
		n *Node
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: `silly test name for non-test`,
			args: args{
				n: &ten,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			PostOrderTraversal(tt.args.n)
		})
	}
}

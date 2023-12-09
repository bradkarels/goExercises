package cracked

import (
	"fmt"
	"reflect"
	"testing"
)

var (
	e2 = Node{Val: 42, Next: nil}
	e1 = Node{Val: 4, Next: &e2}
	e0 = Node{Val: 27, Next: &e1}
	//empt = Stack{
	//	head: nil,
	//}
)
var (
	q3 = Node{Val: 3, Next: nil}
	q2 = Node{Val: 2, Next: &q3}
	q1 = Node{Val: 1, Next: &q2}
)

//func TestStack_Put(t *testing.T) {
//	type fields struct {
//		head *Node
//	}
//	type args struct {
//		v int
//	}
//	tests := []struct {
//		name   string
//		fields fields
//		args   args
//	}{
//		{
//			name: `Putter....`,
//			fields: fields{
//				head: &e0,
//			},
//			args: args{
//				v: 123,
//			},
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			s := &Stack{
//				head: tt.fields.head,
//			}
//			s.Put(tt.args.v)
//		})
//	}
//}
//
//func TestStack_Peek(t *testing.T) {
//	type fields struct {
//		head *Node
//	}
//	tests := []struct {
//		name   string
//		fields fields
//		want   Node
//	}{
//		{
//			name: `peeky peeky!`,
//			fields: fields{
//				head: &e0,
//			},
//			want: e0,
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			s := &Stack{
//				head: tt.fields.head,
//			}
//			if got := s.Peek(); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("Peek() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//

func TestNode_IsEmpty(t *testing.T) {
	type args struct {
		head *Node
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: `isEmtpy test 0`,
			args: args{
				head: nil,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.head.IsEmpty(); got != tt.want {
				t.Errorf("IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNode_Add(t *testing.T) {
	type fields struct {
		Val  int
		Next *Node
	}
	type args struct {
		i int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name: `le experiment...0`,
			fields: fields{
				Val:  1,
				Next: &Node{Val: 2, Next: &Node{Val: 3, Next: nil}},
			},
			args: args{
				i: 0,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &Node{
				Val:  tt.fields.Val,
				Next: tt.fields.Next,
			}
			n = n.Add(tt.args.i)
			if newHeadVal := n.Val; newHeadVal != tt.want {
				t.Errorf("newHeadVal = %v, want %v", newHeadVal, tt.want)
			}
		})
	}
}

var (
	n3 = Node{Val: 3, Next: nil}
	n2 = Node{Val: 2, Next: &n3}
	n1 = Node{Val: 1, Next: &n2}
)

func TestNode_Pop(t *testing.T) {
	type fields struct {
		Val  int
		Next *Node
	}
	tests := []struct {
		name   string
		fields fields
		want   int
		want2  *Node
	}{
		{
			name: `Haar naam is Roos...`,
			fields: fields{
				Val:  0,
				Next: &n1,
			},
			want:  0,
			want2: &n1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &Node{
				Val:  tt.fields.Val,
				Next: tt.fields.Next,
			}
			if head, got := n.Pop(); !reflect.DeepEqual(got, tt.want) || !reflect.DeepEqual(head, tt.want2) {
				t.Errorf("Pop() = %v, want %v", got, tt.want)
				t.Errorf("Pop() = %v, want2 %v", head.Val, tt.want2)
			}
			fmt.Println(n.Val)
		})
	}
}

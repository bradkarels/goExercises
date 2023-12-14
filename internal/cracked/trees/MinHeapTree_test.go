package trees

import (
	"testing"
)

var (
	fiftyFive  = Node{value: 55, left: nil, right: nil}
	ninety     = Node{value: 90, left: nil, right: nil}
	eightSeven = Node{value: 87, left: nil, right: nil}
	fifty      = Node{value: 50, left: &fiftyFive, right: &ninety}
	zeven      = Node{value: 7, left: &eightSeven, right: nil}
	four       = Node{value: 4, left: &fifty, right: &zeven}
	two        = Node{value: 2, left: nil, right: nil}
)

func TestInsert(t *testing.T) {
	type args struct {
		head   *Node
		nieuwe *Node
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: `display operations`,
			args: args{
				head:   &four,
				nieuwe: &two,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Insert(tt.args.head, tt.args.nieuwe)
		})
	}
}

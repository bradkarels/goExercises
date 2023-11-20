package main

import (
	"reflect"
	"testing"
)

func Test_asteroidCollision(t *testing.T) {
	type args struct {
		asteroids []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test case -1",
			args: args{
				[]int{-42, 1, 2, 3, -3, 8, 9, -1, 5, -6, -7, -8},
			},
			want: []int{-42, 1, 2, 8, 9},
		},
		{
			name: "test case 0",
			args: args{
				[]int{5, 10, -5},
			},
			want: []int{5, 10},
		},
		{
			name: "test case 1",
			args: args{
				[]int{8, -8},
			},
			want: []int{},
		},
		{
			name: "test case 2",
			args: args{
				[]int{10, 2, -5},
			},
			want: []int{10},
		},
		{
			name: "test case 2",
			args: args{
				[]int{1, -2, -2, -2},
			},
			want: []int{-2, -2, -2},
		},
		{
			name: "test case 2",
			args: args{
				[]int{1, 1, -1, -2},
			},
			want: []int{-2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := asteroidCollision(tt.args.asteroids); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("asteroidCollision() = %v, want %v", got, tt.want)
			}
		})
	}
}

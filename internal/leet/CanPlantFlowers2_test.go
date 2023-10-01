package main

import "testing"

func Test_canPlaceFlowerz(t *testing.T) {
	type args struct {
		flowerbed []int
		n         int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test case 0",
			args: args{
				flowerbed: []int{1, 0, 0, 0, 1},
				n:         1,
			},
			want: true,
		},
		{
			name: "Test case 1",
			args: args{
				flowerbed: []int{1, 0, 0, 0, 1},
				n:         2,
			},
			want: false,
		},
		{
			name: "Test case 2",
			args: args{
				flowerbed: []int{0, 0, 1, 0, 1, 0},
				n:         1,
			},
			want: true,
		},
		{
			name: "Test case 3",
			args: args{
				flowerbed: []int{0, 0, 0, 1, 0, 1},
				n:         2,
			},
			want: false,
		},
		{
			name: "Test case 4",
			args: args{
				flowerbed: []int{0, 0, 0, 0, 0},
				n:         4,
			},
			want: false,
		},
		{
			name: "Test case 5",
			args: args{
				flowerbed: []int{1, 0, 0, 0, 1, 0, 0},
				n:         2,
			},
			want: true,
		},
		{
			name: "Test case 6",
			args: args{ // 0,1,0,0,0,0,0,1,0
				flowerbed: []int{1, 0, 0, 0, 0, 0, 1},
				n:         2,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canPlaceFlowerz(tt.args.flowerbed, tt.args.n); got != tt.want {
				t.Errorf("canPlaceFlowerz() = %v, want %v", got, tt.want)
			}
		})
	}
}

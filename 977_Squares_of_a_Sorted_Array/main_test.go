package main

import (
	"reflect"
	"testing"
)

func Test_sortedSquares(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{name: "[-4,-1,0,3,10] -> [0,1,9,16,100]", args: args{nums: []int{-4, -1, 0, 3, 10}}, want: []int{0, 1, 9, 16, 100}},
		{name: "[-7,-3,2,3,11] -> [4,9,9,49,121]", args: args{nums: []int{-7, -3, 2, 3, 11}}, want: []int{4, 9, 9, 49, 121}},
		{name: "[-7,-3,1,2,2,3,11] -> [1,4,4,9,9,49,121]", args: args{nums: []int{-7, -3, 1, 2, 2, 3, 11}}, want: []int{1, 4, 4, 9, 9, 49, 121}},
		{name: "[1] -> [1]", args: args{nums: []int{1}}, want: []int{1}},
		{name: "[] -> []", args: args{nums: []int{}}, want: []int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortedSquares(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortedSquares() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sortedSquares2(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{name: "[-4,-1,0,3,10] -> [0,1,9,16,100]", args: args{nums: []int{-4, -1, 0, 3, 10}}, want: []int{0, 1, 9, 16, 100}},
		{name: "[-7,-3,2,3,11] -> [4,9,9,49,121]", args: args{nums: []int{-7, -3, 2, 3, 11}}, want: []int{4, 9, 9, 49, 121}},
		{name: "[-7,-3,1,2,2,3,11] -> [1,4,4,9,9,49,121]", args: args{nums: []int{-7, -3, 1, 2, 2, 3, 11}}, want: []int{1, 4, 4, 9, 9, 49, 121}},
		{name: "[1] -> [1]", args: args{nums: []int{1}}, want: []int{1}},
		{name: "[] -> []", args: args{nums: []int{}}, want: []int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortedSquares2(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortedSquares2() = %v, want %v", got, tt.want)
			}
		})
	}
}

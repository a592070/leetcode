package main

import (
	"reflect"
	"testing"
)

func Test_threeSum(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
		{name: "1", args: args{nums: []int{-1, 0, 1, 2, -1, -4}}, want: [][]int{{-1, 0, 1}, {-1, -1, 2}}},
		{name: "2", args: args{nums: []int{0, 1, 1}}, want: [][]int{}},
		{name: "2", args: args{nums: []int{0, 0, 0}}, want: [][]int{{0, 0, 0}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := threeSum(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("threeSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

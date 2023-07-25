package main

import (
	"reflect"
	"testing"
)

func Test_insert(t *testing.T) {
	type args struct {
		intervals   [][]int
		newInterval []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
		{name: "", args: args{[][]int{{1, 3}, {6, 9}}, []int{2, 5}}, want: [][]int{{1, 5}, {6, 9}}},
		{name: "", args: args{[][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}, []int{4, 8}}, want: [][]int{{1, 2}, {3, 10}, {12, 16}}},
		{name: "", args: args{[][]int{}, []int{5, 7}}, want: [][]int{{5, 7}}},
		{name: "", args: args{[][]int{{1, 5}}, []int{2, 3}}, want: [][]int{{1, 5}}},
		{name: "", args: args{[][]int{{1, 5}, {7, 8}}, []int{2, 3}}, want: [][]int{{1, 5}, {7, 8}}},
		{name: "", args: args{[][]int{{1, 5}}, []int{6, 8}}, want: [][]int{{1, 5}, {6, 8}}},
		{name: "", args: args{[][]int{{2, 5}, {6, 7}, {8, 9}}, []int{0, 1}}, want: [][]int{{0, 1}, {2, 5}, {6, 7}, {8, 9}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := insert(tt.args.intervals, tt.args.newInterval); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("insert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_insert2(t *testing.T) {
	type args struct {
		intervals   [][]int
		newInterval []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
		{name: "", args: args{[][]int{{1, 3}, {6, 9}}, []int{2, 5}}, want: [][]int{{1, 5}, {6, 9}}},
		{name: "", args: args{[][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}, []int{4, 8}}, want: [][]int{{1, 2}, {3, 10}, {12, 16}}},
		{name: "", args: args{[][]int{}, []int{5, 7}}, want: [][]int{{5, 7}}},
		{name: "", args: args{[][]int{{1, 5}}, []int{2, 3}}, want: [][]int{{1, 5}}},
		{name: "", args: args{[][]int{{1, 5}, {7, 8}}, []int{2, 3}}, want: [][]int{{1, 5}, {7, 8}}},
		{name: "", args: args{[][]int{{1, 5}}, []int{6, 8}}, want: [][]int{{1, 5}, {6, 8}}},
		{name: "", args: args{[][]int{{2, 5}, {6, 7}, {8, 9}}, []int{0, 1}}, want: [][]int{{0, 1}, {2, 5}, {6, 7}, {8, 9}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := insert2(tt.args.intervals, tt.args.newInterval); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("insert2() = %v, want %v", got, tt.want)
			}
		})
	}
}

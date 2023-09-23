package main

import (
	"reflect"
	"testing"
)

func Test_compose(t *testing.T) {
	type args struct {
		n int
		k int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		// TODO: Add test cases.
		{name: "C(5, 0)", args: args{n: 5, k: 0}, want: 1},
		{name: "C(5, 1)", args: args{n: 5, k: 1}, want: (5 * 4 * 3 * 2 * 1) / (4 * 3 * 2 * 1) / (1)},
		{name: "C(5, 2)", args: args{n: 5, k: 2}, want: (5 * 4 * 3 * 2 * 1) / (3 * 2 * 1) / (1 * 2)},
		{name: "C(5, 3)", args: args{n: 5, k: 3}, want: (5 * 4 * 3 * 2 * 1) / (2 * 1) / (1 * 2 * 3)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compose(tt.args.n, tt.args.k); got != tt.want {
				t.Errorf("compose() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getRow(t *testing.T) {
	type args struct {
		rowIndex int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{name: "0", args: args{rowIndex: 0}, want: []int{1}},
		{name: "1", args: args{rowIndex: 1}, want: []int{1, 1}},
		{name: "3", args: args{rowIndex: 3}, want: []int{1, 3, 3, 1}},
		{name: "33", args: args{rowIndex: 33}, want: []int{1, 33, 528, 5456, 40920, 237336, 1107568, 4272048, 13884156, 38567100, 92561040, 193536720, 354817320, 573166440, 818809200, 1037158320, 1166803110, 1166803110, 1037158320, 818809200, 573166440, 354817320, 193536720, 92561040, 38567100, 13884156, 4272048, 1107568, 237336, 40920, 5456, 528, 33, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getRow(tt.args.rowIndex); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getRow() = %v, want %v", got, tt.want)
			}
		})
	}
}

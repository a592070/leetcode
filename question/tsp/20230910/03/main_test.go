package main

import "testing"

func Test_findRedundant(t *testing.T) {
	type args struct {
		planets []int32
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{name: "1", args: args{planets: []int32{2, 5, 3, 1}}, want: 1},
		{name: "2", args: args{planets: []int32{2, 3, 1, 4, 4, 10, 5, 5}}, want: 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findRedundant(tt.args.planets); got != tt.want {
				t.Errorf("findRedundant() = %v, want %v", got, tt.want)
			}
		})
	}
}

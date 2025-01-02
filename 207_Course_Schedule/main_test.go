package main

import "testing"

func Test_canFinish(t *testing.T) {
	type args struct {
		numCourses    int
		prerequisites [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "case 1", args: args{numCourses: 2, prerequisites: [][]int{{1, 0}}}, want: true},
		{name: "case 2", args: args{numCourses: 2, prerequisites: [][]int{{0, 1}, {1, 0}}}, want: false},
		{name: "case 2-1", args: args{numCourses: 3, prerequisites: [][]int{{1, 0}, {2, 1}, {0, 2}}}, want: false},
		{name: "case 3", args: args{numCourses: 4, prerequisites: [][]int{{1, 0}, {2, 1}, {3, 2}}}, want: true},
		{name: "case 4", args: args{numCourses: 4, prerequisites: [][]int{{3, 2}, {2, 1}, {1, 0}}}, want: true},
		{name: "case 5", args: args{numCourses: 20, prerequisites: [][]int{{0, 10}, {3, 18}, {5, 5}, {6, 11}, {11, 14}, {13, 1}, {15, 1}, {17, 4}}}, want: false},
		{name: "case 6", args: args{numCourses: 12, prerequisites: [][]int{{1, 0}, {2, 1}, {3, 2}, {4, 0}, {11, 10}, {12, 10}}}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canFinish(tt.args.numCourses, tt.args.prerequisites); got != tt.want {
				t.Errorf("canFinish() = %v, want %v", got, tt.want)
			}
		})
	}
}

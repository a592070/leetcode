package main

import "testing"

func Test_missingNumber(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"[3,0,1]", args{nums: []int{3, 0, 1}}, 2},
		{"[0,1]", args{nums: []int{0, 1}}, 2},
		{"[9,6,4,2,3,5,7,0,1]", args{nums: []int{9, 6, 4, 2, 3, 5, 7, 0, 1}}, 8},
		{"[0,1,2,3,4,5,6,8,9]", args{nums: []int{0, 1, 2, 3, 4, 5, 6, 8, 9}}, 7},
		{"[1,2,3,4,5,6,8,9,0]", args{nums: []int{0, 1, 2, 3, 4, 5, 6, 8, 9}}, 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := missingNumber(tt.args.nums); got != tt.want {
				t.Errorf("missingNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_missingNumber2(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"[3,0,1]", args{nums: []int{3, 0, 1}}, 2},
		{"[0,1]", args{nums: []int{0, 1}}, 2},
		{"[9,6,4,2,3,5,7,0,1]", args{nums: []int{9, 6, 4, 2, 3, 5, 7, 0, 1}}, 8},
		{"[0,1,2,3,4,5,6,8,9]", args{nums: []int{0, 1, 2, 3, 4, 5, 6, 8, 9}}, 7},
		{"[1,2,3,4,5,6,8,9,0]", args{nums: []int{0, 1, 2, 3, 4, 5, 6, 8, 9}}, 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := missingNumber2(tt.args.nums); got != tt.want {
				t.Errorf("missingNumber2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_missingNumber3(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name    string
		args    args
		wantSum int
	}{
		// TODO: Add test cases.
		{"[3,0,1]", args{nums: []int{3, 0, 1}}, 2},
		{"[0,1]", args{nums: []int{0, 1}}, 2},
		{"[9,6,4,2,3,5,7,0,1]", args{nums: []int{9, 6, 4, 2, 3, 5, 7, 0, 1}}, 8},
		{"[0,1,2,3,4,5,6,8,9]", args{nums: []int{0, 1, 2, 3, 4, 5, 6, 8, 9}}, 7},
		{"[1,2,3,4,5,6,8,9,0]", args{nums: []int{0, 1, 2, 3, 4, 5, 6, 8, 9}}, 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSum := missingNumber3(tt.args.nums); gotSum != tt.wantSum {
				t.Errorf("missingNumber3() = %v, want %v", gotSum, tt.wantSum)
			}
		})
	}
}

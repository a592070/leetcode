package main

import "testing"

func Test_isPalindrome(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{name: "121", args: args{121}, want: true},
		{name: "1221", args: args{1221}, want: true},
		{name: "122", args: args{122}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.x); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isPalindrome2(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{name: "121", args: args{121}, want: true},
		{name: "1221", args: args{1221}, want: true},
		{name: "122", args: args{122}, want: false},
		{name: "10", args: args{10}, want: false},
		{name: "23132", args: args{23132}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome2(tt.args.x); got != tt.want {
				t.Errorf("isPalindrome2() = %v, want %v", got, tt.want)
			}
		})
	}
}

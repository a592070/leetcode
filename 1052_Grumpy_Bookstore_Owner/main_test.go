package main

import "testing"

func Test_maxSatisfied3(t *testing.T) {
	type args struct {
		customers []int
		grumpy    []int
		minutes   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{customers: []int{1, 0, 1, 2, 1, 1, 7, 5}, grumpy: []int{0, 1, 0, 1, 0, 1, 0, 1}, minutes: 3},
			want: 16,
		},
		{
			name: "case 2",
			args: args{customers: []int{10, 4}, grumpy: []int{0, 1}, minutes: 2},
			want: 14,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSatisfied3(tt.args.customers, tt.args.grumpy, tt.args.minutes); got != tt.want {
				t.Errorf("maxSatisfied3() = %v, want %v", got, tt.want)
			}
		})
	}
}

package main

import "testing"

func Test_doesValidArrayExist(t *testing.T) {
	type args struct {
		derived []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Case 1",
			args: args{derived: []int{1, 1, 0}},
			want: true,
		},
		{
			name: "Case 2",
			args: args{derived: []int{1, 1}},
			want: true,
		},
		{
			name: "Case 3",
			args: args{derived: []int{1, 0}},
			want: false,
		},
		{
			name: "Case 4",
			args: args{derived: []int{1}},
			want: false,
		},
		{
			name: "Case 5",
			args: args{derived: []int{0}},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := doesValidArrayExist(tt.args.derived); got != tt.want {
				t.Errorf("doesValidArrayExist() = %v, want %v", got, tt.want)
			}
		})
	}
}

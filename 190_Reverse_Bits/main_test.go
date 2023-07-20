package main

import "testing"

func Test_reverseBits(t *testing.T) {
	type args struct {
		num uint32
	}
	tests := []struct {
		name string
		args args
		want uint32
	}{
		// TODO: Add test cases.
		{name: "00000010100101000001111010011100", args: args{43261596}, want: 964176192},
		{name: "11111111111111111111111111111101", args: args{4294967293}, want: 3221225471},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseBits(tt.args.num); got != tt.want {
				t.Errorf("reverseBits() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reverseBits2(t *testing.T) {
	type args struct {
		num uint32
	}
	tests := []struct {
		name string
		args args
		want uint32
	}{
		// TODO: Add test cases.
		{name: "00000010100101000001111010011100", args: args{43261596}, want: 964176192},
		{name: "11111111111111111111111111111101", args: args{4294967293}, want: 3221225471},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseBits2(tt.args.num); got != tt.want {
				t.Errorf("reverseBits2() = %v, want %v", got, tt.want)
			}
		})
	}
}

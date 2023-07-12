package main

import "testing"

func Test_isSymmetric(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"case nil", args{nil}, true},
		{"case [0]", args{&TreeNode{}}, true},
		{"case [0,1,2]", args{&TreeNode{Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 2}}}, false},
		{"case [1,2,2,3,4,4,3]", args{&TreeNode{Val: 1,
			Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 4}},
			Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 3}}},
		},
			true},
		{"case [1,2,2,nil,3,nil,3]", args{&TreeNode{Val: 1,
			Left:  &TreeNode{Val: 2, Left: nil, Right: &TreeNode{Val: 3}},
			Right: &TreeNode{Val: 2, Left: nil, Right: &TreeNode{Val: 3}}},
		},
			true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSymmetric(tt.args.root); got != tt.want {
				t.Errorf("isSymmetric() = %v, want %v", got, tt.want)
			}
		})
	}
}

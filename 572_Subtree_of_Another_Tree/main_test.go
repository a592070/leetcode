package main

import "testing"

func Test_isSubtree(t *testing.T) {
	type args struct {
		root    *TreeNode
		subRoot *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		//root = [3,4,5,1,2], subRoot = [4,1,2]
		{name: "case1", args: args{
			root: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val:   4,
					Left:  &TreeNode{Val: 1},
					Right: &TreeNode{Val: 2},
				},
				Right: &TreeNode{Val: 5},
			},
			subRoot: &TreeNode{
				Val:   4,
				Left:  &TreeNode{Val: 1},
				Right: &TreeNode{Val: 2},
			},
		}, want: true},

		//root = [3,4,5,1,2,null,null,null,null,0], subRoot = [4,1,2]
		{name: "case1", args: args{
			root: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val:   4,
					Left:  &TreeNode{Val: 1},
					Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 0}},
				},
				Right: &TreeNode{Val: 5},
			},
			subRoot: &TreeNode{
				Val:   4,
				Left:  &TreeNode{Val: 1},
				Right: &TreeNode{Val: 2},
			},
		}, want: false},

		//root = [1,1], subRoot = [1]
		{name: "case1", args: args{
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 1,
				},
			},
			subRoot: &TreeNode{
				Val: 1,
			},
		}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSubtree(tt.args.root, tt.args.subRoot); got != tt.want {
				t.Errorf("isSubtree() = %v, want %v", got, tt.want)
			}
		})
	}
}

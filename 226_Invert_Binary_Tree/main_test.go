package main

import (
	"reflect"
	"testing"
)

func Test_invertTree(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		// TODO: Add test cases.
		{"case nil", args{nil}, nil},
		{"case [0]", args{&TreeNode{}}, &TreeNode{}},
		{"case [0,1,2]", args{&TreeNode{Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 2}}}, &TreeNode{Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 1}}},
		{"case [4,2,7,1,3,6,9], except [4,7,2,9,6,3,1]", args{&TreeNode{Val: 4,
			Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}},
			Right: &TreeNode{Val: 7, Left: &TreeNode{Val: 6}, Right: &TreeNode{Val: 9}}},
		},
			&TreeNode{Val: 4,
				Left:  &TreeNode{Val: 7, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 6}},
				Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 1}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := invertTree(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("invertTree() = %v, want %v", got, tt.want)
			}
		})
	}
}

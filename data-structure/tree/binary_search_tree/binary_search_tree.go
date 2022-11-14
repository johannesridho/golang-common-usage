package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// https://leetcode.com/problems/insert-into-a-binary-search-tree/
// binary search the val in tree until it reach a nil node
// insert the new node with val there
// O(log N) time, O(1) space
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{
			Val: val,
		}
	}

	node := root

	var prev *TreeNode
	for node != nil {
		prev = node
		if val > node.Val {
			node = node.Right
		} else {
			node = node.Left
		}
	}

	newNode := &TreeNode{
		Val: val,
	}
	if val > prev.Val {
		prev.Right = newNode
	} else {
		prev.Left = newNode
	}

	return root
}

package main

import (
	"container/list"
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5}}
	root.Right = &TreeNode{Val: 3, Left: &TreeNode{Val: 6}, Right: &TreeNode{Val: 7}}

	fmt.Println("traverse by level with bfs")
	for _, valuesByLevel := range traverseByLevelWithBFS(root) {
		fmt.Println(valuesByLevel)
	}

	fmt.Println("traverse by level with dfs")
	for _, valuesByLevel := range traverseByLevelWithDFS(root) {
		fmt.Println(valuesByLevel)
	}
}

func traverseByLevelWithBFS(root *TreeNode) [][]int {
	queue := list.New()
	queue.PushBack(root)

	var output [][]int

	for queue.Len() != 0 {
		curLen := queue.Len()
		var curRowElements []int

		for i := 0; i < curLen; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			curRowElements = append(curRowElements, node.Val)

			if node.Left != nil {
				queue.PushBack(node.Left)
			}

			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}

		output = append(output, curRowElements)
	}

	return output
}

func traverseByLevelWithDFS(root *TreeNode) [][]int {
	var output [][]int
	dfs(root, 1, &output)
	return output
}

func dfs(node *TreeNode, depth int, output *[][]int) {
	if node == nil {
		return
	}

	// new depth found, need to create the array output[newDepth]
	if len(*output) < depth {
		*output = append(*output, []int{node.Val})
	} else { // depth is already there, append to the existing array (it's zero indexed, so index is depth-1)
		(*output)[depth-1] = append((*output)[depth-1], node.Val)
	}

	depth++
	dfs(node.Left, depth, output)
	dfs(node.Right, depth, output)
}

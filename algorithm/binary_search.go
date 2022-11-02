package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("actual:", binarySearchTemplate([]int{1, 3, 4, 6, 7, 9, 10}, 1), "expected:", 0)
	fmt.Println("actual:", binarySearchTemplate([]int{1, 3, 4, 6, 7, 9, 10}, 4), "expected:", 2)
	fmt.Println("actual:", binarySearchTemplate([]int{1, 3, 4, 6, 7, 9, 10}, 7), "expected:", 4)
	fmt.Println("actual:", binarySearchTemplate([]int{1, 3, 4, 6, 7, 9, 10}, 9), "expected:", 5)
	fmt.Println("actual:", binarySearchTemplate([]int{1, 3, 4, 6, 7, 9, 10}, 10), "expected:", 6)

	fmt.Println("using golang helper function:")
	arr := []int{1, 3, 4, 6, 7, 9, 10}
	result := sort.Search(len(arr)-1, func(i int) bool {
		return isSolutionInTheLeft(arr, i, 7)
	})
	fmt.Println("actual:", result, "expected:", 4)
}

// binary search template
// based on sort.Search and https://leetcode.com/discuss/general-discussion/786126/Python-Powerful-Ultimate-Binary-Search-Template.-Solved-many-problems
// time O(log n), space O(1)
func binarySearchTemplate(arr []int, target int) int {
	// initial value of left = minimum possible solution, right = maximum possible solution
	left, right := 0, len(arr)-1

	for left < right {
		// mid := (left + right) / 2
		// use mid := left + (right-left)/2 in case left + right could overflow,
		mid := left + (right-left)/2

		if isSolutionInTheLeft(arr, mid, target) {
			right = mid // preserves f(right) == true
		} else {
			left = mid + 1 // preserves f(left-1) == false
		}
	}

	// return either left or left-1
	// after exiting the while loop, left is the minimum index satisfying the isSolutionInTheLeft function
	// return left-1 if we want the index before the minimum satisfying index
	return left
}

func isSolutionInTheLeft(arr []int, idx, target int) bool {
	if target <= arr[idx] {
		return true
	}
	return false
}

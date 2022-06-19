//https://go.dev/tour/concurrency/8
//traverse the tree concurrently and share information using channel

package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}

	Walk(t.Left, ch)
	ch <- t.Value
	Walk(t.Right, ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)

	go func() {
		Walk(t1, ch1)
		close(ch1)
	}()

	go func() {
		Walk(t2, ch2)
		close(ch2)
	}()

	var v1, v2 int
	var ok bool
	for {
		select {
		case v1, ok = <-ch1:
			if !ok {
				return true
			}
			v2, ok = <-ch2
			if !ok {
				return false
			}

			if v1 != v2 {
				return false
			}

		case v2, ok = <-ch2:
			if !ok {
				return true
			}

			v1, ok = <-ch1
			if !ok {
				return false
			}

			if v1 != v2 {
				return false
			}
		}
	}

	return true
}

func main() {
	t1 := tree.New(1)
	t2 := tree.New(2)
	t2b := tree.New(2)

	fmt.Println("actual", Same(t1, t2), "expected", false)
	fmt.Println("actual", Same(t2, t2b), "expected", true)
	fmt.Println("actual", Same(t1, t2b), "expected", false)
	fmt.Println("actual", Same(tree.New(11), tree.New(11)), "expected", true)
	fmt.Println("actual", Same(tree.New(1), tree.New(11)), "expected", false)
	fmt.Println("actual", Same(tree.New(10), tree.New(11)), "expected", false)
}

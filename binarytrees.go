package main

import (
	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		Walk(t.Left, ch)
	}
	// println("min = ", t.Value)
	ch <- t.Value
	if t.Right != nil {
		Walk(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)
	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for i := 0; i < 10; i++ {
		v1, v2 := <-ch1, <-ch2
		// println(v1, " ", v2)
		if v1 != v2 {
			println("Not equal")
			return false
		}
	}

	println("Equal")
	return true
}

func main() {
	Same(tree.New(1), tree.New(1))
	Same(tree.New(1), tree.New(2))
	// ch := make(chan int, 10)
	// go Walk(tree.New(1), ch)
	// for i := 0; i < 10; i++ {
	// 	println(<-ch)
	// }
	//	println(Same(tree.New(1), tree.New(1)))
	//	println(Same(tree.New(1), tree.New(2)))
}

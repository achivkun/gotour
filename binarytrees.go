package main

import (
	"errors"
	"fmt"

	"golang.org/x/tour/tree"
)

func FindRemMin(t *tree.Tree, parent *tree.Tree) (int, error, *tree.Tree, *tree.Tree) {
	fmt.Println(t)
	var retVal int = 0
	var err error = nil
	var retParent *tree.Tree = nil
	var retRoot *tree.Tree = nil

	if t == nil {
		println("Error")
		err = errors.New("empty tree")
		return retVal, err, retRoot, retParent
	}

	if t.Left == nil {
		retVal = t.Value

		if parent != nil {
			parent.Left = nil
		}

		if t.Right != nil {
			if parent != nil {
				parent.Left = t.Right
			}
		}

		retRoot = parent
		return retVal, err, retRoot, retParent
	} else {
		return FindRemMin(t.Left, t)
	}
}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	println("start")

	var minVal int
	var err error
	var parent *tree.Tree = nil
	for {
		minVal, err, t, parent = FindRemMin(t, parent)
		println("min val =", minVal)
		if err != nil {
			break
		}
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch := make(chan int)
	go Walk(t1, ch)
	go Walk(t2, ch)
	v1, v2 := <-ch, <-ch

	if v1 != v2 {
		return true
	}

	return false
}

func main() {
	ch := make(chan int, 10)
	println("before")
	Walk(tree.New(1), ch)
	println("after")
	for i := 0; i < 10; i++ {
		ch <- i
	}
	//	println(Same(tree.New(1), tree.New(1)))
	//	println(Same(tree.New(1), tree.New(2)))
}

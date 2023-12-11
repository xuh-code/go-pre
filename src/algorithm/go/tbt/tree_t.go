package main

import (
	"fmt"
	"github.com/emirpasic/gods/trees/redblacktree"
)

func main() {
	tree := redblacktree.NewWithIntComparator()

	tree.Put(1, 1)
	s := tree.String()
	fmt.Printf(s)
	tree.Put(2, 1)
	s = tree.String()
	fmt.Printf(s)
	tree.Put(3, 1)
	s = tree.String()
	fmt.Printf(s)
	tree.Put(4, 1)
	s = tree.String()
	fmt.Printf(s)
	tree.Put(5, 1)
	s = tree.String()
	fmt.Printf(s)
	tree.Put(6, 1)
	s = tree.String()
	fmt.Printf(s)
	tree.Put(7, 1)
	s = tree.String()
	fmt.Printf(s)
	tree.Put(8, 1)
	s = tree.String()
	fmt.Printf(s)
	tree.Put(9, 1)
	s = tree.String()
	fmt.Printf(s)
	//tree.Put(8, 1)
	//tree.Put(9, 1)

	s = tree.String()
	fmt.Printf(s)
}

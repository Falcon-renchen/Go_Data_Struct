package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Tree struct {
	Left *Tree
	Value int
	Right *Tree
}

func traverse(t *Tree)  {
	if t == nil {
		return
	}
	traverse(t.Left)
	fmt.Println(t.Value, " ")
	traverse(t.Right)
}

func Create(n int) *Tree {
	var t *Tree
	rand.Seed(time.Now().Unix())
	for i:=0; i<2*n; i++ {
		temp := rand.Intn(n*2)
		t = insert(t,temp)
	}
	return t
}

func insert(t *Tree, v int) *Tree {
	if t == nil {
		return &Tree{nil, v, nil}
	}

	if v == t.Value {
		return t
	}

	if v < t.Value {
		t.Left = insert(t.Left, v)
		return t
	}
	t.Right = insert(t.Right, v)
	return t
}

func main() {
	tree := Create(10)
	fmt.Println("The value of the root of the tree is", tree.Value)
	traverse(tree)
	fmt.Println()
	tree = insert(tree, -10)
	tree = insert(tree, -2)
	traverse(tree)
	fmt.Println()
	fmt.Println("The value of the root of the tree is", tree.Value)
}
/*
Output:

The value of the root of the tree is 0
0
1
2
4
5
8
9
11
14
15
17
18
19

-10
-2
0
1
2
4
5
8
9
11
14
15
17
18
19

The value of the root of the tree is 0

 */
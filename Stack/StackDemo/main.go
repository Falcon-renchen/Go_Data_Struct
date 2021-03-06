package main

import (
	"fmt"
	stack "github.com/golang-collections/collections/stack"
)
func main() {
	// Create a new stack
	fmt.Println("Creating New SingleStack")
	exstack := stack.New()
	fmt.Println("Pushing 1 to stack")
	exstack.Push(1) // push 1 to stack
	fmt.Println("Top of SingleStack is : ", exstack.Peek())
	fmt.Println("Popping 1 from stack")
	exstack.Pop() // remove 1 from stack
	fmt.Println("SingleStack length is : ", exstack.Len())
}

/*
Output:
Creating New SingleStack
Pushing 1 to stack
Top of SingleStack is :  1
Popping 1 from stack
SingleStack length is :  0
 */
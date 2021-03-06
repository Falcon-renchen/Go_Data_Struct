package main

import (
	"errors"
	"fmt"
)

type Stack struct {
	MaxTop int
	Top int
	arr [5]int
}

func (s *Stack) Push(val int) (err error) {
	if s.Top == s.MaxTop-1 {
		fmt.Println("栈满")
		return
	}
	s.Top++
	s.arr[s.Top] = val
	return 
}

func (s *Stack) Pop() (val int, err error) {
	if s.Top == -1 {
		fmt.Println("栈空")
		return 0, errors.New("empty stack")
	}
	val = s.arr[s.Top]
	s.Top--
	return val,nil
}

func (s *Stack) List()  {
	if s.Top == -1 {
		fmt.Println("栈空")
		return
	}
	for i:=s.Top; i>=0; i-- {
		fmt.Printf("arr[%d]=%d\n", i, s.arr[i])
	}
}

func main() {
	stack := &Stack{
		MaxTop: 5,
		Top:    -1,
	}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)
	stack.List()
	fmt.Println()
	pop, _ := stack.Pop()
	fmt.Println(pop)
	stack.List()
}

/*
Output:
arr[4]=5
arr[3]=4
arr[2]=3
arr[1]=2
arr[0]=1

5
arr[3]=4
arr[2]=3
arr[1]=2
arr[0]=1
 */
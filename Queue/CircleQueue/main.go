package main

import (
	"errors"
	"fmt"
)

type CircleQueue struct {
	maxSize int // 5
	array [5]int
	tail int	//指向队尾 0
	head int	//指向队首 0
}

func (c *CircleQueue) Push(val int) (err error) {
	if c.IsFull() {
		return errors.New("queue full")
	}
	c.array[c.tail] = val

	c.tail = (c.tail+1)%c.maxSize
	return
}

func (c *CircleQueue) Pop() (val int, err error) {
	if c.IsEmpty() {
		return 0, errors.New("queue empty.")
	}
	val = c.array[c.head]

	c.head = (c.head+1)%c.maxSize
	return 
}

//显示队列
func (c *CircleQueue) ListQueue() {
	fmt.Println("环形队列如下：")
	size := c.Size()
	if size==0 {
		errors.New("queue is empty!")
	}
	tempHead := c.head
	for i:=0; i<size; i++ {
		fmt.Printf("arr[%d]=%d",tempHead, c.array[tempHead])
		tempHead = (tempHead+1) % c.maxSize
	}
	fmt.Println()


}

func (c *CircleQueue) IsFull() bool {
	return (c.tail+1) % c.maxSize == c.head
}

func (c *CircleQueue) IsEmpty() bool {
	return c.tail == c.head
}


//取出环形链表有多少个元素
func (c *CircleQueue) Size() int {
	return (c.tail + c.maxSize - c.head) % c.maxSize
}

func main() {
	
}

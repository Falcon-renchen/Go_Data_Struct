package main

import (
	"errors"
	"fmt"
)

type Queue struct {
	MaxSize int
	array [5]int
	front int
	rear int
}

func (q *Queue) ShowQueue() {
	for i:=q.front+1; i<q.rear; i++ {
		fmt.Printf("array[%d]=%d",i,q.array[i])
	}
}

func (q *Queue) AddQueue(val int) (err error) {
	if q.rear == q.MaxSize-1 {
		return errors.New("queue full")
	}
	q.rear ++
	q.array[q.rear] = val
	return
}

func (q *Queue) GetQueue() (val int, err error) {
	if q.rear == q.front {
		return -1,errors.New("queue empty")
	}
	q.front++
	val = q.array[q.front]
	return val, err
}

func (q *Queue) DelQueue() (err error) {
	if q.rear == q.front {
		return errors.New("queue empty!")
	}
	q.rear--
	return
}

func (q *Queue) UpdateQueue(j int, val int) (err error) {
	if q.rear == q.front {
		return errors.New("queue empty!")
	}
	for i:=q.front; i<q.rear; i++ {
		if j == q.front {
			q.array[q.front] = val
		}
	}
	return
}


func main() {
	q := &Queue{
		MaxSize: 10,
		array:   [5]int{},
		front:   -1,
		rear:    -1,
	}
	fmt.Println(q)
}

package main

import (
	"container/list"
	"fmt"
)
//双向链表
func main() {
	ll := list.New()
	three :=  ll.PushBack(3)	// [3]
	four := ll.InsertBefore(4,three) 	// [4,3]
	ll.InsertBefore(2,three)	//  	[4,2,3]
	ll.InsertAfter(5,three)	//		[4,2,3,5]
	ll.MoveToBack(four)	//		[2,3,5,4]
	ll.PushFront(1)	// [1,2,3,5,4]
	ll.PushBack(5)//		[1,2,3,5,4,5]
	ll.Remove(three)	//移除3

	listLenght := ll.Len()
	fmt.Println(listLenght)
	for e := ll.Front(); e != nil; e = e.Next() {
		fmt.Printf("%d ",e.Value)
	}
}
/*
Output:
5
1 2 5 4 5
*/
package main

import "fmt"

type SingleLinkedList struct {
	head *LinkedListNode
}

type LinkedListNode struct {
	data string
	next *LinkedListNode
}

func (ll *SingleLinkedList) Append(node *LinkedListNode) {
	if ll.head == nil {
		ll.head = node
		return
	}
	
	currentNode := ll.head
	for currentNode.next != nil {
		currentNode = currentNode.next
	}
	currentNode.next = node
}

func main() {
	ll := &SingleLinkedList{}
	ll.Append(&LinkedListNode{data: "hello"})
	ll.Append(&LinkedListNode{data: "high"})
	ll.Append(&LinkedListNode{data: "performance"})
	ll.Append(&LinkedListNode{data: "go"})

	for e := ll.head; e != nil; e = e.next {
		fmt.Println(e.data)
	}

}

/*
Output:
hello
high
performance
go
*/
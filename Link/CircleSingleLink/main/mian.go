package main

import (
	"fmt"
)

type CatNode struct {
	no int
	name string
	next *CatNode
}

func InsertNode(head *CatNode, newCatNode *CatNode)  {
	//第一只猫
	if head.next == nil {
		head.no = newCatNode.no
		head.name = newCatNode.name
		head.next = head
		return
	}

	temp := head
	for {
		if temp.next == head {
			break
		}
		temp = temp.next
	}
	temp.next = newCatNode
	newCatNode.next = head
}

func ListCatNode(head *CatNode)  {
	temp := head

		if temp.next == nil {
			fmt.Println("空空如也")
			return
		}

	for {
		fmt.Printf("猫的信息为[id=%d, name=%s]->\n",temp.no,temp.name)
		if temp.next == head {
			break
		}
		temp = temp.next
	}
}

func DelCatNode(head *CatNode, id int) *CatNode {
	temp := head
	helper := head

	if temp.next == nil {
		fmt.Println("空空如也")
		return head
	}
	if temp.next == head {
		temp.next = nil
		return head
	}
	//将链表定位到最后
	for {
		if helper.next == head {
			break
		}
		helper = helper.next
	}

	flag := true
	for {
		if temp.no == id {
			if temp == head {
				head = head.next
			}
		}
		if temp.next == head {	//说明找到最后一个， 即没有找到
			break
		}
		if temp.no == id {	//找到了，可以直接删除
			helper.next = temp.next
			flag = false
			break
		}
		temp = temp.next  	//移动
		helper = helper.next //价值移动
	}
	if flag {	//如果为真，说明上面没有删除
		helper.next = temp.next
	} else {
		fmt.Println("没有这只猫")
	}
	return head
}


func main() {
	head := &CatNode{}

	cat1 := &CatNode{
		no:   1,
		name: "tom",
	}
	cat2 := &CatNode{
		no:   2,
		name: "tom2",
	}
	cat3 := &CatNode{
		no:   3,
		name: "tom3",
	}

	InsertNode(head, cat1)
	InsertNode(head, cat2)
	InsertNode(head, cat3)
	ListCatNode(head)
	fmt.Println()
	head = DelCatNode(head,1)
	ListCatNode(head)

}
/*
Output:
猫的信息为[id=1, name=tom]->
猫的信息为[id=2, name=tom2]->
猫的信息为[id=3, name=tom3]->

猫的信息为[id=2, name=tom2]->
猫的信息为[id=3, name=tom3]->

 */
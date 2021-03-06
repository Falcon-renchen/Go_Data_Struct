package main

import "fmt"

type HeroNode struct {
	no int
	name string
	nickName string
	pre *HeroNode
	next *HeroNode
}


func InsertHeroNode(head *HeroNode, newHeroNode *HeroNode)  {
	temp := head
	for {
		if temp.next == nil {
			break
		}
		temp = temp.next
	}
	temp.next = newHeroNode
	newHeroNode.pre = temp
}

func InsertHeroNode2(head *HeroNode, newHeroNode *HeroNode)  {
	temp := head
	flag := true
	for {
		if temp.next == nil {
			break
		} else if temp.next.no > newHeroNode.no {
			break
		} else if temp.next.no == newHeroNode.no {
			flag = false
			break
		}
		temp = temp.next
	}
	if !flag {
		fmt.Println("对不起，已经存在no=", newHeroNode.no)
		return
	} else {
		newHeroNode.next = temp.next
		newHeroNode.pre = temp
		temp.next.pre = newHeroNode
		temp.next = newHeroNode
	}
}

func DelHeroNode(head *HeroNode, id int)  {
	temp := head
	flag := false
	for {
		if temp.next == nil {
			break
		} else if temp.next.no == id {
			flag = true
			break
		}
		temp = temp.next
	}
	if flag {
		if temp.next != nil {
			temp.next = temp.next.next
			temp.next.pre = temp
		}

	} else {
		fmt.Println("要删除的id不存在")
	}
}

func ListHeroNode(head *HeroNode)  {
	temp := head

	if temp.next == nil {
		fmt.Println("空空如也")
		return
	}
	for {
		fmt.Printf("[%d, %s, %s]==>", temp.next.no,temp.next.name,temp.next.nickName)
		temp  = temp.next
		if temp.next==nil {
			break
		}
	}
}

func ListHeroNode2(head *HeroNode)  {
	temp := head

	//判断链表是不是一个空链表
	if temp.next == nil {
		fmt.Println("空空如也")
		return
	}
	//让temp定位到双向链表的最后节点
	for {
		if temp.next==nil {
			break
		}
		temp = temp.next
	}

	for {
		fmt.Printf("[%d, %s, %s]==>", temp.no,temp.name,temp.nickName)
		temp  = temp.pre
		if temp.pre==nil {
			break
		}
	}

}

func main() {
	head := &HeroNode{}
	hero1 := &HeroNode{
		no:       1,
		name:     "宋江",
		nickName: "及时雨",
	}
	hero2 := &HeroNode{
		no:       2,
		name:     "李奎",
		nickName: "黑旋风",
	}
	hero3 := &HeroNode{
		no:       3,
		name:     "鲁智深",
		nickName: "花和尚",
	}

	InsertHeroNode(head,hero1)
	InsertHeroNode(head,hero2)
	InsertHeroNode(head,hero3)
	ListHeroNode(head)
	fmt.Println()
	//DelHeroNode(head,2)
	ListHeroNode2(head)
}

/*
Output:
[1, 宋江, 及时雨]==>[2, 李奎, 黑旋风]==>[3, 鲁智深, 花和尚]==>
[3, 鲁智深, 花和尚]==>[2, 李奎, 黑旋风]==>[1, 宋江, 及时雨]==>
 */
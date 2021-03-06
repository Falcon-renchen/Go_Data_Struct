package main

import (
	"fmt"
)

type Emp struct {
	Id int
	Name string
	Next *Emp
}

//第一个结点存放雇员,,不带表头
type EmpLink struct {
	Head *Emp
}

//定义hashtable，含有一个链表数组
type HashTable struct {
	LinkArr [7]EmpLink
}

//显示链表信息
func (e *EmpLink) ShowLink(no int) {
	if e.Head == nil {
		fmt.Printf("链表%d为空\n",no)
		return
	}
	//变量当前的链表，并显示数据
	cur := e.Head	//辅助指针
	for {
		if cur != nil {
			fmt.Printf("链表%d, 雇员id=%d,名字=%s ->",no,cur.Id,cur.Name)
			cur = cur.Next
		} else {
			break
		}
	}
	fmt.Println()

}

func (e *EmpLink) Insert(emp *Emp) {
	cur := e.Head //这是一个辅助指针
	var pre *Emp = nil//这是一个辅助指针，pre在cur之前
	//如果当前的EmpLink就是一个空链表
	if cur == nil {
		e.Head = emp	//完成
		return
	}
	//如果不是一个空链表，给emp找到对应的位置并插入
	//思路是让cur和emp比较，然后让pre始终保持在cur前面
	for {
		if cur!=nil {
			if cur.Id >= emp.Id {
				//找到位置
				break
			}
			pre = cur // 同步
			cur = cur.Next
		} else {
			break
		}
	}

	//推出时，我们将看一下emp是否添加到链表最后
	pre.Next = emp
	emp.Next = cur
}

//根据id查找对应的雇员，如果没有就返回nil
func (e *EmpLink) FindById(id int) *Emp {
	cur := e.Head
	for {
		if cur != nil && cur.Id == id {
			return cur
		} else if cur == nil {
			break
		}
		cur = cur.Next
	}
	return nil
}

func (e *Emp) ShowMe() {
	fmt.Printf("链表%d 找到该雇员%d\n", e.Id % 7,e.Id)
}

//给hashtable添加insert雇员方法
func (h *HashTable) Insert(emp *Emp) {
	//使用散列函数 ，确定将该雇员添加到哪个链表
	linkNo := h.HashFunc(emp.Id)
	//使用对应的链表添加
	h.LinkArr[linkNo].Insert(emp)
}

//编写一个散列方法
func (h *HashTable) HashFunc(id int) int {
	return id % 7 //得到了一个值，就是对于链表的下标
}

//显示所有雇员信息
func (h *HashTable) ShowAll() {
	for i:=0; i<len(h.LinkArr); i++ {
		h.LinkArr[i].ShowLink(i)
	}
}


func main() {

}
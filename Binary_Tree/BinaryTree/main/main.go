package main

import "fmt"

type HeroName struct {
	No int
	Name string
	Left *HeroName
	Right *HeroName
}

//前序遍历
func PreOrder(node *HeroName)  {
	if node!=nil {
		fmt.Printf("no=%d, name=%s\n", node.No,node.Name)
		PreOrder(node.Left)
		PreOrder(node.Right)
	}
}
//中序遍历
func InfixOrder(node *HeroName)  {
	if node!=nil {
		InfixOrder(node.Left)
		fmt.Printf("no=%d, name=%s\n", node.No,node.Name)
		InfixOrder(node.Right)
	}
}

//后序遍历
func PostOrder(node *HeroName)  {
	if node!=nil {
		InfixOrder(node.Left)
		InfixOrder(node.Right)
		fmt.Printf("no=%d, name=%s\n", node.No,node.Name)
	}
}


func main() {
	root := &HeroName{
		No:       1,
		Name:     "宋江",
	}
	left1 := &HeroName{
		No:       2,
		Name:     "李奎",
	}
	right1 := &HeroName{
		No:       3,
		Name:     "吴用",
	}
	root.Left = left1
	root.Right = right1

	left2 := &HeroName{
		No:    4,
		Name:  "hhh",
	}

	right2 :=  &HeroName{
		No:    5,
		Name:  "卢俊义",
	}
	right1.Right = right2
	right1.Left = left2

	//PreOrder(root)
	//InfixOrder(root)
	PostOrder(root)
}

/*
Output:
no=2, name=李奎
no=4, name=hhh
no=3, name=吴用
no=5, name=卢俊义
no=1, name=宋江

 */
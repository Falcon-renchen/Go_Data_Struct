package main

import "fmt"

type Boy struct {
	No int
	Next *Boy
}

func AddBoy(num int) *Boy {
	first := &Boy{}
	curBoy := &Boy{}

	if num < 1 {
		fmt.Println("num的值不对")
		return first
	}
	//循环的构建环形链表
	for i:=1; i<=num; i++ {
		boy := &Boy{
			No:   i,
		}
		//第一个孩子比较特殊
		if i==1 {
			first = boy
			curBoy = boy
			curBoy.Next = first
		} else {
			curBoy.Next = boy
			curBoy = boy
			curBoy.Next  = first
		}
	}
	return first
}

func ShowBoy(first *Boy)  {
	if first.Next == nil {
		fmt.Println("链表为空")
		return
	}
	//创建一个指针帮助遍历
	curBoy := first
	for {
		fmt.Printf("小孩的编号=%d ->",curBoy.No)
		if curBoy.Next == first {
			break
		}
		curBoy = curBoy.Next
	}
}

//五个人围成一圈， 从第2个人数数， 喊到第m个数字的 推出队列。
func PlayGames(first *Boy, startNo int, countNum int)  {
	if first.Next == nil {
		fmt.Println("空链表")
		return
	}
	//留一个判断startNo 《 小孩 总数
	//定义辅助指针
	tail := first
	//让tail执行环形链表的最后一个小孩，
	for {
		if tail.Next == first {
			break
		}
		tail = tail.Next
	}
	//让first移动到startNo
	for i:=1; i <= startNo-1; i++ {
		first = first.Next
		tail = tail.Next
	}
	fmt.Println()
	//刚开始数countNum， 就删除first指向的小孩
	for {
		//开始数countNum-1
		for i:=1; i<=countNum-1; i++ {
			first = first.Next
			tail = tail.Next
		}
		fmt.Printf("最后出圈的小孩编号为:%d ->", first.No)
		//删除first执行的小孩
		first = first.Next
		tail.Next = first
		if tail == first {
			 break
		}
	}
	fmt.Printf("最后出圈的小孩编号为:%d ->", first.No)

}

func main() {
	first := AddBoy(5)
	ShowBoy(first)
	PlayGames(first,2,3)
}

/*
Output:
小孩的编号=1 ->小孩的编号=2 ->小孩的编号=3 ->小孩的编号=4 ->小孩的编号=5 ->
最后出圈的小孩编号为:4 ->最后出圈的小孩编号为:2 ->最后出圈的小孩编号为:1 ->最后出圈的小孩编号为:3 ->最后出圈的小孩编号为:5 ->
 */
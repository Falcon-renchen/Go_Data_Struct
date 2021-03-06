package main

import (
	"errors"
	"fmt"
	"strconv"
)

type Stack struct {
	MaxTop int
	Top int
	arr [20]int
}

func (s *Stack) Push(val int) (err error) {
	if s.Top == s.MaxTop-1 {
		fmt.Println("栈满")
		return errors.New("stack full")
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

func (s *Stack) IsOper(val int) bool {
	if val==42 || val==43 || val==45 || val==47  {
		return true
	} else {
		return false
	}
}

func (s *Stack) Cal(num1 int, num2 int, oper int) int {
	res := 0
	switch oper {
	case 42:
		res = num2 * num1
	case 43:
		res = num2 + num1
	case 45:
		res = num2 - num1
	case 47:
		res = num2 / num1
	default:
		fmt.Println("运算符错误")
	}
	return res
}
//编写  一个方法，返回某个运算符的优先级
func (s *Stack) Priority(oper int) int {
	res := 0
	if oper==42 || oper==47 {
		res = 1
	} else if oper==43 || oper== 45 {
		res = 0
	}
	return res
}


func main() {
	numStack := &Stack{
		MaxTop: 20,
		Top:    -1,
	}
	operStack := &Stack{
		MaxTop: 20,
		Top:    -1,
	}
	exp := "30+2*6-2"

	index := 0
	num1 := 0
	num2 := 0
	oper := 0
	result := 0
	keepNum:=""
	for {
		ch := exp[index:index+1]
		temp := int([]byte(ch)[0])	//字符对应的ASCII码
		if operStack.IsOper(temp) {
			if operStack.Top == -1 {
				operStack.Push(temp)
			} else {
				//如果发现栈顶符号优先级大于当前准备入栈的符号，则pop出一个符号和两个数字
				if operStack.Priority(operStack.arr[operStack.Top]) >= operStack.Priority(temp) {
					num1, _ = numStack.Pop()
					num2, _ = numStack.Pop()
					oper, _ = operStack.Pop()
					result = operStack.Cal(num1,num2,oper)
					numStack.Push(result)	//将计算结果重新压入栈中
					operStack.Push(temp)	//当前符号压入栈中
				} else {
					operStack.Push(temp)  	//如果栈顶优先级不大于准备入栈的符号优先级，则直接入栈
				}
			}
		}else {

			keepNum +=  ch
			//数字
			if index==len(exp)-1 {
				//如果已经到了表达式 最后，直接压入计算
				val,_  := strconv.ParseInt(keepNum,10,64)
				numStack.Push(int(val))
			} else {

				//向后看是不是一个操作符，
				if operStack.IsOper(int([]byte(exp[index+1:index+2])[0])) {
					val,_ := strconv.ParseInt(keepNum,10,64)
					numStack.Push(int(val))
					keepNum=""
				}
			}


			//val,_ := strconv.ParseInt(ch,10,64)
			//numStack.Push(int(val))
		}
		//先判断index是否已经扫描到计算表达式的最后	
		if index+1 == len(exp) {
			break
		}
			index++
	}
	//如果扫描表达式完毕，依次从符号栈取出符号，从数栈中取出两个数，
	for {
		if operStack.Top==-1 {
			break
		}
		num1, _ := numStack.Pop()
		num2, _ := numStack.Pop()
		oper, _ := operStack.Pop()
		result := operStack.Cal(num1,num2,oper)
		numStack.Push(result)
	}
	res, _ := numStack.Pop()
	fmt.Printf("表达式%s= %v", exp,res)
}

/*
Output:
表达式30+2*6-2= 40
 */
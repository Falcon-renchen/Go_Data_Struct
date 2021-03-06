package main

import (
	"fmt"
	"gopkg.in/karalabe/cookiejar.v1/collections/deque"
)

//双端 队列
func main() {
	d := deque.New()
	element := []string{"foo","bar","baz"}
	for i := range element {
		d.PushLeft(element[i])
	}
	fmt.Println(d.PopLeft())
	fmt.Println(d.PopRight())
	fmt.Println(d.PopLeft())
}

/*
Output:
baz
foo
bar
*/
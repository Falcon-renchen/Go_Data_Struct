package main

import (
	"container/ring"
	"fmt"
)

var size int = 10


func main() {
	myRing := ring.New(size + 1)
	fmt.Println("Empty ring:", *myRing)

	for i := 0; i < myRing.Len()-1; i++ {
		myRing.Value = i
		myRing = myRing.Next()
	}
	myRing.Value = 2

	sum := 0
	myRing.Do(func(x interface{}) {
		t := x.(int)
		sum = sum + t
	})
	fmt.Println("Sum:", sum)

	for i := 0; i < myRing.Len()+2; i++ {
		myRing = myRing.Next()
		fmt.Print(myRing.Value, " ")
	}
	fmt.Println()
}

/*
Output:
Empty ring: {0xc00000c0a0 0xc00000c1c0 <nil>}
Sum: 47
0 1 2 3 4 5 6 7 8 9 2 0 1
*/
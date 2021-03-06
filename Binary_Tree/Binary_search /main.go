package main

import (
	"fmt"
	"sort"
)

func main() {
	intArray := []int{1,3,9,11,22,56,99,34}
	searchNum := 34
	sorted := sort.SearchInts(intArray, searchNum)	//如果切片不是升序排列，该api自动将切片升序排列然后输出找到的数字
	if sorted < len(intArray) {
		fmt.Printf("We found the number %d at %d",searchNum,  sorted)
	} else {
		fmt.Println("We not found number")
	}
}

//Output:
//We found the number 34 at 5
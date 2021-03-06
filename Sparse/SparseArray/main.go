package main

import "fmt"

type ValNode struct {
	row int
	col int
	val int
}

func main() {
	 var chessMap [11][11]int
	 chessMap[1][2] = 1	//黑子
	 chessMap[2][3] = 2 //白子

	for _,v := range chessMap {
		for _,v2 := range v {
			fmt.Printf("%d\t",v2)
		}
		fmt.Println()
	}

	var sparsearray []ValNode

	valNode := ValNode{
		row: 11,
		col: 11,
		val: 0,
	}
	sparsearray = append(sparsearray, valNode)


	for i,v := range chessMap {
		for j,v2 := range v {
			if v2 != 0 {
				valNode = ValNode{
					row: i,
					col: j,
					val: v2,
				}
				sparsearray = append(sparsearray, valNode)
			}
		}
	}

	for i,valNode := range sparsearray {
		fmt.Printf("%d: %d %d %d\n",i,valNode.row,valNode.col,valNode.val)
	}

	var chessMap2 [11][11]int
	for i, valNode := range sparsearray {
		if i!=0 {
			chessMap2[valNode.row][valNode.col] = valNode.val
		}
	}
	for _, v := range chessMap2{
		for _,v2 := range v {
			fmt.Printf("%d\t",v2)
		}
		fmt.Println()
	}

}

/*
Output:
0       0       0       0       0       0       0       0       0       0       0
0       0       1       0       0       0       0       0       0       0       0
0       0       0       2       0       0       0       0       0       0       0
0       0       0       0       0       0       0       0       0       0       0
0       0       0       0       0       0       0       0       0       0       0
0       0       0       0       0       0       0       0       0       0       0
0       0       0       0       0       0       0       0       0       0       0
0       0       0       0       0       0       0       0       0       0       0
0       0       0       0       0       0       0       0       0       0       0
0       0       0       0       0       0       0       0       0       0       0
0       0       0       0       0       0       0       0       0       0       0
0: 11 11 0
1: 1 2 1
2: 2 3 2
0       0       0       0       0       0       0       0       0       0       0
0       0       1       0       0       0       0       0       0       0       0
0       0       0       2       0       0       0       0       0       0       0
0       0       0       0       0       0       0       0       0       0       0
0       0       0       0       0       0       0       0       0       0       0
0       0       0       0       0       0       0       0       0       0       0
0       0       0       0       0       0       0       0       0       0       0
0       0       0       0       0       0       0       0       0       0       0
0       0       0       0       0       0       0       0       0       0       0
0       0       0       0       0       0       0       0       0       0       0
0       0       0       0       0       0       0       0       0       0       0
*/
package main

import "fmt"

const SIZE = 15

type Node struct {
	Value int
	Next *Node
}

type HashTable struct {
	Table map[int] *Node
	Size int
}

func hashFunction(i, size int) int {
	return i%size
}

func insert(hash *HashTable, value int) int {
	index := hashFunction(value, hash.Size)
	element := Node{Value: value, Next: hash.Table[index]}
	hash.Table[index] = &element
	return index
}

func traverse(hash *HashTable) {
	for k := range hash.Table {
		if hash.Table[k] != nil {
			t := hash.Table[k]
			for t != nil {
				fmt.Printf("%d -> ", t.Value)
				t = t.Next
			}
			fmt.Println()
		}
	}
}

func lookup(hash *HashTable, value int) bool {
	index := hashFunction(value, hash.Size)
	if hash.Table[index] != nil {
		t := hash.Table[index]
		for t != nil {
			if t.Value == value {
				return true
			}
			t = t.Next
		}
	}
	return false
}

func main() {
	table := make(map[int]*Node, SIZE)
	hash := &HashTable{Table: table, Size: SIZE}
	fmt.Println("Number of spaces:", hash.Size)
	for i := 0; i < 120; i++ {
		insert(hash, i)
	}
	traverse(hash)
	fmt.Println(lookup(hash,69))
}

/*
Output:
Number of spaces: 15
119 -> 104 -> 89 -> 74 -> 59 -> 44 -> 29 -> 14 ->
116 -> 101 -> 86 -> 71 -> 56 -> 41 -> 26 -> 11 ->
118 -> 103 -> 88 -> 73 -> 58 -> 43 -> 28 -> 13 ->
110 -> 95 -> 80 -> 65 -> 50 -> 35 -> 20 -> 5 ->
113 -> 98 -> 83 -> 68 -> 53 -> 38 -> 23 -> 8 ->
114 -> 99 -> 84 -> 69 -> 54 -> 39 -> 24 -> 9 ->
115 -> 100 -> 85 -> 70 -> 55 -> 40 -> 25 -> 10 ->
106 -> 91 -> 76 -> 61 -> 46 -> 31 -> 16 -> 1 ->
107 -> 92 -> 77 -> 62 -> 47 -> 32 -> 17 -> 2 ->
105 -> 90 -> 75 -> 60 -> 45 -> 30 -> 15 -> 0 ->
117 -> 102 -> 87 -> 72 -> 57 -> 42 -> 27 -> 12 ->
111 -> 96 -> 81 -> 66 -> 51 -> 36 -> 21 -> 6 ->
112 -> 97 -> 82 -> 67 -> 52 -> 37 -> 22 -> 7 ->
108 -> 93 -> 78 -> 63 -> 48 -> 33 -> 18 -> 3 ->
109 -> 94 -> 79 -> 64 -> 49 -> 34 -> 19 -> 4 ->
true
 */
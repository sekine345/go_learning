package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("len=%d, cap=%d, %v\n", len(s), cap(s), s)
}

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)
	// fmt.Println(&s)
	// len=6, cap=6, [2 3 5 7 11 13]

	s = s[:0] // sliceは参照型配列、大本を書き換えない
	printSlice(s)
	// fmt.Println(&s)
	// len=0, cap=6, []

	s = s[:4] // 切り取れる部分があると上書き？
	printSlice(s)
	// fmt.Println(&s)

	s = s[2:]
	printSlice(s)
	// fmt.Println(&s)
}

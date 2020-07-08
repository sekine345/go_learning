package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(s)

	s = s[1:4]
	fmt.Println(s)

	s = s[:2]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)
}

// output
// slice したもので書き換えられていき、slice の slice を作成していく
// [2 3 5 7 11 13]
// [3 5 7]
// [3 5]
// [5]

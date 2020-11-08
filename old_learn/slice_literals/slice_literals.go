package main

import "fmt"

func main() {
	q := []int{2, 3, 5, 7, 11, 13} //[]に配列の長さを書かなくても定義できる
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct { //struct も配列でかける
		i int
		b bool // カンマなし
	}{
		{2, true}, //要素は{}で囲んで表す、カンマあり
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
}

// [2 3 5 7 11 13]
// [true false true true false true]
// [{2 true} {3 false} {5 true} {7 true} {11 false} {13 true}]

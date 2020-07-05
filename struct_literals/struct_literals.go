package main

import "fmt"

type Vertex struct {
	X, Y int // 型が同じなら列挙できる
}

var (
	v1 = Vertex{1, 2} // X, Y を初期化
	v2 = Vertex{X: 1} // X のみ初期化
	v3 = Vertex{}     // 初期化なし、X, Y のどちらも 0
	p1 = &Vertex{1, 2}
	p2 = &v1
) // struct を列挙し定義できる

func main() {
	fmt.Println(v1, v2, v3)
	fmt.Println(p1, p2)
}

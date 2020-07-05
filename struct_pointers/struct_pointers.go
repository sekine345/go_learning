package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	p := &v        // struct 型の v のメモリアドレスを取得
	fmt.Println(p) // アドレスでは表示されない
	fmt.Println(*p)
	fmt.Println(v)
	p.X = 1e9 // ポインタを通して Vertex にアクセス、X を上書き
	fmt.Println(v)
	fmt.Println(*p)
	fmt.Println(p)
}

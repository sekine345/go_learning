package main

import "fmt"

func main() {
	i, j := 42, 2701
	fmt.Println(i, j)

	p := &i         // i のメモリアドレスを取得、ポインタを定義するため ":="
	fmt.Println(*p) // p(ポインタ)を通して、i の値を表示
	fmt.Println(p)

	*p = 21 // p(ポインタ)を通して、i の値を書き換える
	fmt.Println(i)
	fmt.Println(*p)
	fmt.Println(p)

	p = &j       // j のメモリアドレスを取得、すでに定義されている p を変更するため "="
	*p = *p / 37 // ポインタを通して、j の値を書き換える
	fmt.Println(j)
	fmt.Println(*p)
	fmt.Println(p)
}

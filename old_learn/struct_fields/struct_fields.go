package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2} // Vertex 定義
	fmt.Println(v)

	v.X = 4 // VertexのXにアクセス、書き換え
	fmt.Println(v)
}

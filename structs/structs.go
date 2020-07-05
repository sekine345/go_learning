package main

import "fmt"

type Vertex struct { //struct: 複数の値を保持できる型
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{1, 2}) //Vertex に 1, 2という値を定義
}

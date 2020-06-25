package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ { // iを0と初期化・定義; iが10より小さい限り; ++:1ずつ加算
		sum += i
		fmt.Printf("sum: %v, i %v \n", sum, i)
	}
	fmt.Println(sum)
}

package main

import "fmt"

func main() {
	i := 42
	f := 3.14
	g := 0.867 + 0.5i
	fmt.Printf("i, %T, %v \n", i, i)
	fmt.Printf("f, %T, %v \n", f, f)
	fmt.Printf("g, %T, %v \n", g, g)
}

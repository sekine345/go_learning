package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Pi)
	// fmt.Println(math.pi)
	// # command-line-arguments
	// ./exported_names.go:9:14: cannot refer to unexported name math.pi
	// ./exported_names.go:9:14: undefined: math.pi
}

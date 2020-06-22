package main

import (
	"fmt"
	"math/cmplx"
)

//複数の package を一度に import できる。","" はなしに、並べる事ができる

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1            //符号なし整数型
	z      complex128 = cmplx.Sqrt(-5 + 12i) //実数部・虚数部をfloat64で表現する複素数型
)

func main() {
	fmt.Printf("Type: %T , Value: %v\n", ToBe, ToBe) //%T:Type, %v:values
	fmt.Printf("Type: %T , Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T , Value: %v\n", z, z)
}

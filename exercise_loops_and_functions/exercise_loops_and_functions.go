package main

import (
	"fmt"
)

// func sqrt(x float64) float64 {
// 	z := float64(1)
// 	for i := 0; i < 10; i++ {
// 		z -= (z*z - x) / (2 * z)
// 		fmt.Println(z)
// 	}
// 	return z
// }

func sqrt(x float64) float64 {
	z := 1.0
	i := 1
	for d := 1.0; d*d > 1e-20; z -= d { //for文内で定義したdのような変数もfor文の定義に使える
		d = (z*z - x) / (2 * z)
		fmt.Println(i, z, d*d)
		i += 1
	}
	return z
}

func main() {
	fmt.Println(sqrt(2))
}

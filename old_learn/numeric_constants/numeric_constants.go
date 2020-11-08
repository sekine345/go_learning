package main

import "fmt"

const (
	// Big is huge number
	Big = 1 << 100
	// Small is shift it right again 99 places
	Small = Big >> 99
)

func needInt(x int) int {
	return x*10 + 1
}

func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	//fmt.Println(Big)
	//fmt.Println(Small)

	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}

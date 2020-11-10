package main

import (
	"fmt"
	"math/rand"
)

/*パッケージ名はインポートパスの最後の要素と同じ名前になる*/

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
}

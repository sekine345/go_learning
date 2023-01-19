// main.goの分割

/*
同じ階層、パッケージのapp.goの関数はimportなしに呼び出せる
*/

/*[terminal]
$ go run main.go
# command-line-arguments
./main.go:18:17: undefined: AppName
$
$ go run main.go app.go
Zoo Aplication
Grass
Banana
Carrot
$
$ go run *.go  // ワイルドカードを使って実行できる
Zoo Aplication
Grass
Banana
Carrot
*/


package main

import (
    "fmt"
    "zoo/animals"
)


func main() {
    fmt.Println(AppName()) // "main.AppName()でないことに注目"

    fmt.Println(animals.ElephantFeed())
    fmt.Println(animals.MonkeyFeed())
    fmt.Println(animals.RabbitFeed())
}


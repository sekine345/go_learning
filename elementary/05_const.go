// 定数
// 一度定義したら変更しないもの

package main

import "fmt"


func main(){
    //const name = "takeuchi"
    //name = "akimoto"
    //fmt.Println(name)
    /*
    error:
        ./const.go:11:5: cannot assign to name (untyped string constant "takeuchi")
    */

    const (
        sun = iota  // iota: 0から順に1ずつ増やす
        mon  // 1
        tue  // 2
    )
    fmt.Println(sun, mon, tue)
}

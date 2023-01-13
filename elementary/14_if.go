// if
/*

- 比較演算子
> >= < <= == !=

- 論理演算子
&& || !
*/


package main

import "fmt"


func check_score(score int) {
    if score > 80 {
        fmt.Println("Great!")
    } else if score > 60{
        fmt.Println("Good!")
    } else {
        fmt.Println("so, so...")
    }
}

func main() {
    check_score(85)  // "Great!"
    check_score(65)  // "Good!"
    check_score(40)  // "so, so..."

    // スコープ
    // if節でしか使わない変数を定義できるが、if節以外では使えないことに注意
    // else は省略可能
    if ok := true; ok {
        fmt.Println("OK!")
    }
    // fmt.Println(ok)  // undefined: ok
}

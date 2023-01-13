// ポインタ
// アドレスを指し示す変数、メモリの削減に使える
// 演算はできない


package main

import "fmt"


func main(){
    a := 5
    var pa *int
    /*
    pa: aのポインタの宣言
    *int: int型のポインタと明示
    */

    pa = &a  // paにaのアドレスを入れる
    // paの領域にあるデータの値 = *pa
    fmt.Println(pa)
    fmt.Println(*pa)
}

// goroutine: 並行処理
// 複数の処理を同時に実行する


package main

import (
    "fmt"
    "time"
)


func task1() {
    // 思い処理のシミュレート
    time.Sleep(time.Second * 2)  // 2秒 Sleep
    fmt.Println("Task1 finished!")
}

func task2() {
    fmt.Println("Task2 finished!")
}

func main() {
    // いつも通りの実行、task1が終わったらtask2を実行している
    // task1()
    // task2()
    /* [terminal]
    Task1 finished!
    Task2 finished!
    */


    // 並行処理での実行
    go task1()
    go task2()  // goroutine が終わる前に main 関数自体が終わるので、これだけだとターミナルに何も出ない
    time.Sleep(time.Second * 3)
    /* [terminal]
    Task2 finished!
    Task1 finished!
    */
}

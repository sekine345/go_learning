// channel
// 並行処理で実行しているタスクから値を受け取る時に使える


package main

import (
    "fmt"
    "time"
)

// channelを使いたい間数で引数に"chan"を明示
func task1(result chan string) {
    time.Sleep(time.Second * 2)
    result <- "Task1 finished!" // channelへの値の格納
}

func task2() {
    fmt.Println("Task2 finished!")
}

func main() {
    // channelの定義
    result := make(chan string)

    go task1(result)
    go task2()

    fmt.Println(<-result) // resultに値が格納されるまで待機する
    /* [terminal]
    Task2 finished!
    Task1 finished!
    */
}

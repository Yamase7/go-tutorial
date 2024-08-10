package main

import (
    "fmt"
    "github.com/Yamase7/go-tutorial/greetings"
)

func main(){
    // 挨拶のメッセージをもらって、出力をする
    message := greetings.Hello("Gladys")
    fmt.Println(message)
}

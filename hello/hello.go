package main

import (
    "fmt"
    "log"

    "github.com/Yamase7/go-tutorial/greetings"
)

func main() {
    // ログエントリーの接頭辞と、時間・ソースファイル・行番号の表示を無効にするフラグなど、定義済みロガーのプロパティを設定します。
    log.SetPrefix("greetings: ")
    log.SetFlags(0)

    // あいさつを要求.
    message, err := greetings.Hello("Gladys")
    
    // エラーが返ってきたらエラーログをコンソールに出力してプログラムを終了する。
    if err != nil {
        log.Fatal(err)
    }

    // エラーがなければ受け取ったメッセージを返す。
    fmt.Println(message)
}

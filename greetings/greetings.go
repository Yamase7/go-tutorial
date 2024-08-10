package greetings

import (
    "errors"
    "fmt"
)

// Hello関数は指定された人物への挨拶を返す。
func Hello(name string) (string, error) {
    // 名前が与えられていない場合はメッセージとともにエラーを返す。
    if name == "" {
        return "", errors.New("empty name")
    }

    // 名前を受け取った場合は、名前を埋め込んだ挨拶メッセージを返す。
    message := fmt.Sprintf("Hi, %v. Welcome!", name)
    return message, nil
}

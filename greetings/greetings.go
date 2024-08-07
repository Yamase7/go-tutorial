package greetings

import "fmt"

// 指定された名前の挨拶を返す
func Hello(name string) string {
    // Return a greeting that embeds the name in a message.
    message := fmt.Sprintf("Hi, %v. Welcome!", name)
    return message
}
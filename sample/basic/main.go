package main

import "fmt"

func main() {
	fmt.Println("=== Basic DI Sample with Wire ===")

	// Wireが生成したInitializeUserHandler関数を使用
	handler := InitializeUserHandler()

	// ハンドラーを使ってリクエストを処理
	handler.Handle(1)
	handler.Handle(2)
	handler.Handle(3)
}

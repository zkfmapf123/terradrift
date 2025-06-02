package main

import (
	"fmt"
	"os"
)

func main() {
	// GitHub Actions에서 입력값을 환경 변수로 받습니다
	name := os.Getenv("INPUT_NAME")
	if name == "" {
		name = "World" // 기본값
	}

	// Hello World 메시지 출력
	fmt.Printf("Hello, %s!\n", name)
}

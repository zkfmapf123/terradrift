package main

import (
	"fmt"
	"os"

	"github.com/zkfmapf123/terradrift/models"
)

// 파라미터 설정 메서드
func parameterInit() models.TerraDriftInputParams {

	return models.TerraDriftInput(
		models.WithConcurreny(os.Getenv("INPUT_CONCURRENCY")),
		models.WithSlackChannel(os.Getenv("INPUT_SLACK_CHANNEL")),
		models.WithSlackToken(os.Getenv("INPUT_SLACK_TOKEN")),
	)

}

/*
- 파일구조 파악
- 파일내에서 terraform , terragrunt path 확인하기
- path 에서 paln 후 결과 모으기
- slack message
*/
func main() {
	params := parameterInit()

	// Hello World 메시지 출력
	fmt.Printf("%s %s %d\n", params.SlackParams.Token, params.SlackParams.Channel, params.Concurrency)
}

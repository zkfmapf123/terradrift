package main

import (
	"fmt"
	"os"

	"github.com/zkfmapf123/terradrift/models"
)

func parameterInit() models.TerraDriftInputParams {

	return models.TerraDriftInput(
		models.WithConcurreny(os.Getenv("INPUT_CONCURRENCY")),
		models.WithSlackChannel(os.Getenv("INPUT_SLACK_CHANNEL")),
		models.WithSlackToken(os.Getenv("INPUT_SLACK_TOKEN")),
	)

}

func main() {
	params := parameterInit()

	// Hello World 메시지 출력
	fmt.Printf("%s %s %d\n", params.SlackParams.Token, params.SlackParams.Channel, params.Concurrency)
}

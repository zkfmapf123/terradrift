package main

import (
	"fmt"
	"os"

	"github.com/zkfmapf123/terradrift/models"
)

func parameterInit() models.TerraDriftInputParams {

	return models.TerraDriftInput(
		models.WithConcurreny(os.Getenv("INPUT_CONCURRENCY")),
		models.WithIsUseTerraformPath(os.Getenv("INPUT_WITH_IS_USE_TERRAFORM_PATH")),
		models.WithIsUseTerragruntPath(os.Getenv("INPUT_WITH_IS_USE_TERRAGRUNT_PATH")),
		models.WithSlackChannel(os.Getenv("INPUT_SLACK_CHANNEL")),
		models.WithSlackToken(os.Getenv("INPUT_SLACK_TOKEN")),
	)

}

func main() {
	params := parameterInit()
	fmt.Println(params)

	// Hello World 메시지 출력
	fmt.Printf("Hello, %s!\n", name)
}

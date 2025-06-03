package main

import (
	"fmt"
	"os"
	"time"

	"github.com/zkfmapf123/donggo"
	"github.com/zkfmapf123/terradrift/intenral/cmd"
	"github.com/zkfmapf123/terradrift/intenral/terraform"
	"github.com/zkfmapf123/terradrift/intenral/terragrunt"
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
  - Terraform -> provider.tf
  - Terragrunt -> terragrunt.hcl

- path 에서 paln 후 결과 모으기
- slack message
*/

var (
	COMMAND_LOOP = []string{"terraform", "terragrunt"}
)

func main() {
	params := parameterInit()

	tfPaths, tgPaths, err := cmd.GetCurrentDirOrFile()
	if err != nil {
		panic(err)
	}

	iacManager := map[string]models.DriftResultFuncs{
		"terraform":  terraform.New(),
		"terragrunt": terragrunt.New(),
	}

	iacManager["terraform"].AllPush(donggo.OKeys(tfPaths))
	iacManager["terragrunt"].AllPush(donggo.OKeys(tgPaths))

	started := time.Now()
	// plan
	for _, v := range COMMAND_LOOP {
		iacManager[v].Plan()
	}

	end := time.Since(started)
	fmt.Println("method time : %d ms\n", end.Milliseconds())
}

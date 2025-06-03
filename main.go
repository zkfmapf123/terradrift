package main

import (
	"fmt"
	"maps"
	"os"
	"time"

	"github.com/ashwanthkumar/slack-go-webhook"
	"github.com/zkfmapf123/terradrift/intenral/cmd"
	"github.com/zkfmapf123/terradrift/intenral/strings"
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
- 1. 파일구조 파악
- 2. 파일내에서 terraform , terragrunt path 확인하기
  - Terraform -> provider.tf
  - Terragrunt -> terragrunt.hcl
- 2.1 TG / TF Path 별 중복검사

- 3. path 에서 paln 후 결과 모으기
- 4. slack message
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

	tgPathArr, tfPathArr := strings.ParsingClear(tgPaths, tfPaths)
	iacManager["terraform"].AllPush(tfPathArr)
	iacManager["terragrunt"].AllPush(tgPathArr)

	// Run....
	started := time.Now()
	resultReport := run(params.Concurrency, iacManager)

	// slack send
	if params.SlackParams.Token == "" || params.SlackParams.Channel == "" {
		fmt.Println("No Slack Values...")

		for path, res := range resultReport {
			fmt.Printf("path : %s add : %s changes : %s destory : %s", path, res.Add, res.Change, res.Destroy)
		}

	} else {
		sendSlack(params.SlackParams.Channel, params.SlackParams.Token, resultReport)
	}

	end := time.Since(started)
	fmt.Printf("method time : %d ms\n", end.Milliseconds())
	return
}

func run(concurreny int, iacManager map[string]models.DriftResultFuncs) map[string]models.DriftResultsParams {

	result := make(map[string]models.DriftResultsParams)

	// plan
	for _, v := range COMMAND_LOOP {

		r := iacManager[v].Plan(concurreny)
		maps.Copy(result, r)
	}

	return result
}

func sendSlack(slackChannel string, slackToken string, report map[string]models.DriftResultsParams) {

	if slackChannel == "" || slackToken == "" {
		fmt.Println("No Slack Values...")
		return
	}

	attach := slack.Attachment{}

	for path, result := range report {
		attach.AddField(slack.Field{Title: "Path", Value: path, Short: true})

		if result.Add == "0" && result.Change == "0" && result.Destroy == "0" {
			attach.AddField(slack.Field{Title: "Result", Value: "No Changes...", Short: false})
			continue
		}

		attach.AddField(slack.Field{Title: "Result", Value: fmt.Sprintf("Add : %s, Change : %s, Destory : %s", result.Add, result.Change, result.Destroy), Short: false})
	}

	p := slack.Payload{
		Text:        "TerraDirft Reports",
		Channel:     slackChannel,
		Attachments: []slack.Attachment{attach},
	}

	err := slack.Send(slackToken, "", p)
	if len(err) > 0 {
		fmt.Println(err)
	}
}

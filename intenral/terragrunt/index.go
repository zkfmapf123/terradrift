package terragrunt

import (
	"fmt"
	"log"
	"sync"

	"github.com/zkfmapf123/terradrift/intenral/cmd"
	"github.com/zkfmapf123/terradrift/intenral/strings"
	"github.com/zkfmapf123/terradrift/models"
)

type TerragruntParams struct {
	Method    string
	IaCParams models.IaCParams
}

func New() *TerragruntParams {

	m := make(map[string]models.DriftResultsParams)

	return &TerragruntParams{
		Method: "terragrunt",
		IaCParams: models.IaCParams{
			PlanPath: nil,
			Results:  m,
		},
	}
}

func (t *TerragruntParams) AllPush(paths []string) {
	m := make([]string, len(paths))

	copy(m, paths)
	t.IaCParams.PlanPath = m
}

func (t *TerragruntParams) Push(path string) {
	t.IaCParams.PlanPath = append(t.IaCParams.PlanPath, path)
}

// terragrunt 가 terraform 자체를 한번더 감싸서 뭔ㄱ ㅏ안잡히는것같음...
func (t *TerragruntParams) Plan(concurreny int) map[string]models.DriftResultsParams {
	planResult := make(map[string]models.DriftResultsParams)
	resultChan := make(chan struct {
		path   string
		result models.DriftResultsParams
	})

	workerpool := make(chan struct{}, concurreny)
	var wg sync.WaitGroup

	for _, path := range t.IaCParams.PlanPath {
		wg.Add(1)

		go func(p string) {

			defer wg.Done()
			workerpool <- struct{}{}
			defer func() { <-workerpool }()

			cmd.Exec("terragrunt", fmt.Sprintf("--terragrunt-working-dir=%s", p), "init")
			b, err := cmd.Exec("terragrunt", fmt.Sprintf("--terragrunt-working-dir=%s", p), "plan")
			if err != nil {
				log.Fatalln("[Terragrunt Error] path : ", p, " output : ", string(b), " err : ", err)
			}

			result := strings.IaCParsing(b)
			resultChan <- struct {
				path   string
				result models.DriftResultsParams
			}{
				path: p,
				result: models.DriftResultsParams{
					Add:     result.Add,
					Change:  result.Change,
					Destroy: result.Destroy,
				},
			}

		}(path)

		go func() {
			wg.Wait()
			close(resultChan)
		}()

		for r := range resultChan {
			planResult[r.path] = r.result
		}
	}

	return planResult
}

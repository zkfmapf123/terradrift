package terraform

import (
	"fmt"
	"log"
	"sync"

	"github.com/zkfmapf123/terradrift/intenral/cmd"
	"github.com/zkfmapf123/terradrift/intenral/strings"
	"github.com/zkfmapf123/terradrift/models"
)

type TerraformParams struct {
	Method    string
	IaCParams models.IaCParams
}

func New() *TerraformParams {

	m := make(map[string]models.DriftResultsParams)

	return &TerraformParams{
		Method: "terraform",
		IaCParams: models.IaCParams{
			PlanPath: nil,
			Results:  m,
		},
	}
}

func (t *TerraformParams) AllPush(paths []string) {
	m := make([]string, len(paths))

	copy(m, paths)
	t.IaCParams.PlanPath = m
}

func (t *TerraformParams) Push(path string) {
	t.IaCParams.PlanPath = append(t.IaCParams.PlanPath, path)
}

func (t *TerraformParams) Plan(concurrency int) map[string]models.DriftResultsParams {
	planResult := make(map[string]models.DriftResultsParams)
	resultChan := make(chan struct {
		path   string
		result models.DriftResultsParams
	})

	workerpool := make(chan struct{}, concurrency)
	var wg sync.WaitGroup

	for _, path := range t.IaCParams.PlanPath {
		wg.Add(1)

		go func(p string) {

			defer wg.Done()
			workerpool <- struct{}{} // worker
			defer func() { <-workerpool }()

			cmd.Exec("terraform", fmt.Sprintf("-chdir=%s", p), "init")
			b, err := cmd.Exec("terraform", fmt.Sprintf("-chdir=%s", p), "plan")
			if err != nil {
				log.Fatalln("[Terraform Error] path : ", p, " output : ", string(b), " err : ", err)
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
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	for r := range resultChan {
		planResult[r.path] = r.result
	}

	return planResult
}

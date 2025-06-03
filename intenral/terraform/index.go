package terraform

import (
	"fmt"
	"log"

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
	for _, path := range t.IaCParams.PlanPath {
		cmd.Exec("terraform", fmt.Sprintf("-chdir=%s", path), "init")
		b, err := cmd.Exec("terraform", fmt.Sprintf("-chdir=%s", path), "plan")
		if err != nil {
			log.Fatalln("output : ", string(b), "err : ", err)
		}

		result := strings.IaCParsing(b)
		planResult[path] = models.DriftResultsParams{
			Add:     result.Add,
			Change:  result.Change,
			Destroy: result.Destroy,
		}
	}

	return planResult
}

package terragrunt

import (
	"fmt"
	"log"

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
	for _, path := range t.IaCParams.PlanPath {

		cmd.Exec("terragrunt", fmt.Sprintf("--terragrunt-working-dir=%s", path), "init")
		b, err := cmd.Exec("terragrunt", fmt.Sprintf("--terragrunt-working-dir=%s", path), "plan")
		if err != nil {
			log.Fatalln("[Terragrunt Error] path : ", path, " output : ", string(b), " err : ", err)
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

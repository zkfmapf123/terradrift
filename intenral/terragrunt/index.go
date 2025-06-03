package terragrunt

import (
	"fmt"

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

func (t *TerragruntParams) Plan() {
	for _, path := range t.IaCParams.PlanPath {
		fmt.Printf("terragrunt : %s\n", path)
	}
}

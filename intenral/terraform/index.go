package terraform

import (
	"fmt"

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

	copy(paths, m)
	t.IaCParams.PlanPath = m
}

func (t *TerraformParams) Push(path string) {
	t.IaCParams.PlanPath = append(t.IaCParams.PlanPath, path)
}

func (t *TerraformParams) Plan() {

	for _, path := range t.IaCParams.PlanPath {
		fmt.Printf("terraform : %s\n", path)
	}
}

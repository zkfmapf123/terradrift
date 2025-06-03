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

	copy(m, paths)
	t.IaCParams.PlanPath = m
}

func (t *TerraformParams) Push(path string) {
	t.IaCParams.PlanPath = append(t.IaCParams.PlanPath, path)
}

func (t *TerraformParams) Plan(concurrency int) {

	for _, path := range t.IaCParams.PlanPath {
		// b, err := cmd.Exec("terraform", fmt.Sprintf("-chdir=%s", path), "plan")
		// if err != nil {
		// 	panic(err)
		// }

		fmt.Println("terraform : ", path)
	}
}

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
			PlanPath: []string{},
			Results:  m,
		},
	}
}

func (tf *TerraformParams) Push(path string) {
	tf.IaCParams.PlanPath = append(tf.IaCParams.PlanPath, path)
}

func Plan(concurrency int) {
	fmt.Println("terraform")
}

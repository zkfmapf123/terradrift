package strings

import (
	"fmt"
	"regexp"

	"github.com/zkfmapf123/terradrift/models"
)

func TerraformParsing(b []byte) models.DriftResultsParams {

	re := regexp.MustCompile(`Plan: (\d+) to add, (\d+) to change, (\d+) to destroy\.`)
	matches := re.FindStringSubmatch(string(b))

	fmt.Println(matches)

	// plan
	if len(matches) == 4 {
		return models.DriftResultsParams{
			Add:     matches[1],
			Change:  matches[2],
			Destroy: matches[3],
		}
	}

	// no changes
	return models.DriftResultsParams{
		Add:     "0",
		Change:  "0",
		Destroy: "0",
	}
}

// parsing clean
// Terragrunt, Terraform 중복 path 삭제 (tgPaths > tfPaths)
func ParsingClear(tgPaths map[string]bool, tfPaths map[string]bool) ([]string, []string) {

	tgPathArr, tfPathArr := []string{}, []string{}

	// terragrunt
	for path, _ := range tgPaths {
		tfPaths[path] = false

		tgPathArr = append(tgPathArr, path)
	}

	// terraform
	for path, isValue := range tfPaths {

		if isValue {
			tfPathArr = append(tfPathArr, path)
		}
	}

	return tgPathArr, tfPathArr

}

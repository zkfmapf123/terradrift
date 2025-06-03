package strings

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/zkfmapf123/terradrift/models"
)

func IaCParsing(b []byte) models.DriftResultsParams {

	lines := strings.Split(string(b), "\n")

	for _, line := range lines {
		if strings.Contains(line, "Plan:") {

			planLines := strings.Split(line, "Plan:")
			re := regexp.MustCompile(`(\d+)\s+to add,\s+(\d+)\s+to change,\s+(\d+)\s+to destroy`)
			matches := re.FindStringSubmatch(planLines[1])

			return models.DriftResultsParams{
				Add:     matches[1],
				Change:  matches[2],
				Destroy: matches[3],
			}
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

func ParsingSlackMessage(path string, result models.DriftResultsParams) (string, string) {

	path = strings.ReplaceAll(path, "/github/workspace", "")
	res := ""

	if result.Add == "0" || result.Change == "0" || result.Destroy == "0" {
		res = "No Changes"
	} else {
		res = fmt.Sprintf("Add %s, Change %s, Destory %s", result.Add, result.Change, result.Destroy)
	}

	return path, fmt.Sprintf("Result : %s", res)

}

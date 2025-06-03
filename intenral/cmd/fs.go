package cmd

import (
	"os"
	"path/filepath"
	"strings"
)

var (
	EXCLUDE_DIR = map[string]bool{
		".git":         true,
		".github":      true,
		"node_modules": true,
		"vendor":       true,
		"dist":         true,
		"build":        true,
		".terraform":   true,
	}
)

func GetCurrentDirOrFile() (map[string]bool, map[string]bool, error) {

	tfPaths, tgPaths := map[string]bool{}, map[string]bool{}

	currentDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	return getWalk(currentDir, tfPaths, tgPaths)

}

func getWalk(currentDir string, tfPaths map[string]bool, tgPaths map[string]bool) (map[string]bool, map[string]bool, error) {
	entries, err := os.ReadDir(currentDir)

	if err != nil {
		return nil, nil, err
	}

	for _, entry := range entries {

		path := filepath.Join(currentDir, entry.Name())

		if entry.IsDir() {

			// skip
			if EXCLUDE_DIR[entry.Name()] {
				continue
			}

			getWalk(path, tfPaths, tgPaths)
		} else {

			parentDir := filepath.Dir(path)
			// terragrunt 검사
			if entry.Name() == "terragrunt.hcl" {
				tgPaths[parentDir] = true
				continue
			}

			// terraform 검사
			if strings.HasSuffix(entry.Name(), ".tf") {
				tfPaths[parentDir] = true
				continue
			}
		}
	}

	return tfPaths, tgPaths, nil
}

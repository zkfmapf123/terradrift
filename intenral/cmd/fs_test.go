package cmd

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Shell(t *testing.T) {

	_, err := Exec("ls", "-al")

	assert.Equal(t, err, nil)
}

func Test_CurrentDir(t *testing.T) {

	tgPath, tfPath := map[string]bool{}, map[string]bool{}

	tfPath, tgPath, _ = getWalk("/Users/idong-gyu/dev/github-actions/terradrift/__test__", tfPath, tgPath)

	fmt.Println("terraform path >> ", tfPath)
	fmt.Println("terragrunt path >> ", tgPath)

	assert.Equal(t, len(tfPath), 3)
	assert.Equal(t, len(tgPath), 2)
}

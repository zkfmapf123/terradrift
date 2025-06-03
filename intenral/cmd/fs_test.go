package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Shell(t *testing.T) {

	_, err := Exec("ls", "-al")

	assert.Equal(t, err, nil)
}

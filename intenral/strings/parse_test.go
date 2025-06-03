package strings

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ParsingClear(t *testing.T) {

	tgPaths := map[string]bool{
		"/a": true,
		"/b": true,
		"/c": true,
	}

	tfPaths := map[string]bool{
		"/a": true,
		"/f": true,
		"/g": true,
	}

	tg, tf := ParsingClear(tgPaths, tfPaths)

	assert.Equal(t, len(tg), 3)
	assert.Equal(t, len(tf), 2)

	for _, t := range tf {
		if t == "/a" {
			log.Fatalf("Must be not path %s\n", t)
		}
	}
}

func Test_TerraformParsing_1(t *testing.T) {

	text := `
		Terraform will perform the following actions:

  # module.vpc.aws_eip.nat_eip[0] will be created
  ...
Plan: 2 to add, 1 to change, 0 to destroy.
	`

	TerraformParsing([]byte(text))

}

func Test_TerraformParsing_2(t *testing.T) {

	text := `
		Changes to Outputs:
  + v = "hello world"
	`

	TerraformParsing([]byte(text))

}

func Test_TerraformParsing_3(t *testing.T) {

	text := `
		No changes. Your infrastructure matches the configuration.
	`

	TerraformParsing([]byte(text))

}

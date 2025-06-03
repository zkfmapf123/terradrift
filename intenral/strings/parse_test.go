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

	v := IaCParsing([]byte(text))
	assert.Equal(t, v.Add, "2")
	assert.Equal(t, v.Change, "1")
	assert.Equal(t, v.Destroy, "0")

}

func Test_TerraformParsing_2(t *testing.T) {

	text := `
		Changes to Outputs:
  + v = "hello world"
	`

	v := IaCParsing([]byte(text))
	assert.Equal(t, v.Add, "0")
	assert.Equal(t, v.Change, "0")
	assert.Equal(t, v.Destroy, "0")

}

func Test_TerraformParsing_3(t *testing.T) {

	text := `
		No changes. Your infrastructure matches the configuration.
	`
	v := IaCParsing([]byte(text))
	assert.Equal(t, v.Add, "0")
	assert.Equal(t, v.Change, "0")
	assert.Equal(t, v.Destroy, "0")

}

func Test_TerraformParsing_4(t *testing.T) {
	text := `
		15:23:13.637 STDOUT terraform:     }
15:23:13.637 STDOUT terraform:   # module.alb.aws_lb_target_group.lb_443_tg will be created
15:23:13.637 STDOUT terraform:   + resource "aws_lb_target_group" "lb_443_tg" {
15:23:13.637 STDOUT terraform:       + arn                                = (known after apply)
15:23:13.637 STDOUT terraform:       + arn_suffix                         = (known after apply)
15:23:13.637 STDOUT terraform:       + connection_termination             = (known after apply)
15:23:13.637 STDOUT terraform:       + deregistration_delay               = "300"
15:23:13.637 STDOUT terraform:       + id                                 = (known after apply)
15:23:13.637 STDOUT terraform:       + ip_address_type                    = (known after apply)
15:23:13.637 STDOUT terraform:       + lambda_multi_value_headers_enabled = false
15:23:13.637 STDOUT terraform:       + load_balancer_arns                 = (known after apply)
15:23:13.637 STDOUT terraform:       + load_balancing_algorithm_type      = (known after apply)
15:23:13.637 STDOUT terraform:       + load_balancing_anomaly_mitigation  = (known after apply)
15:23:13.637 STDOUT terraform:       + load_balancing_cross_zone_enabled  = (known after apply)
15:23:13.637 STDOUT terraform:       + name                               = "test-default-tg"
15:23:13.637 STDOUT terraform:       + name_prefix                        = (known after apply)
15:23:13.637 STDOUT terraform:       + port                               = 80
15:23:13.637 STDOUT terraform:       + preserve_client_ip                 = (known after apply)
15:23:13.637 STDOUT terraform:       + protocol                           = "HTTP"
15:23:13.637 STDOUT terraform:       + protocol_version                   = (known after apply)
15:23:13.637 STDOUT terraform:       + proxy_protocol_v2                  = false
15:23:13.637 STDOUT terraform:       + slow_start                         = 0
15:23:13.637 STDOUT terraform:       + tags_all                           = (known after apply)
15:23:13.637 STDOUT terraform:       + target_type                        = "ip"
15:23:13.637 STDOUT terraform:       + vpc_id                             = "vpc-0b61fb68b5edbdd48"
15:23:13.637 STDOUT terraform:       + health_check (known after apply)
15:23:13.637 STDOUT terraform:       + stickiness (known after apply)
15:23:13.637 STDOUT terraform:       + target_failover (known after apply)
15:23:13.637 STDOUT terraform:       + target_group_health (known after apply)
15:23:13.637 STDOUT terraform:       + target_health_state (known after apply)
15:23:13.637 STDOUT terraform:     }
15:23:13.637 STDOUT terraform: Plan: 7 to add, 0 to change, 0 to destroy.
15:23:13.637 STDOUT terraform:
15:23:13.637 STDOUT terraform: ─────────────────────────────────────────────────────────────────────────────
15:23:13.637 STDOUT terraform: Note: You didn't use the -out option to save this plan, so Terraform can't
15:23:13.637 STDOUT terraform: guarantee to take exactly these actions if you run "terraform apply" now.
	`

	v := IaCParsing([]byte(text))
	assert.Equal(t, v.Add, "7")
	assert.Equal(t, v.Change, "0")
	assert.Equal(t, v.Destroy, "0")

}

package aws

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccDataSourceAwsRoute53ResolverRules_basic(t *testing.T) {
	dsResourceName := "data.aws_route53_resolver_rules.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t); testAccPreCheckAWSRoute53Resolver(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceAwsRoute53ResolverRules_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dsResourceName, "resolver_rule_ids.#", "1"),
					resource.TestCheckResourceAttr(dsResourceName, "resolver_rule_ids.1743502667", "rslvr-autodefined-rr-internet-resolver"),
				),
			},
		},
	})
}

func TestAccDataSourceAwsRoute53ResolverRules_ResolverEndpointId(t *testing.T) {
	rName1 := fmt.Sprintf("tf-testacc-r53-resolver-%s", acctest.RandStringFromCharSet(8, acctest.CharSetAlphaNum))
	rName2 := fmt.Sprintf("tf-testacc-r53-resolver-%s", acctest.RandStringFromCharSet(8, acctest.CharSetAlphaNum))
	ds1ResourceName := "data.aws_route53_resolver_rules.by_resolver_endpoint_id"
	ds2ResourceName := "data.aws_route53_resolver_rules.by_rule_type"
	ds3ResourceName := "data.aws_route53_resolver_rules.by_share_status"
	ds4ResourceName := "data.aws_route53_resolver_rules.by_invalid_owner_id"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t); testAccPreCheckAWSRoute53Resolver(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceAwsRoute53ResolverRules_resolverEndpointId(rName1, rName2),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(ds1ResourceName, "resolver_rule_ids.#", "1"),
					resource.TestCheckResourceAttr(ds2ResourceName, "resolver_rule_ids.#", "1"),
					resource.TestCheckResourceAttr(ds3ResourceName, "resolver_rule_ids.#", "2"),
					resource.TestCheckResourceAttr(ds4ResourceName, "resolver_rule_ids.#", "0"),
				),
			},
		},
	})
}

const testAccDataSourceAwsRoute53ResolverRules_basic = `
# The default Internet Resolver rule.
data "aws_route53_resolver_rules" "test" {
  owner_id     = "Route 53 Resolver"
  rule_type    = "RECURSIVE"
  share_status = "NOT_SHARED"
}
`

func testAccDataSourceAwsRoute53ResolverRules_resolverEndpointId(rName1, rName2 string) string {
	return testAccRoute53ResolverRuleConfig_resolverEndpoint(rName1) + fmt.Sprintf(`
resource "aws_route53_resolver_rule" "forward" {
  domain_name = "example.com"
  rule_type   = "FORWARD"
  name        = %[1]q

  resolver_endpoint_id = "${aws_route53_resolver_endpoint.bar.id}"

  target_ip {
    ip = "192.0.2.7"
  }
}

resource "aws_route53_resolver_rule" "system" {
  domain_name = "example.org"
  rule_type   = "SYSTEM"
  name        = %[2]q
}

data "aws_route53_resolver_rules" "by_resolver_endpoint_id" {
  owner_id             = "${aws_route53_resolver_rule.system.owner_id}"
  resolver_endpoint_id = "${aws_route53_resolver_rule.forward.resolver_endpoint_id}"
}

data "aws_route53_resolver_rules" "by_rule_type" {
  owner_id  = "${aws_route53_resolver_rule.forward.owner_id}"
  rule_type = "${aws_route53_resolver_rule.system.rule_type}"
}

data "aws_route53_resolver_rules" "by_share_status" {
  owner_id     = "${aws_route53_resolver_rule.forward.owner_id}"
  share_status = "${aws_route53_resolver_rule.system.share_status}"
}

data "aws_route53_resolver_rules" "by_invalid_owner_id" {
  owner_id     = "000000000000"
  share_status = "SHARED_WITH_ME"
}
`, rName1, rName2)
}

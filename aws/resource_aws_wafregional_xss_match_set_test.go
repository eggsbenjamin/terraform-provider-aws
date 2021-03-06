package aws

import (
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/waf"
	"github.com/aws/aws-sdk-go/service/wafregional"
	"github.com/hashicorp/errwrap"
	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccAWSWafRegionalXssMatchSet_basic(t *testing.T) {
	var v waf.XssMatchSet
	xssMatchSet := fmt.Sprintf("xssMatchSet-%s", acctest.RandString(5))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAWSWafRegionalXssMatchSetDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccAWSWafRegionalXssMatchSetConfig(xssMatchSet),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAWSWafRegionalXssMatchSetExists("aws_wafregional_xss_match_set.xss_match_set", &v),
					resource.TestCheckResourceAttr(
						"aws_wafregional_xss_match_set.xss_match_set", "name", xssMatchSet),
					resource.TestCheckResourceAttr(
						"aws_wafregional_xss_match_set.xss_match_set", "xss_match_tuple.#", "2"),
					resource.TestCheckResourceAttr(
						"aws_wafregional_xss_match_set.xss_match_set", "xss_match_tuple.2018581549.field_to_match.#", "1"),
					resource.TestCheckResourceAttr(
						"aws_wafregional_xss_match_set.xss_match_set", "xss_match_tuple.2018581549.field_to_match.2316364334.data", ""),
					resource.TestCheckResourceAttr(
						"aws_wafregional_xss_match_set.xss_match_set", "xss_match_tuple.2018581549.field_to_match.2316364334.type", "QUERY_STRING"),
					resource.TestCheckResourceAttr(
						"aws_wafregional_xss_match_set.xss_match_set", "xss_match_tuple.2018581549.text_transformation", "NONE"),
					resource.TestCheckResourceAttr(
						"aws_wafregional_xss_match_set.xss_match_set", "xss_match_tuple.2786024938.field_to_match.#", "1"),
					resource.TestCheckResourceAttr(
						"aws_wafregional_xss_match_set.xss_match_set", "xss_match_tuple.2786024938.field_to_match.3756326843.data", ""),
					resource.TestCheckResourceAttr(
						"aws_wafregional_xss_match_set.xss_match_set", "xss_match_tuple.2786024938.field_to_match.3756326843.type", "URI"),
					resource.TestCheckResourceAttr(
						"aws_wafregional_xss_match_set.xss_match_set", "xss_match_tuple.2786024938.text_transformation", "NONE"),
				),
			},
		},
	})
}

func TestAccAWSWafRegionalXssMatchSet_changeNameForceNew(t *testing.T) {
	var before, after waf.XssMatchSet
	xssMatchSet := fmt.Sprintf("xssMatchSet-%s", acctest.RandString(5))
	xssMatchSetNewName := fmt.Sprintf("xssMatchSetNewName-%s", acctest.RandString(5))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAWSWafRegionalXssMatchSetDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAWSWafRegionalXssMatchSetConfig(xssMatchSet),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAWSWafRegionalXssMatchSetExists("aws_wafregional_xss_match_set.xss_match_set", &before),
					resource.TestCheckResourceAttr(
						"aws_wafregional_xss_match_set.xss_match_set", "name", xssMatchSet),
					resource.TestCheckResourceAttr(
						"aws_wafregional_xss_match_set.xss_match_set", "xss_match_tuple.#", "2"),
				),
			},
			{
				Config: testAccAWSWafRegionalXssMatchSetConfigChangeName(xssMatchSetNewName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAWSWafRegionalXssMatchSetExists("aws_wafregional_xss_match_set.xss_match_set", &after),
					resource.TestCheckResourceAttr(
						"aws_wafregional_xss_match_set.xss_match_set", "name", xssMatchSetNewName),
					resource.TestCheckResourceAttr(
						"aws_wafregional_xss_match_set.xss_match_set", "xss_match_tuple.#", "2"),
				),
			},
		},
	})
}

func TestAccAWSWafRegionalXssMatchSet_disappears(t *testing.T) {
	var v waf.XssMatchSet
	xssMatchSet := fmt.Sprintf("xssMatchSet-%s", acctest.RandString(5))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAWSWafRegionalXssMatchSetDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAWSWafRegionalXssMatchSetConfig(xssMatchSet),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAWSWafRegionalXssMatchSetExists("aws_wafregional_xss_match_set.xss_match_set", &v),
					testAccCheckAWSWafRegionalXssMatchSetDisappears(&v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccAWSWafRegionalXssMatchSet_changeTuples(t *testing.T) {
	var before, after waf.XssMatchSet
	setName := fmt.Sprintf("xssMatchSet-%s", acctest.RandString(5))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAWSWafRegionalXssMatchSetDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAWSWafRegionalXssMatchSetConfig(setName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckAWSWafRegionalXssMatchSetExists("aws_wafregional_xss_match_set.xss_match_set", &before),
					resource.TestCheckResourceAttr(
						"aws_wafregional_xss_match_set.xss_match_set", "name", setName),
					resource.TestCheckResourceAttr(
						"aws_wafregional_xss_match_set.xss_match_set", "xss_match_tuple.#", "2"),
					resource.TestCheckResourceAttr(
						"aws_wafregional_xss_match_set.xss_match_set", "xss_match_tuple.2018581549.field_to_match.#", "1"),
					resource.TestCheckResourceAttr(
						"aws_wafregional_xss_match_set.xss_match_set", "xss_match_tuple.2018581549.field_to_match.2316364334.data", ""),
					resource.TestCheckResourceAttr(
						"aws_wafregional_xss_match_set.xss_match_set", "xss_match_tuple.2018581549.field_to_match.2316364334.type", "QUERY_STRING"),
					resource.TestCheckResourceAttr(
						"aws_wafregional_xss_match_set.xss_match_set", "xss_match_tuple.2018581549.text_transformation", "NONE"),
					resource.TestCheckResourceAttr(
						"aws_wafregional_xss_match_set.xss_match_set", "xss_match_tuple.2786024938.field_to_match.#", "1"),
					resource.TestCheckResourceAttr(
						"aws_wafregional_xss_match_set.xss_match_set", "xss_match_tuple.2786024938.field_to_match.3756326843.data", ""),
					resource.TestCheckResourceAttr(
						"aws_wafregional_xss_match_set.xss_match_set", "xss_match_tuple.2786024938.field_to_match.3756326843.type", "URI"),
					resource.TestCheckResourceAttr(
						"aws_wafregional_xss_match_set.xss_match_set", "xss_match_tuple.2786024938.text_transformation", "NONE"),
				),
			},
			{
				Config: testAccAWSWafRegionalXssMatchSetConfig_changeTuples(setName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckAWSWafRegionalXssMatchSetExists("aws_wafregional_xss_match_set.xss_match_set", &after),
					resource.TestCheckResourceAttr(
						"aws_wafregional_xss_match_set.xss_match_set", "name", setName),
					resource.TestCheckResourceAttr(
						"aws_wafregional_xss_match_set.xss_match_set", "xss_match_tuple.#", "2"),
					resource.TestCheckResourceAttr(
						"aws_wafregional_xss_match_set.xss_match_set", "xss_match_tuple.2893682529.field_to_match.#", "1"),
					resource.TestCheckResourceAttr(
						"aws_wafregional_xss_match_set.xss_match_set", "xss_match_tuple.2893682529.field_to_match.4253810390.data", "GET"),
					resource.TestCheckResourceAttr(
						"aws_wafregional_xss_match_set.xss_match_set", "xss_match_tuple.2893682529.field_to_match.4253810390.type", "METHOD"),
					resource.TestCheckResourceAttr(
						"aws_wafregional_xss_match_set.xss_match_set", "xss_match_tuple.2893682529.text_transformation", "HTML_ENTITY_DECODE"),
					resource.TestCheckResourceAttr(
						"aws_wafregional_xss_match_set.xss_match_set", "xss_match_tuple.4270311415.field_to_match.#", "1"),
					resource.TestCheckResourceAttr(
						"aws_wafregional_xss_match_set.xss_match_set", "xss_match_tuple.4270311415.field_to_match.281401076.data", ""),
					resource.TestCheckResourceAttr(
						"aws_wafregional_xss_match_set.xss_match_set", "xss_match_tuple.4270311415.field_to_match.281401076.type", "BODY"),
					resource.TestCheckResourceAttr(
						"aws_wafregional_xss_match_set.xss_match_set", "xss_match_tuple.4270311415.text_transformation", "CMD_LINE"),
				),
			},
		},
	})
}

func TestAccAWSWafRegionalXssMatchSet_noTuples(t *testing.T) {
	var ipset waf.XssMatchSet
	setName := fmt.Sprintf("xssMatchSet-%s", acctest.RandString(5))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAWSWafRegionalXssMatchSetDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAWSWafRegionalXssMatchSetConfig_noTuples(setName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckAWSWafRegionalXssMatchSetExists("aws_wafregional_xss_match_set.xss_match_set", &ipset),
					resource.TestCheckResourceAttr(
						"aws_wafregional_xss_match_set.xss_match_set", "name", setName),
					resource.TestCheckResourceAttr(
						"aws_wafregional_xss_match_set.xss_match_set", "xss_match_tuple.#", "0"),
				),
			},
		},
	})
}

func testAccCheckAWSWafRegionalXssMatchSetDisappears(v *waf.XssMatchSet) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		conn := testAccProvider.Meta().(*AWSClient).wafregionalconn
		region := testAccProvider.Meta().(*AWSClient).region

		wr := newWafRegionalRetryer(conn, region)
		_, err := wr.RetryWithToken(func(token *string) (interface{}, error) {
			req := &waf.UpdateXssMatchSetInput{
				ChangeToken:   token,
				XssMatchSetId: v.XssMatchSetId,
			}

			for _, xssMatchTuple := range v.XssMatchTuples {
				xssMatchTupleUpdate := &waf.XssMatchSetUpdate{
					Action: aws.String("DELETE"),
					XssMatchTuple: &waf.XssMatchTuple{
						FieldToMatch:       xssMatchTuple.FieldToMatch,
						TextTransformation: xssMatchTuple.TextTransformation,
					},
				}
				req.Updates = append(req.Updates, xssMatchTupleUpdate)
			}
			return conn.UpdateXssMatchSet(req)
		})
		if err != nil {
			return errwrap.Wrapf("[ERROR] Error updating regional WAF XSS Match Set: {{err}}", err)
		}

		_, err = wr.RetryWithToken(func(token *string) (interface{}, error) {
			opts := &waf.DeleteXssMatchSetInput{
				ChangeToken:   token,
				XssMatchSetId: v.XssMatchSetId,
			}
			return conn.DeleteXssMatchSet(opts)
		})
		if err != nil {
			return errwrap.Wrapf("[ERROR] Error deleting regional WAF XSS Match Set: {{err}}", err)
		}
		return nil
	}
}

func testAccCheckAWSWafRegionalXssMatchSetExists(n string, v *waf.XssMatchSet) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No regional WAF XSS Match Set ID is set")
		}

		conn := testAccProvider.Meta().(*AWSClient).wafregionalconn
		resp, err := conn.GetXssMatchSet(&waf.GetXssMatchSetInput{
			XssMatchSetId: aws.String(rs.Primary.ID),
		})

		if err != nil {
			return err
		}

		if *resp.XssMatchSet.XssMatchSetId == rs.Primary.ID {
			*v = *resp.XssMatchSet
			return nil
		}

		return fmt.Errorf("Regional WAF XSS Match Set (%s) not found", rs.Primary.ID)
	}
}

func testAccCheckAWSWafRegionalXssMatchSetDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "aws_wafregional_xss_match_set" {
			continue
		}

		conn := testAccProvider.Meta().(*AWSClient).wafregionalconn
		resp, err := conn.GetXssMatchSet(
			&waf.GetXssMatchSetInput{
				XssMatchSetId: aws.String(rs.Primary.ID),
			})

		if err == nil {
			if *resp.XssMatchSet.XssMatchSetId == rs.Primary.ID {
				return fmt.Errorf("Regional WAF XSS Match Set %s still exists", rs.Primary.ID)
			}
		}

		// Return nil if the regional WAF XSS Match Set is already destroyed
		if isAWSErr(err, wafregional.ErrCodeWAFNonexistentItemException, "") {
			return nil
		}

		return err
	}

	return nil
}

func testAccAWSWafRegionalXssMatchSetConfig(name string) string {
	return fmt.Sprintf(`
resource "aws_wafregional_xss_match_set" "xss_match_set" {
  name = "%s"
  xss_match_tuple {
    text_transformation = "NONE"
    field_to_match {
      type = "URI"
    }
  }

  xss_match_tuple {
    text_transformation = "NONE"
    field_to_match {
      type = "QUERY_STRING"
    }
  }
}`, name)
}

func testAccAWSWafRegionalXssMatchSetConfigChangeName(name string) string {
	return fmt.Sprintf(`
resource "aws_wafregional_xss_match_set" "xss_match_set" {
  name = "%s"
  xss_match_tuple {
    text_transformation = "NONE"
    field_to_match {
      type = "URI"
    }
  }

  xss_match_tuple {
    text_transformation = "NONE"
    field_to_match {
      type = "QUERY_STRING"
    }
  }
}`, name)
}

func testAccAWSWafRegionalXssMatchSetConfig_changeTuples(name string) string {
	return fmt.Sprintf(`
resource "aws_wafregional_xss_match_set" "xss_match_set" {
  name = "%s"
  xss_match_tuple {
    text_transformation = "CMD_LINE"
    field_to_match {
      type = "BODY"
    }
  }
  xss_match_tuple {
    text_transformation = "HTML_ENTITY_DECODE"
    field_to_match {
      type = "METHOD"
      data = "GET"
    }
  }
}`, name)
}

func testAccAWSWafRegionalXssMatchSetConfig_noTuples(name string) string {
	return fmt.Sprintf(`
resource "aws_wafregional_xss_match_set" "xss_match_set" {
  name = "%s"
}`, name)
}

package ncloud

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"testing"
)

func TestAccDataSourceNcloudRegionsBasic(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceNcloudRegionsConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNcloudRegionsDataSourceID("data.ncloud_regions.regions"),
				),
			},
		},
	})
}

func testAccCheckNcloudRegionsDataSourceID(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("can't find Region data source: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("region data source ID not set")
		}
		return nil
	}
}

var testAccDataSourceNcloudRegionsConfig = `
data "ncloud_regions" "regions" {}
`
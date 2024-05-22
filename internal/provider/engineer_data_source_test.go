package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccEngineersDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			// Depends on present existing resources.  How can this be optimized?
			{
				Config: providerConfig + `data "devops-bootcamp_engineer" "test" {}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify number of coffees returned
					resource.TestCheckResourceAttr("data.devops-bootcamp_engineer.test", "engineer.#", "7"),
					// Verify the first coffee to ensure all attributes are set
					resource.TestCheckResourceAttr("data.devops-bootcamp_engineer.test", "engineer.0.id", "4FFLH"),
					resource.TestCheckResourceAttr("data.devops-bootcamp_engineer.test", "engineer.0.name", "eve"),
					resource.TestCheckResourceAttr("data.devops-bootcamp_engineer.test", "engineer.0.email", "eve@finches.com"),
				),
			},
		},
	})
}

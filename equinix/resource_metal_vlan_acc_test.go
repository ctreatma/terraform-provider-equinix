package equinix

import (
	"fmt"
	"testing"

	"github.com/equinix/terraform-provider-equinix/internal/config"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/packethost/packngo"
)

func testAccCheckMetalVlanConfig_metro(projSuffix, metro, desc string) string {
	return fmt.Sprintf(`
resource "equinix_metal_project" "foobar" {
    name = "tfacc-vlan-%s"
}

resource "equinix_metal_vlan" "foovlan" {
    project_id = equinix_metal_project.foobar.id
    metro = "%s"
    description = "%s"
    vxlan = 5
}
`, projSuffix, metro, desc)
}

func TestAccMetalVlan_metro(t *testing.T) {
	rs := acctest.RandString(10)
	metro := "sv"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ExternalProviders:        testExternalProviders,
		ProtoV5ProviderFactories: testAccProtoV5ProviderFactories,
		CheckDestroy:             testAccMetalVlanCheckDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckMetalVlanConfig_metro(rs, metro, "tfacc-vlan"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"equinix_metal_vlan.foovlan", "metro", metro),
					resource.TestCheckResourceAttr(
						"equinix_metal_vlan.foovlan", "facility", ""),
				),
			},
		},
	})
}

func TestAccMetalVlan_basic(t *testing.T) {
	var vlan packngo.VirtualNetwork
	rs := acctest.RandString(10)
	fac := "ny5"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ExternalProviders:        testExternalProviders,
		ProtoV5ProviderFactories: testAccProtoV5ProviderFactories,
		CheckDestroy:             testAccMetalVlanCheckDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccMetalVlanConfig_var(rs, fac, "tfacc-vlan"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMetalVlanExists("equinix_metal_vlan.foovlan", &vlan),
					resource.TestCheckResourceAttr(
						"equinix_metal_vlan.foovlan", "description", "tfacc-vlan"),
					resource.TestCheckResourceAttr(
						"equinix_metal_vlan.foovlan", "facility", fac),
				),
			},
		},
	})
}

func testAccCheckMetalVlanExists(n string, vlan *packngo.VirtualNetwork) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("No Record ID is set")
		}

		client := testAccProvider.Meta().(*config.Config).Metal

		foundVlan, _, err := client.ProjectVirtualNetworks.Get(rs.Primary.ID, nil)
		if err != nil {
			return err
		}
		if foundVlan.ID != rs.Primary.ID {
			return fmt.Errorf("Record not found: %v - %v", rs.Primary.ID, foundVlan)
		}

		*vlan = *foundVlan

		return nil
	}
}

func testAccMetalVlanCheckDestroyed(s *terraform.State) error {
	client := testAccProvider.Meta().(*config.Config).Metal

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "equinix_metal_vlan" {
			continue
		}
		if _, _, err := client.ProjectVirtualNetworks.Get(rs.Primary.ID, nil); err == nil {
			return fmt.Errorf("Metal Vlan still exists")
		}
	}

	return nil
}

func testAccMetalVlanConfig_var(projSuffix, facility, desc string) string {
	return fmt.Sprintf(`
resource "equinix_metal_project" "foobar" {
    name = "tfacc-vlan-%s"
}

resource "equinix_metal_vlan" "foovlan" {
    project_id = "${equinix_metal_project.foobar.id}"
    facility = "%s"
    description = "%s"
}
`, projSuffix, facility, desc)
}

func TestAccMetalVlan_importBasic(t *testing.T) {
	rs := acctest.RandString(10)
	fac := "ny5"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ExternalProviders:        testExternalProviders,
		ProtoV5ProviderFactories: testAccProtoV5ProviderFactories,
		CheckDestroy:             testAccMetalVlanCheckDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccMetalVlanConfig_var(rs, fac, "tfacc-vlan"),
			},
			{
				ResourceName:      "equinix_metal_vlan.foovlan",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}
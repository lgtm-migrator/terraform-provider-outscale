package outscale

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	oscgo "github.com/outscale/osc-sdk-go/v2"
	"github.com/spf13/cast"
)

func TestAccOutscaleClientGateway_basic(t *testing.T) {
	t.Parallel()
	resourceName := "outscale_client_gateway.foo"
	rBgpAsn := acctest.RandIntRange(64512, 65534)
	rBgpAsnUpdated := acctest.RandIntRange(64512, 65534)

	resource.Test(t, resource.TestCase{
		PreCheck:      func() { testAccPreCheck(t) },
		IDRefreshName: resourceName,
		Providers:     testAccProviders,
		CheckDestroy:  testAccCheckClientGatewayDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccClientGatewayConfig(rBgpAsn),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckClientGatewayExists(resourceName),
					resource.TestCheckResourceAttrSet(resourceName, "bgp_asn"),
					resource.TestCheckResourceAttrSet(resourceName, "public_ip"),
					resource.TestCheckResourceAttrSet(resourceName, "connection_type"),

					resource.TestCheckResourceAttr(resourceName, "bgp_asn", cast.ToString(rBgpAsn)),
				),
			},
			{
				Config: testAccClientGatewayConfig(rBgpAsnUpdated),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckClientGatewayExists(resourceName),
					resource.TestCheckResourceAttrSet(resourceName, "bgp_asn"),
					resource.TestCheckResourceAttrSet(resourceName, "public_ip"),
					resource.TestCheckResourceAttrSet(resourceName, "connection_type"),

					resource.TestCheckResourceAttr(resourceName, "bgp_asn", cast.ToString(rBgpAsnUpdated)),
				),
			},
		},
	})
}

func TestAccOutscaleClientGateway_withTags(t *testing.T) {
	resourceName := "outscale_client_gateway.foo"
	value := fmt.Sprintf("testacc-%s", acctest.RandString(5))
	valueUpdated := fmt.Sprintf("testacc-%s", acctest.RandString(5))

	resource.Test(t, resource.TestCase{
		PreCheck:      func() { testAccPreCheck(t) },
		IDRefreshName: resourceName,
		Providers:     testAccProviders,
		CheckDestroy:  testAccCheckClientGatewayDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccClientGatewayConfigWithTags(value),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckClientGatewayExists(resourceName),
					resource.TestCheckResourceAttrSet(resourceName, "bgp_asn"),
					resource.TestCheckResourceAttrSet(resourceName, "public_ip"),
					resource.TestCheckResourceAttrSet(resourceName, "connection_type"),
				),
			},
			{
				Config: testAccClientGatewayConfigWithTags(valueUpdated),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckClientGatewayExists(resourceName),
					resource.TestCheckResourceAttrSet(resourceName, "bgp_asn"),
					resource.TestCheckResourceAttrSet(resourceName, "public_ip"),
					resource.TestCheckResourceAttrSet(resourceName, "connection_type"),
				),
			},
		},
	})
}

func testAccCheckClientGatewayExists(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("Not found: %s", resourceName)
		}

		conn := testAccProvider.Meta().(*OutscaleClient).OSCAPI

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Client Gateway ID is set")
		}

		filter := oscgo.ReadClientGatewaysRequest{
			Filters: &oscgo.FiltersClientGateway{
				ClientGatewayIds: &[]string{rs.Primary.ID},
			},
		}

		resp, _, err := conn.ClientGatewayApi.ReadClientGateways(context.Background()).ReadClientGatewaysRequest(filter).Execute()
		if err != nil || len(resp.GetClientGateways()) < 1 {
			return fmt.Errorf("Outscale Client Gateway not found (%s)", rs.Primary.ID)
		}
		return nil
	}
}

func testAccCheckClientGatewayDestroy(s *terraform.State) error {
	conn := testAccProvider.Meta().(*OutscaleClient).OSCAPI

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "outscale_client_gateway" {
			continue
		}

		filter := oscgo.ReadClientGatewaysRequest{
			Filters: &oscgo.FiltersClientGateway{
				ClientGatewayIds: &[]string{rs.Primary.ID},
			},
		}

		resp, _, err := conn.ClientGatewayApi.ReadClientGateways(context.Background()).ReadClientGatewaysRequest(filter).Execute()
		if err != nil ||
			len(resp.GetClientGateways()) > 0 && resp.GetClientGateways()[0].GetState() != "deleted" {
			return fmt.Errorf("Outscale Client Gateway still exists (%s): %s", rs.Primary.ID, err)
		}
	}
	return nil
}

func testAccClientGatewayConfig(rBgpAsn int) string {
	return fmt.Sprintf(`
		resource "outscale_client_gateway" "foo" {
			bgp_asn         = %d
			public_ip       = "172.0.0.1"
			connection_type = "ipsec.1"
		}
	`, rBgpAsn)
}

func testAccClientGatewayConfigWithTags(value string) string {
	return fmt.Sprintf(`
		resource "outscale_client_gateway" "foo" {
			bgp_asn         = 3
			public_ip       = "172.0.0.1"
			connection_type = "ipsec.1"

			tags {
				key = "Name"
				value = "%s"
			}
		}
	`, value)
}

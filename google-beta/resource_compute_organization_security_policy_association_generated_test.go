// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"log"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccComputeOrganizationSecurityPolicyAssociation_organizationSecurityPolicyAssociationBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"org_id":        getTestOrgFromEnv(t),
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProvidersOiCS,
		CheckDestroy: testAccCheckComputeOrganizationSecurityPolicyAssociationDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeOrganizationSecurityPolicyAssociation_organizationSecurityPolicyAssociationBasicExample(context),
			},
		},
	})
}

func testAccComputeOrganizationSecurityPolicyAssociation_organizationSecurityPolicyAssociationBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_organization_security_policy" "policy" {
  provider = google-beta

  display_name = "tf-test%{random_suffix}"
  parent       = "organizations/%{org_id}"
}

resource "google_compute_organization_security_policy_rule" "policy" {
  provider = google-beta

  policy_id = google_compute_organization_security_policy.policy.id
  action = "allow"

  direction = "INGRESS"
  enable_logging = true
  match {
    config {
      src_ip_ranges = ["192.168.0.0/16", "10.0.0.0/8"]
      layer4_config {
        ip_protocol = "tcp"
        ports = ["22"]
      }
      layer4_config {
        ip_protocol = "icmp"
      }
    }
  }
  priority = 100
}

resource "google_compute_organization_security_policy_association" "policy" {
  provider = google-beta

  name          = "tf-test%{random_suffix}"
  attachment_id = google_compute_organization_security_policy.policy.parent
  policy_id     = google_compute_organization_security_policy.policy.id
}
`, context)
}

func testAccCheckComputeOrganizationSecurityPolicyAssociationDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_compute_organization_security_policy_association" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			log.Printf("[DEBUG] Ignoring destroy during test")
		}

		return nil
	}
}

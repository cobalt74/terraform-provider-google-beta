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
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccSecurityScannerScanConfig_scanConfigBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProvidersOiCS,
		CheckDestroy: testAccCheckSecurityScannerScanConfigDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccSecurityScannerScanConfig_scanConfigBasicExample(context),
			},
		},
	})
}

func testAccSecurityScannerScanConfig_scanConfigBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_address" "scanner_static_ip" {
  provider = google-beta
  name     = "tf-test-scan-basic-static-ip%{random_suffix}"
}

resource "google_security_scanner_scan_config" "scan-config" {
  provider         = google-beta
  display_name     = "tf-test-terraform-scan-config%{random_suffix}"
  starting_urls    = ["http://${google_compute_address.scanner_static_ip.address}"]
  target_platforms = ["COMPUTE"]
}

provider "google-beta" {
  region = "us-central1"
  zone   = "us-central1-a"
}
`, context)
}

func testAccCheckSecurityScannerScanConfigDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_security_scanner_scan_config" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := googleProviderConfig(t)

			url, err := replaceVarsForTest(config, rs, "{{SecurityScannerBasePath}}{{name}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = sendRequest(config, "GET", billingProject, url, config.userAgent, nil)
			if err == nil {
				return fmt.Errorf("SecurityScannerScanConfig still exists at %s", url)
			}
		}

		return nil
	}
}

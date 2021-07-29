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

func TestAccCloudAssetOrganizationFeed_cloudAssetOrganizationFeedExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project":       getTestProjectFromEnv(),
		"org_id":        getTestOrgFromEnv(t),
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudAssetOrganizationFeedDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccCloudAssetOrganizationFeed_cloudAssetOrganizationFeedExample(context),
			},
			{
				ResourceName:            "google_cloud_asset_organization_feed.organization_feed",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"billing_project", "feed_id", "org_id"},
			},
		},
	})
}

func testAccCloudAssetOrganizationFeed_cloudAssetOrganizationFeedExample(context map[string]interface{}) string {
	return Nprintf(`
# Create a feed that sends notifications about network resource updates under a
# particular organization.
resource "google_cloud_asset_organization_feed" "organization_feed" {
  billing_project = "%{project}"
  org_id          = "%{org_id}"
  feed_id         = "tf-test-network-updates%{random_suffix}"
  content_type    = "RESOURCE"

  asset_types = [
    "compute.googleapis.com/Subnetwork",
    "compute.googleapis.com/Network",
  ]

  feed_output_config {
    pubsub_destination {
      topic = google_pubsub_topic.feed_output.id
    }
  }

  condition {
    expression = <<-EOT
    !temporal_asset.deleted &&
    temporal_asset.prior_asset_state == google.cloud.asset.v1.TemporalAsset.PriorAssetState.DOES_NOT_EXIST
    EOT
    title = "created"
    description = "Send notifications on creation events"
  }

  # Wait for the permission to be ready on the destination topic.
  depends_on = [
    google_pubsub_topic_iam_member.cloud_asset_writer,
  ]
}

# The topic where the resource change notifications will be sent.
resource "google_pubsub_topic" "feed_output" {
  project  = "%{project}"
  name     = "tf-test-network-updates%{random_suffix}"
}

# Find the project number of the project whose identity will be used for sending
# the asset change notifications.
data "google_project" "project" {
  project_id = "%{project}"
}

# Allow the publishing role to the Cloud Asset service account of the project that
# was used for sending the notifications.
resource "google_pubsub_topic_iam_member" "cloud_asset_writer" {
  project = "%{project}"
  topic   = google_pubsub_topic.feed_output.id
  role    = "roles/pubsub.publisher"
  member  = "serviceAccount:service-${data.google_project.project.number}@gcp-sa-cloudasset.iam.gserviceaccount.com"
}
`, context)
}

func testAccCheckCloudAssetOrganizationFeedDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_cloud_asset_organization_feed" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := googleProviderConfig(t)

			url, err := replaceVarsForTest(config, rs, "{{CloudAssetBasePath}}{{name}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = sendRequest(config, "GET", billingProject, url, config.userAgent, nil)
			if err == nil {
				return fmt.Errorf("CloudAssetOrganizationFeed still exists at %s", url)
			}
		}

		return nil
	}
}

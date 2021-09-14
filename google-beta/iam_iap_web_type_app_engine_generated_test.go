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
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIapWebTypeAppEngineIamBindingGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
		"role":          "roles/iap.httpsResourceAccessor",
		"project_id":    fmt.Sprintf("tf-test%s", randString(t, 10)),
		"org_id":        getTestOrgFromEnv(t),

		"condition_title": "expires_after_2019_12_31",
		"condition_expr":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		ExternalProviders: map[string]resource.ExternalProvider{
			"random": {},
			"time":   {},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccIapWebTypeAppEngineIamBinding_basicGenerated(context),
			},
			{
				ResourceName:      "google_iap_web_type_app_engine_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/iap_web/appengine-%s roles/iap.httpsResourceAccessor", context["project_id"], context["project_id"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				// Test Iam Binding update
				Config: testAccIapWebTypeAppEngineIamBinding_updateGenerated(context),
			},
			{
				ResourceName:      "google_iap_web_type_app_engine_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/iap_web/appengine-%s roles/iap.httpsResourceAccessor", context["project_id"], context["project_id"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccIapWebTypeAppEngineIamMemberGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
		"role":          "roles/iap.httpsResourceAccessor",
		"project_id":    fmt.Sprintf("tf-test%s", randString(t, 10)),
		"org_id":        getTestOrgFromEnv(t),

		"condition_title": "expires_after_2019_12_31",
		"condition_expr":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		ExternalProviders: map[string]resource.ExternalProvider{
			"random": {},
			"time":   {},
		},
		Steps: []resource.TestStep{
			{
				// Test Iam Member creation (no update for member, no need to test)
				Config: testAccIapWebTypeAppEngineIamMember_basicGenerated(context),
			},
			{
				ResourceName:      "google_iap_web_type_app_engine_iam_member.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/iap_web/appengine-%s roles/iap.httpsResourceAccessor user:admin@hashicorptest.com", context["project_id"], context["project_id"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccIapWebTypeAppEngineIamPolicyGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
		"role":          "roles/iap.httpsResourceAccessor",
		"project_id":    fmt.Sprintf("tf-test%s", randString(t, 10)),
		"org_id":        getTestOrgFromEnv(t),

		"condition_title": "expires_after_2019_12_31",
		"condition_expr":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		ExternalProviders: map[string]resource.ExternalProvider{
			"random": {},
			"time":   {},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccIapWebTypeAppEngineIamPolicy_basicGenerated(context),
			},
			{
				ResourceName:      "google_iap_web_type_app_engine_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/iap_web/appengine-%s", context["project_id"], context["project_id"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccIapWebTypeAppEngineIamPolicy_emptyBinding(context),
			},
			{
				ResourceName:      "google_iap_web_type_app_engine_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/iap_web/appengine-%s", context["project_id"], context["project_id"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccIapWebTypeAppEngineIamBindingGenerated_withCondition(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
		"role":          "roles/iap.httpsResourceAccessor",
		"project_id":    fmt.Sprintf("tf-test%s", randString(t, 10)),
		"org_id":        getTestOrgFromEnv(t),

		"condition_title": "expires_after_2019_12_31",
		"condition_expr":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		ExternalProviders: map[string]resource.ExternalProvider{
			"random": {},
			"time":   {},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccIapWebTypeAppEngineIamBinding_withConditionGenerated(context),
			},
			{
				ResourceName:      "google_iap_web_type_app_engine_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/iap_web/appengine-%s roles/iap.httpsResourceAccessor %s", context["project_id"], context["project_id"], context["condition_title"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccIapWebTypeAppEngineIamBindingGenerated_withAndWithoutCondition(t *testing.T) {
	// Multiple fine-grained resources
	skipIfVcr(t)
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
		"role":          "roles/iap.httpsResourceAccessor",
		"project_id":    fmt.Sprintf("tf-test%s", randString(t, 10)),
		"org_id":        getTestOrgFromEnv(t),

		"condition_title": "expires_after_2019_12_31",
		"condition_expr":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		ExternalProviders: map[string]resource.ExternalProvider{
			"random": {},
			"time":   {},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccIapWebTypeAppEngineIamBinding_withAndWithoutConditionGenerated(context),
			},
			{
				ResourceName:      "google_iap_web_type_app_engine_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/iap_web/appengine-%s roles/iap.httpsResourceAccessor", context["project_id"], context["project_id"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				ResourceName:      "google_iap_web_type_app_engine_iam_binding.foo2",
				ImportStateId:     fmt.Sprintf("projects/%s/iap_web/appengine-%s roles/iap.httpsResourceAccessor %s", context["project_id"], context["project_id"], context["condition_title"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccIapWebTypeAppEngineIamMemberGenerated_withCondition(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
		"role":          "roles/iap.httpsResourceAccessor",
		"project_id":    fmt.Sprintf("tf-test%s", randString(t, 10)),
		"org_id":        getTestOrgFromEnv(t),

		"condition_title": "expires_after_2019_12_31",
		"condition_expr":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		ExternalProviders: map[string]resource.ExternalProvider{
			"random": {},
			"time":   {},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccIapWebTypeAppEngineIamMember_withConditionGenerated(context),
			},
			{
				ResourceName:      "google_iap_web_type_app_engine_iam_member.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/iap_web/appengine-%s roles/iap.httpsResourceAccessor user:admin@hashicorptest.com %s", context["project_id"], context["project_id"], context["condition_title"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccIapWebTypeAppEngineIamMemberGenerated_withAndWithoutCondition(t *testing.T) {
	// Multiple fine-grained resources
	skipIfVcr(t)
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
		"role":          "roles/iap.httpsResourceAccessor",
		"project_id":    fmt.Sprintf("tf-test%s", randString(t, 10)),
		"org_id":        getTestOrgFromEnv(t),

		"condition_title": "expires_after_2019_12_31",
		"condition_expr":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		ExternalProviders: map[string]resource.ExternalProvider{
			"random": {},
			"time":   {},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccIapWebTypeAppEngineIamMember_withAndWithoutConditionGenerated(context),
			},
			{
				ResourceName:      "google_iap_web_type_app_engine_iam_member.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/iap_web/appengine-%s roles/iap.httpsResourceAccessor user:admin@hashicorptest.com", context["project_id"], context["project_id"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				ResourceName:      "google_iap_web_type_app_engine_iam_member.foo2",
				ImportStateId:     fmt.Sprintf("projects/%s/iap_web/appengine-%s roles/iap.httpsResourceAccessor user:admin@hashicorptest.com %s", context["project_id"], context["project_id"], context["condition_title"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccIapWebTypeAppEngineIamPolicyGenerated_withCondition(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
		"role":          "roles/iap.httpsResourceAccessor",
		"project_id":    fmt.Sprintf("tf-test%s", randString(t, 10)),
		"org_id":        getTestOrgFromEnv(t),

		"condition_title": "expires_after_2019_12_31",
		"condition_expr":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		ExternalProviders: map[string]resource.ExternalProvider{
			"random": {},
			"time":   {},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccIapWebTypeAppEngineIamPolicy_withConditionGenerated(context),
			},
			{
				ResourceName:      "google_iap_web_type_app_engine_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/iap_web/appengine-%s", context["project_id"], context["project_id"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccIapWebTypeAppEngineIamMember_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_project" "project" {
  name       = "%{project_id}"
  project_id = "%{project_id}"
  org_id     = "%{org_id}"
}

resource "time_sleep" "wait_60_seconds" {
  depends_on = [google_project.project]

  create_duration = "60s"
}

resource "google_project_service" "project_service" {
  project = google_project.project.project_id
  service = "iap.googleapis.com"

  # Needed for CI tests for permissions to propagate, should not be needed for actual usage
  depends_on = [time_sleep.wait_60_seconds]
}


resource "google_app_engine_application" "app" {
  project     = google_project_service.project_service.project
  location_id = "us-central"
}

resource "google_iap_web_type_app_engine_iam_member" "foo" {
  project = google_app_engine_application.app.project
  app_id = google_app_engine_application.app.app_id
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
}
`, context)
}

func testAccIapWebTypeAppEngineIamPolicy_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_project" "project" {
  name       = "%{project_id}"
  project_id = "%{project_id}"
  org_id     = "%{org_id}"
}

resource "time_sleep" "wait_60_seconds" {
  depends_on = [google_project.project]

  create_duration = "60s"
}

resource "google_project_service" "project_service" {
  project = google_project.project.project_id
  service = "iap.googleapis.com"

  # Needed for CI tests for permissions to propagate, should not be needed for actual usage
  depends_on = [time_sleep.wait_60_seconds]
}


resource "google_app_engine_application" "app" {
  project     = google_project_service.project_service.project
  location_id = "us-central"
}

data "google_iam_policy" "foo" {
  binding {
    role = "%{role}"
    members = ["user:admin@hashicorptest.com"]
  }
}

resource "google_iap_web_type_app_engine_iam_policy" "foo" {
  project = google_app_engine_application.app.project
  app_id = google_app_engine_application.app.app_id
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}

func testAccIapWebTypeAppEngineIamPolicy_emptyBinding(context map[string]interface{}) string {
	return Nprintf(`
resource "google_project" "project" {
  name       = "%{project_id}"
  project_id = "%{project_id}"
  org_id     = "%{org_id}"
}

resource "time_sleep" "wait_60_seconds" {
  depends_on = [google_project.project]

  create_duration = "60s"
}

resource "google_project_service" "project_service" {
  project = google_project.project.project_id
  service = "iap.googleapis.com"

  # Needed for CI tests for permissions to propagate, should not be needed for actual usage
  depends_on = [time_sleep.wait_60_seconds]
}


resource "google_app_engine_application" "app" {
  project     = google_project_service.project_service.project
  location_id = "us-central"
}

data "google_iam_policy" "foo" {
}

resource "google_iap_web_type_app_engine_iam_policy" "foo" {
  project = google_app_engine_application.app.project
  app_id = google_app_engine_application.app.app_id
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}

func testAccIapWebTypeAppEngineIamBinding_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_project" "project" {
  name       = "%{project_id}"
  project_id = "%{project_id}"
  org_id     = "%{org_id}"
}

resource "time_sleep" "wait_60_seconds" {
  depends_on = [google_project.project]

  create_duration = "60s"
}

resource "google_project_service" "project_service" {
  project = google_project.project.project_id
  service = "iap.googleapis.com"

  # Needed for CI tests for permissions to propagate, should not be needed for actual usage
  depends_on = [time_sleep.wait_60_seconds]
}


resource "google_app_engine_application" "app" {
  project     = google_project_service.project_service.project
  location_id = "us-central"
}

resource "google_iap_web_type_app_engine_iam_binding" "foo" {
  project = google_app_engine_application.app.project
  app_id = google_app_engine_application.app.app_id
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
}
`, context)
}

func testAccIapWebTypeAppEngineIamBinding_updateGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_project" "project" {
  name       = "%{project_id}"
  project_id = "%{project_id}"
  org_id     = "%{org_id}"
}

resource "time_sleep" "wait_60_seconds" {
  depends_on = [google_project.project]

  create_duration = "60s"
}

resource "google_project_service" "project_service" {
  project = google_project.project.project_id
  service = "iap.googleapis.com"

  # Needed for CI tests for permissions to propagate, should not be needed for actual usage
  depends_on = [time_sleep.wait_60_seconds]
}


resource "google_app_engine_application" "app" {
  project     = google_project_service.project_service.project
  location_id = "us-central"
}

resource "google_iap_web_type_app_engine_iam_binding" "foo" {
  project = google_app_engine_application.app.project
  app_id = google_app_engine_application.app.app_id
  role = "%{role}"
  members = ["user:admin@hashicorptest.com", "user:paddy@hashicorp.com"]
}
`, context)
}

func testAccIapWebTypeAppEngineIamBinding_withConditionGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_project" "project" {
  name       = "%{project_id}"
  project_id = "%{project_id}"
  org_id     = "%{org_id}"
}

resource "time_sleep" "wait_60_seconds" {
  depends_on = [google_project.project]

  create_duration = "60s"
}

resource "google_project_service" "project_service" {
  project = google_project.project.project_id
  service = "iap.googleapis.com"

  # Needed for CI tests for permissions to propagate, should not be needed for actual usage
  depends_on = [time_sleep.wait_60_seconds]
}


resource "google_app_engine_application" "app" {
  project     = google_project_service.project_service.project
  location_id = "us-central"
}

resource "google_iap_web_type_app_engine_iam_binding" "foo" {
  project = google_app_engine_application.app.project
  app_id = google_app_engine_application.app.app_id
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
  condition {
    title       = "%{condition_title}"
    description = "Expiring at midnight of 2019-12-31"
    expression  = "%{condition_expr}"
  }
}
`, context)
}

func testAccIapWebTypeAppEngineIamBinding_withAndWithoutConditionGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_project" "project" {
  name       = "%{project_id}"
  project_id = "%{project_id}"
  org_id     = "%{org_id}"
}

resource "time_sleep" "wait_60_seconds" {
  depends_on = [google_project.project]

  create_duration = "60s"
}

resource "google_project_service" "project_service" {
  project = google_project.project.project_id
  service = "iap.googleapis.com"

  # Needed for CI tests for permissions to propagate, should not be needed for actual usage
  depends_on = [time_sleep.wait_60_seconds]
}


resource "google_app_engine_application" "app" {
  project     = google_project_service.project_service.project
  location_id = "us-central"
}

resource "google_iap_web_type_app_engine_iam_binding" "foo" {
  project = google_app_engine_application.app.project
  app_id = google_app_engine_application.app.app_id
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
}

resource "google_iap_web_type_app_engine_iam_binding" "foo2" {
  project = google_app_engine_application.app.project
  app_id = google_app_engine_application.app.app_id
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
  condition {
    title       = "%{condition_title}"
    description = "Expiring at midnight of 2019-12-31"
    expression  = "%{condition_expr}"
  }
}
`, context)
}

func testAccIapWebTypeAppEngineIamMember_withConditionGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_project" "project" {
  name       = "%{project_id}"
  project_id = "%{project_id}"
  org_id     = "%{org_id}"
}

resource "time_sleep" "wait_60_seconds" {
  depends_on = [google_project.project]

  create_duration = "60s"
}

resource "google_project_service" "project_service" {
  project = google_project.project.project_id
  service = "iap.googleapis.com"

  # Needed for CI tests for permissions to propagate, should not be needed for actual usage
  depends_on = [time_sleep.wait_60_seconds]
}


resource "google_app_engine_application" "app" {
  project     = google_project_service.project_service.project
  location_id = "us-central"
}

resource "google_iap_web_type_app_engine_iam_member" "foo" {
  project = google_app_engine_application.app.project
  app_id = google_app_engine_application.app.app_id
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
  condition {
    title       = "%{condition_title}"
    description = "Expiring at midnight of 2019-12-31"
    expression  = "%{condition_expr}"
  }
}
`, context)
}

func testAccIapWebTypeAppEngineIamMember_withAndWithoutConditionGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_project" "project" {
  name       = "%{project_id}"
  project_id = "%{project_id}"
  org_id     = "%{org_id}"
}

resource "time_sleep" "wait_60_seconds" {
  depends_on = [google_project.project]

  create_duration = "60s"
}

resource "google_project_service" "project_service" {
  project = google_project.project.project_id
  service = "iap.googleapis.com"

  # Needed for CI tests for permissions to propagate, should not be needed for actual usage
  depends_on = [time_sleep.wait_60_seconds]
}


resource "google_app_engine_application" "app" {
  project     = google_project_service.project_service.project
  location_id = "us-central"
}

resource "google_iap_web_type_app_engine_iam_member" "foo" {
  project = google_app_engine_application.app.project
  app_id = google_app_engine_application.app.app_id
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
}

resource "google_iap_web_type_app_engine_iam_member" "foo2" {
  project = google_app_engine_application.app.project
  app_id = google_app_engine_application.app.app_id
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
  condition {
    title       = "%{condition_title}"
    description = "Expiring at midnight of 2019-12-31"
    expression  = "%{condition_expr}"
  }
}
`, context)
}

func testAccIapWebTypeAppEngineIamPolicy_withConditionGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_project" "project" {
  name       = "%{project_id}"
  project_id = "%{project_id}"
  org_id     = "%{org_id}"
}

resource "time_sleep" "wait_60_seconds" {
  depends_on = [google_project.project]

  create_duration = "60s"
}

resource "google_project_service" "project_service" {
  project = google_project.project.project_id
  service = "iap.googleapis.com"

  # Needed for CI tests for permissions to propagate, should not be needed for actual usage
  depends_on = [time_sleep.wait_60_seconds]
}


resource "google_app_engine_application" "app" {
  project     = google_project_service.project_service.project
  location_id = "us-central"
}

data "google_iam_policy" "foo" {
  binding {
    role = "%{role}"
    members = ["user:admin@hashicorptest.com"]
    condition {
      title       = "%{condition_title}"
      description = "Expiring at midnight of 2019-12-31"
      expression  = "%{condition_expr}"
    }
  }
}

resource "google_iap_web_type_app_engine_iam_policy" "foo" {
  project = google_app_engine_application.app.project
  app_id = google_app_engine_application.app.app_id
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}

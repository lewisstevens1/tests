package test

import (
	"encoding/json"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

type identityProviderPermission struct {
	Effect   string
	Resource string
	Action   []string
}

func getIdentityProviderPermissions() string {
	var ipp []identityProviderPermission

	ipp = append(ipp, identityProviderPermission{
		Effect:   "*",
		Resource: "*",
		Action: []string{
			"s3:*",
		},
	})

	ipp_json, _ := json.Marshal(ipp)
	return string(ipp_json)
}

func TestOpenIdConnectApply(t *testing.T) {

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../",

		Vars: map[string]interface{}{
			"identity_provider_permissions": getIdentityProviderPermissions(),
		},
	})

	defer terraform.Destroy(t, terraformOptions)
	terraform.InitAndApply(t, terraformOptions)

	output := terraform.Output(t, terraformOptions, "complete")
	assert.Equal(t, "true", output)

}

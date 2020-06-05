package rancher2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

const (
	bootstrapDefaultTokenDesc   = "Terraform bootstrap admin token"
	bootstrapDefaultSessionDesc = "Terraform bootstrap admin session"
	bootstrapDefaultUser        = "admin"
	bootstrapDefaultPassword    = "admin"
	bootstrapDefaultTTL         = "60000"
	bootstrapSettingURL         = "server-url"
	bootstrapSettingTelemetry   = "telemetry-opt"
)

//Schemas

func bootstrapFields() map[string]*schema.Schema {
	s := map[string]*schema.Schema{
		"current_password": {
			Type:      schema.TypeString,
			Optional:  true,
			Computed:  true,
			Sensitive: true,
		},
		"password": {
			Type:      schema.TypeString,
			Optional:  true,
			Computed:  true,
			Sensitive: true,
		},
		"token_ttl": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  0,
		},
		"token": {
			Type:      schema.TypeString,
			Computed:  true,
			Sensitive: true,
		},
		"token_id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"token_update": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"telemetry": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"temp_token": {
			Type:      schema.TypeString,
			Computed:  true,
			Sensitive: true,
		},
		"temp_token_id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"url": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"user": {
			Type:     schema.TypeString,
			Computed: true,
		},
	}

	return s
}

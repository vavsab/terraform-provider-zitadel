package application_api

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/zitadel/terraform-provider-zitadel/zitadel/helper"
)

func GetDatasource() *schema.Resource {
	return &schema.Resource{
		Description: "Datasource representing an API application belonging to a project, with all configuration possibilities.",
		Schema: map[string]*schema.Schema{
			AppIDVar: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The ID of this resource.",
			},
			helper.OrgIDVar: helper.OrgIDDatasourceField,
			ProjectIDVar: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "ID of the project",
			},
			NameVar: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Name of the application",
			},
			authMethodTypeVar: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Auth method type",
			},
		},
		ReadContext: read,
	}
}

package servicenow

import (
	"github.com/coveooss/terraform-provider-servicenow/servicenow/client"
	"github.com/coveooss/terraform-provider-servicenow/servicenow/resources"
	"github.com/hashicorp/terraform/helper/schema"
)

// Provider is a Terraform Provider to that manages objects in a ServiceNow instance.
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"instance_url": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("SN_INSTANCE_URL", nil),
				Description: "The URL of the ServiceNow instance to connect to.",
			},
			"username": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("SN_USERNAME", nil),
				Description: "The username to use for authentication.",
			},
			"password": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("SN_PASSWORD", nil),
				Description: "The password to use for authentication.",
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"servicenow_application": resources.ResourceApplication(),
		},

		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(data *schema.ResourceData) (interface{}, error) {
	// Create a new client to talk to the instance.
	client := client.NewClient(
		data.Get("instance_url").(string),
		data.Get("username").(string),
		data.Get("password").(string))

	return client, nil
}

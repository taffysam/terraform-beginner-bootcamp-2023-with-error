package main

import (
	"context"
	"log"
    "github.com/hashicorp/terraform-plugin-sdk/v2/diag" 
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"github.com/google/uuid"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: Provider,
	})

}

type Config struct {
	Endpoint string
	UserUUID string
	Token    string

}

func Provider() *schema.Provider {
	return &schema.Provider{

		ResourcesMap: map[string]*schema.Resource{
			"terratowns_home": Resource(),
		},

		ResourcesMap:   map[string]*schema.Resource{},

		DataSourcesMap: map[string]*schema.Resource{},
		Schema: map[string]*schema.Schema{
			"endpoint": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The endpoint of an external service",
			},
			"token": {
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true, // Marks the Token as sensitive to hide it in the logs
				Description: "The bearer Token for authorization",
			},
			"user_uuid": {
				Type:        schema.TypeString,
				Required:    true,
				ValidateFunc: ValidateUUID,
				Description: "UUID for configuration",
			},
		},
	}
}

func providerConfigure(p *schema.Provider) schema.ConfigureContextFunc {
	return func(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {

		log.Print("providerConfigure:start")
		config := Config{
			Endpoint: d.Get("endpoint").(string),
			Token:    d.Get("token").(string),
			UserUUID: d.Get("user_uuid").(string),
		}
		log.Print("providerConfigure:end")

		log.Print("providerConfigure: start")
		config := Config{
			Endpoint: d.Get("endpoint").(string),
			Token:    d.Get("token").(string),
			UserUuid: d.Get("user_uuid").(string),
		}
		log.Print("providerConfigure: end")

		return &config, nil
	}
}


func Resource() *schema.Resource {
	log.Print("Resource:start")
	return &schema.Resource{
		CreateContext: resourceHouseCreate,
		ReadContext:   resourceHouseRead,
		UpdateContext: resourceHouseUpdate,
		DeleteContext: resourceHouseDelete,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of Home",
			},
			"description": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Home Description",
			},
			"domain_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Home Domain Name",
			},
			"town": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Town to which the home will belong",
			},
			"content_version": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "Content Version",
			},
		},
	}

type Config struct {
	Endpoint string
	Token    string
	UserUuid string

}

func resourceHouseCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diag diag.Diagnostics
	return diag
}

func resourceHouseRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diag diag.Diagnostics
	return diag
}

func resourceHouseUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diag diag.Diagnostics
	return diag
}

func resourceHouseDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diag diag.Diagnostics
	return diag
}

func ValidateUUID(i interface{}, k string) (ws []string, errors []error) {
	log.Print("validateUUID: start")
	value, ok := i.(string)
	if !ok {
		return []string{"Invalid data type for UUID"}, []error{nil}
	}

	_, err := uuid.Parse(value)
	if err != nil {
		return []string{err.Error()}, []error{err}
	}

	return nil, nil
}

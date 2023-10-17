package main

import (
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
    "github.com/google/uuid"
    "log"
)

func main() {
    plugin.Serve(&plugin.ServeOpts{
        ProviderFunc: Provider,
    })
}

func Provider() *schema.Provider {
    return &schema.Provider{
        ResourcesMap: map[string]*schema.Resource{},
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

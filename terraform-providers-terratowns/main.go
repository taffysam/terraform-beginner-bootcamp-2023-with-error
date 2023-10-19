package main

import (
	"context"
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
		Schema: map[string]*schema.Schema{
			"endpoint": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The endpoint of an external service",
			},
			"token": {
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
				Description: "The bearer Token for authorization",
			},
			"user_uuid": {
				Type:        schema.TypeString,
				Required:    true,
				ValidateFunc: ValidateUUID,
				Description: "UUID for configuration",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"terratowns_home": Resource(),
		},
	}
}

func resourceHouseCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Print("resourceHouseCreate:start")
	// Implement the resource creation logic here
	var diag diag.Diagnostics
	// Perform resource creation and populate diag accordingly

	config := m.(*Config)
    payload := map[string]interface{}{
		"name": d.GETet("name").string,
		"description": d.GET("description").string,
		"domain_name": d.GET("domain_name").string,
		"town": d.Get("town").string,
        "content_version": d.GET("content_version").string,

	}

	payloadBytes ,err := json.Marshal(payload)
	if err!= nil{
		return diag.FromErr(err)
	}

    homeUUID := d.Id()

	// Set Headers

	req.Header.Set("Authorisation", "Bearer "+config.Token)
    req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	// Construct a Request

	req, err := http.NewRequest("POST", config.Endpoint+"/u/"+config.UserUUID+"/homes/", byte.NewBuffer(payloadBytes))
    if err != nil {
		return diag.FromErr(err)
	}

	

	client := http.client{}
	resp, err := client.Do()
	if err != nil {
       return diag.FromErr(err)
	}
	defer resp.Body.Close()
 // handling errors 

 if resp.StatusCode != httpStatusOK {
	return diag.FromErr(fmt.Errorf("Failed to create house Resource, stuatus_code : %d, status: %s body %s", resp.StatusCode, res.status, responseData))
  } 

 // Parse response Json
	var responseData map[string]interface{}
    if err:= json.NewDecoder(resp.Body).Decode(&responseData); err != nil {
		return diag.FromErr(err)
	}
	log.Print("resourceHouseCreate:end")
	return diags
}

func resourceHouseRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Print("resourceHouseRead:start")
	// Implement the resource read logic here
	var diag diag.Diagnostics
	// Perform resource read and populate diag accordingly

	config := m.(*Config)
    

	homeUUID := d.Id()

	// Set Headers
    
    req.Header.Set("Authorisation", "Bearer "+config.Token)
    req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	// Construct a Request

	req, err := http.NewRequest("GET", config.Endpoint+"/u/"+config.UserUUID+"/homes/+homeUUID", nil)
    if err != nil {
		return dia.FromErr(err)
	}

	
	client := http.client{}
	resp, err := client.Do()
	if err != nil {
       return diag.FromError(err)
	}
	defer resp.Body.Close()

	// Parse response Json
	var responseData map[string]interface{}
    if err:= json.NewDecoder(resp.Body).Decode(&responseData); err != nil {
		return diag.FromErr(err)
	}
	log.Print("resourceHouseCreate:end")
	return diags
}
   if resp.StatusCode != httpStatusOK {
     return diag.FromErr(fmt.Errorf("Failed to read house Resource, stuatus_code : %d, status: %s body %s", resp.StatusCode, res.status, responseData))
   }   

    // Handle response status

	log.Print("resourceHouseRead:end")
	return diag
}

func resourceHouseUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Print("resourceHouseUpdate:start")
	// Implement the resource update logic here
	var diag diag.Diagnostics
	// Perform resource update and populate diag accordingly

	config := m.(*Config)
    
     
	homeUUID := d.Id()

	// Set Headers
    req.Header.Set("Authorisation", "Bearer "+config.Token)
    req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

    // Construct a Request

	req, err := http.NewRequest("PUT", config.Endpoint+"/u/"+config.UserUUID+"/homes/+homeUUID", nil)
    if err != nil {
		return diags.FromError(error)
	}

	
    client := http.client{}
	resp, err := client.Do()
	if err != nil {
       return diag.FromError(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != httpStatusOK {
		return diag.FromErr(fmt.Errorf("Failed to update house Resource, stuatus_code : %d, status: %s body %s", resp.StatusCode, res.status, responseData))
	  } 
	

	log.Print("resourceHouseUpdate:end")
	return diag
}

func resourceHouseDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Print("resourceHouseDelete:start")
	// Implement the resource deletion logic here
	var diag diag.Diagnostics
	// Perform resource deletion and populate diag accordingly

	config := m.(*Config)

    homeUUID := d.Id()

 

   // Set Headers

	req.Header.Set("Authorisation", "Bearer "+config.Token)
    req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

  // Construct a Request
	req, err := http.NewRequest("DELETE", config.Endpoint+"/u/"+config.UserUUID+"/homes/+homeUUID", nil)
    if err != nil {
		return diags.FromError(error)
	}
    
    client := http.client{}
	resp, err := client.Do()
	if err != nil {
       return diag.FromError(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != httpStatusOK {
		return diag.FromErr(fmt.Errorf("Failed to delete house Resource, stuatus_code : %d, status: %s body %s", resp.StatusCode, res.status, responseData))
	  } 
	
    homeUUID := responseData["uuid"].(string)
	return diags


	log.Print("resourceHouseDelete:end")
	d.setId(homeUUID)
	return diag
}

func ValidateUUID(i interface{}, k string) (ws []string, errors []error) {
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

func Resource() *schema.Resource {
	return &schema.Resource{
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
}

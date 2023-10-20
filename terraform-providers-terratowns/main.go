package main

import (

	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"bytes"
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
		"name":           d.Get("name").(string),
		"description":    d.Get("description").(string),
		"domain_name":    d.Get("domain_name").(string),
		"town":           d.Get("town").(string),
		"content_version": d.Get("content_version").(int),
	}

	payloadBytes ,err := json.Marshal(payload)
	if err!= nil{
		return diag.FromError(err)
	}

	homeUUID := responseData["uuid"].(string)
	d.SetId(homeUUID)

	// Construct a Request
	req, err := http.NewRequest("POST", config.Endpoint+"/u/"+config.UserUUID+"/homes", bytes.NewBuffer(payloadBytes))
    if err != nil {
       return diag.FromErr(err)

		}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
       return diag.FromErr(err)
	}
	defer resp.Body.Close()
 // handling errors 

 // Set Headers

     req.Header.Set("Authorization", "Bearer "+config.Token)
     req.Header.Set("Content-Type", "application/json")
     req.Header.Set("Accept", "application/json")

 var responseData map[string]interface{}

 if resp.StatusCode != http.StatusOK {
	return diag.FromErr(fmt.Errorf("Failed to create house Resource, status_code: %d, status: %s, body: %s", resp.StatusCode, resp.Status, responseData))
}

 // Parse response Json

  if err:= json.NewDecoder(resp.Body).Decode(&responseData); err != nil {
		return diag.FromErr(err)
	}
   log.Print("resourceHouseCreate:end")
	return diag
}

func resourceHouseRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Print("resourceHouseRead:start")
	// Implement the resource read logic here
	var diag diag.Diagnostics
	// Perform resource read and populate diag accordingly

	config := m.(*Config)
    
	// homeUUID := d.Id()
  
	// Construct a Request
req, err := http.NewRequest("POST", config.Endpoint+"/u/"+config.UserUUID+"/homes", bytes.NewBuffer(payloadBytes))
if err != nil {
    if err != nil {
       return diag.FromErr(err)
		}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
       return diag.FromErr(err)
	}
	defer resp.Body.Close()

	// handling errors 

 if resp.StatusCode != http.StatusOK {
	return diag.FromErr(fmt.Errorf("Failed to create house Resource, status_code: %d, status: %s, body: %s", resp.StatusCode, resp.Status, responseData))
}
// Parse response Json
var responseData map[string]interface{}
    if err:= json.NewDecoder(resp.Body).Decode(&responseData); err != nil {

		return diag.FromErr(err)

		}
	
  
	homeUUID := responseData["uuid"].(string)
	d.SetId(homeUUID)
	

	log.Print("resourceHouseCreate:end")
	return diag

		// Set Headers
		req.Header.Set("Authorization", "Bearer "+config.Token)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Accept", "application/json")
	

 payload := map[string]interface{}{
   d.Set("name", responseData["name"].(string)),
   d.Set("description", responseData["description"].(string)),
   d.Set("domain_name", responseData["domain_name"].(string)),
   d.Set("content_version", responseData["content_version"].(int)),
   //d.Set("town", responseData["town"].(string))
 }
   if resp.StatusCode!= http.StatusNotFound {
	d.Set("")
   }

   if resp.StatusCode != http.StatusOK {
     return diag.FromErr(fmt.Errorf("Failed to read house Resource, stuatus_code : %d, status: %s body %s", resp.StatusCode, resp.Status, responseData))
   }   

    // Handle response status

	log.Print("resourceHouseRead:end")
	return diag
}

func resourceHouseRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
    log.Print("resourceHouseUpdate: start")

    // Initialize an empty diagnostics slice
    var diags diag.Diagnostics

    // Assert the configuration
    config := m.(*Config)

    // Create a payload for the update request
    payload := map[string]interface{}{
        "name":           d.Get("name").(string),
        "description":    d.Get("description").(string),
        "content_version": d.Get("content_version").(string),
    }

    payloadBytes, err := json.Marshal(payload)
    if err != nil {
        return diag.FromErr(err)
    }

    // Construct an HTTP request
    req, err := http.NewRequest("POST", config.Endpoint+"/u/"+config.UserUUID+"/homes", bytes.NewBuffer(payloadBytes))
    if err != nil {
        return diag.FromErr(err)
    }

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return diag.FromErr(err)
    }
    defer resp.Body.Close()

    // Set Headers
    req.Header.Set("Authorization", "Bearer "+config.Token)
    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("Accept", "application/json")

    if resp.StatusCode != http.StatusOK {
        return diag.FromErr(fmt.Errorf("Failed to update house resource, status_code: %d, status: %s", resp.StatusCode, resp.Status))
    }

    log.Print("resourceHouseUpdate: end")

    // Set the updated attributes in the Terraform state
    d.Set("name", payload["name"].(string))
    d.Set("description", payload["description"].(string))
    d.Set("content_version", payload["content_version"].(string))

    return diags
}

func resourceHouseDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Print("resourceHouseDelete:start")
	// Implement the resource deletion logic here
	var diag diag.Diagnostics
	// Perform resource deletion and populate diag accordingly

	config := m.(*Config)

	homeUUID := responseData["uuid"].(string)
	d.SetId(homeUUID)
  	   

    

  // Construct a Request
	req, err := http.NewRequest("DELETE", config.Endpoint+"/u/"+config.UserUUID+"/homes/+homeUUID", nil)
    if err != nil {

		return diag.FromErr(err)
	}
    
      // Set Headers

      req.Header.Set("Authorisation", "Bearer "+config.Token)
      req.Header.Set("Content-Type", "application/json")
      req.Header.Set("Accept", "application/json")

    client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
       return diag.FromErr(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return diag.FromErr(fmt.Errorf("Failed to delete house Resource, stuatus_code : %d, status: %s body %s", resp.StatusCode, resp.Status, responseData))
	  } 
	  
	
	d.SetId("")
	log.Print("resourceHouseDelete:end")
	
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

package provider

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform/helper/schema"
)

func validateName(v interface{}, k string) (ws []string, es []error) {
	var errs []error
	var warns []string
	value, ok := v.(string)
	if !ok {
		errs = append(errs, fmt.Errorf("Expected name to be string"))
		return warns, errs
	}
	whiteSpace := regexp.MustCompile(`\s+`)
	if whiteSpace.Match([]byte(value)) {
		errs = append(errs, fmt.Errorf("name cannot contain whitespace. Got %s", value))
		return warns, errs
	}
	return warns, errs
}

func resourceProfile() *schema.Resource {
	fmt.Print()
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				Description:  "The name of the employee",
				ForceNew:     true,
				ValidateFunc: validateName,
			},
			"email": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "employee email id",
			},
		},
		Create: resourceCreateProfile,
		Read:   resourceReadProfile,
		Update: resourceUpdateProfile,
		Delete: resourceDeleteProfile,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
	}
}

func resourceCreateProfile(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	tfTags := d.Get("tags").(*schema.Set).List()
	tags := make([]string, len(tfTags))
	for i, tfTag := range tfTags {
		tags[i] = tfTag.(string)
	}

	item := server.Item{
		Name:        d.Get("name").(string),
		Email: d.Get("description").(string),
	}

	err := apiClient.NewProfile(&profile)

	if err != nil {
		return err
	}
	d.SetId(profile.Name)
	return nil
}

func resourceReadProfile(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	itemId := d.Id()
	item, err := apiClient.GetItem(itemId)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			d.SetId("")
		} else {
			return fmt.Errorf("error finding Item with ID %s", itemId)
		}
	}

	d.SetId(item.Name)
	d.Set("name", item.Name)
	d.Set("description", item.Description)
	d.Set("tags", item.Tags)
	return nil
}

func resourceUpdateProfile(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	tfTags := d.Get("tags").(*schema.Set).List()
	tags := make([]string, len(tfTags))
	for i, tfTag := range tfTags {
		tags[i] = tfTag.(string)
	}

	item := server.Item{
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
		Tags:        tags,
	}

	err := apiClient.UpdateItem(&item)
	if err != nil {
		return err
	}
	return nil
}

func resourceDeleteProfile(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	itemId := d.Id()

	err := apiClient.DeleteItem(itemId)
	if err != nil {
		return err
	}
	d.SetId("")
	return nil
}

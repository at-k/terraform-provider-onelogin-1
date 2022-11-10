package onelogin

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/onelogin/onelogin-go-sdk/pkg/client"
)

// Users returns a resource with the CRUD methods and Terraform Schema defined
func UserCustomAttribute() *schema.Resource {
	return &schema.Resource{
		Create:   userCustomAttributeCreate,
		Read:     userCustomAttributeRead,
		Update:   userCustomAttributeUpdate,
		Delete:   userCustomAttributeDelete,
		Schema:   map[string]*schema.Schema{
			"user_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
      "custom_attributes": &schema.Schema{
        Type:     schema.TypeMap,
        Optional: true,
        Elem:     &schema.Schema{Type: schema.TypeString},
      },
    },
	}
}

func userCustomAttributeCreate(d *schema.ResourceData, m interface{}) error {
  return userCustomAttributeUpdate(d,m)
}

func userCustomAttributeUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*client.APIClient)

  userId := d.Get("user_id")
  attrs := d.Get("custom_attributes")

  user, err := client.Services.UsersV2.GetOne(int32(userId.(int)))
	if err != nil {
		log.Printf("[ERROR] There was a problem reading the user!")
		log.Println(err)
		return err
	}

  user.CustomAttributes = attrs.(map[string]interface{})

	err = client.Services.UsersV2.Update(user)
	if err != nil {
		log.Println("[ERROR] There was a problem updating the user!", err)
		return err
	}
	log.Printf("[UPDATED] Updated user with %d", *(user.ID))

	d.SetId(fmt.Sprintf("%d", *(user.ID)))
	return userCustomAttributeRead(d, m)
}

func userCustomAttributeRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*client.APIClient)
	uid, _ := strconv.Atoi(d.Id())
	user, err := client.Services.UsersV2.GetOne(int32(uid))
	if err != nil {
		log.Printf("[ERROR] There was a problem reading the user!")
		log.Println(err)
		return err
	}
	if user == nil {
		d.SetId("")
		return nil
	}
	log.Printf("[READ] Reading user with %d", *(user.ID))

	d.Set("custom_attributes", user.CustomAttributes)

	return nil
}

func userCustomAttributeDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}

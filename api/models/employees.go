package employees

import (
	"fmt"
)

func get(id string) map[string]interface{} {
  return map[string]string{
    "Id": "9B405619-4656-458B-ABC1-B9A4D5A65105",
  	"Name": "Sam",
  	"Email": "sam@myob.com",
  	"Phone": "",
  	"Department": map[string]string{
      "Id": "8A34C7DE-5F10-4F24-BCDC-E1ACA97003EE",
      "Name": "sme",
      "Description": "sme"
    }
  }
}

get("asdf")

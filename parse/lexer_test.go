package parse

import (
	"fmt"
	"testing"
)

func TestValidIni(t *testing.T) {
	ini := `
[SectionName]
key1=value 1
key2=value 2`
	expected := `
{
  "Sections": [
    {
      "Name": "SectionName",
      "KeyValuePairs": [
        {
          "Key": "key1",
          "Value": "value 1"
        },
        {
          "Key": "key2",
          "Value": "value 2"
        }
      ]
    }
  ]
}`
	fmt.Println(ini, expected)
}

package parse

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestParse(t *testing.T) {
	in := `
[SectionName]
key1=value 1
key2=value 2
`
	expected := `{
  "FileName": "%s",
  "Sections": [
    {
      "Name": "SectionName",
      "KeyValues": [
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
	name := "TestEqual"
	ini := Parse(name, in)
	b, err := json.MarshalIndent(ini, "", "  ")
	if err != nil {
		t.Error(err)
	}
	actual := string(b)
	expected = fmt.Sprintf(expected, name)
	if actual != expected {
		t.Errorf("%s: \ngot\n\t%+v\nexpected\n\t%v", name, actual, expected)
		return
	}
	t.Log(name, "OK")
}

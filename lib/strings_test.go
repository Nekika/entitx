package lib

import (
	"reflect"
	"testing"
)

func TestCapitalize(t *testing.T) {
	cases := []struct {
		Name     string
		Input    string
		Expected string
	}{
		{
			Name:     "Empty string",
			Input:    "",
			Expected: "",
		},
		{
			Name:     "Character",
			Input:    "a",
			Expected: "A",
		},
		{
			Name:     "Word",
			Input:    "hello",
			Expected: "Hello",
		},
		{
			Name:     "Sentence",
			Input:    "hello everyone.",
			Expected: "Hello everyone.",
		},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			output := Capitalize(c.Input)
			if !reflect.DeepEqual(c.Expected, output) {
				t.Errorf("values are not equal: expected %s but got %s", c.Expected, output)
			}
		})
	}
}

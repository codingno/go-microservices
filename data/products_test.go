package data

import "testing"

func TestChecksValidation(t *testing.T) {
	p := &Product{
		Name:  "frappuchino",
		Price: 2.55,
		SKU:   "abc-def-uuu",
	}

	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}

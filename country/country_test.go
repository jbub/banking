package country

import "testing"

func TestCountryExists(t *testing.T) {
	code := "GB"
	ok := Exists(code)
	if !ok {
		t.Errorf("expected country %v to exist", code)
	}
}

func TestCountryNotPresent(t *testing.T) {
	code := "XX"
	ok := Exists(code)
	if ok {
		t.Errorf("country %v should not exist", code)
	}
}

func TestValidCountry(t *testing.T) {
	alpha2 := "GB"
	alpha3 := "GBR"
	name := "United Kingdom"
	c, ok := Get(alpha2)
	if !ok {
		t.Errorf("expected country %v to exist", alpha2)
	}
	if alpha2 != c.Alpha2Code {
		t.Errorf("expected %v got %v", alpha2, c.Alpha2Code)
	}
	if alpha3 != c.Alpha3Code {
		t.Errorf("expected %v got %v", alpha3, c.Alpha3Code)
	}
	if name != c.Name {
		t.Errorf("expected %v got %v", name, c.Name)
	}
	if c.Name != c.String() {
		t.Errorf("expected %v got %v", c.Name, c.String())
	}
}

func TestInvalidCountry(t *testing.T) {
	code := "XX"
	alpha2 := ""
	alpha3 := ""
	name := ""
	c, ok := Get(code)
	if ok {
		t.Errorf("country %v should not exist", code)
	}
	if alpha2 != c.Alpha2Code {
		t.Errorf("expected %v got %v", alpha2, c.Alpha2Code)
	}
	if alpha3 != c.Alpha3Code {
		t.Errorf("expected %v got %v", alpha3, c.Alpha3Code)
	}
	if name != c.Name {
		t.Errorf("expected %v got %v", name, c.Name)
	}
	if c.Name != c.String() {
		t.Errorf("expected %v got %v", c.Name, c.String())
	}
}

func TestValidBbanStructure(t *testing.T) {
	code := "GB"
	structure, ok := GetBbanStructure(code)
	if !ok {
		t.Errorf("bban structure for %v should exist", code)
	}
	want := 18
	if want != structure.Length() {
		t.Errorf("expected %v got %v", want, structure.Length())
	}
}

func TestInvalidBbanStructure(t *testing.T) {
	code := "XX"
	structure, ok := GetBbanStructure(code)
	if ok {
		t.Errorf("bban structure for %v should not exist", code)
	}
	want := 0
	if want != structure.Length() {
		t.Errorf("expected %v got %v", want, structure.Length())
	}
}

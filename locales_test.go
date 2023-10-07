package locales_test

import (
	"testing"

	locales "github.com/iMed-Cloud-Services-Inc/imedcs-locale"
)

func TestIsValid(t *testing.T) {
	us := locales.LocaleUS
	if !us.IsValid() {
		t.Errorf("%s - is actually valid ISO code", us)
	}
	inv := locales.Locale("XYS")
	if inv.IsValid() {
		t.Errorf("%s - is actually invalid ISO code", inv)
	}
}

func TestCountry(t *testing.T) {
	ae := locales.LocaleAE
	if ae.Country() != "United Arab Emirates" {
		t.Errorf("%s - represents 'United Arab Emirates'", ae)
	}
	gl := locales.LocaleWildcard
	if gl.Country() != "Global" {
		t.Errorf("%s - represents any country", ae)
	}
	xyz := locales.Locale("XYZ")
	if xyz.Country() != "" {
		t.Errorf("%s - does not represent any country", xyz)
	}
}

func TestMake(t *testing.T) {
	ca, err := locales.Make("   Ca      ")
	if err != nil {
		t.Error("Ca - should be trimmed and converted to LocaleCA")
	} else if ca != locales.LocaleCA {
		t.Error("Ca - should be trimmed and converted to LocaleCA")
	}
	_, err = locales.Make("  Invalid ")
	if err == nil {
		t.Error("Invalid ISO code should return an error")
	}
}

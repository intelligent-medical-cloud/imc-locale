package locale_test

import (
	"testing"

	locale "github.com/intelligent-medical-cloud/imc-locale"
)

func TestIsValid(t *testing.T) {
	us := locale.LocaleUS
	if !us.IsValid() {
		t.Errorf("%s - is actually valid ISO code", us)
	}

	inv := locale.Locale("XYS")
	if inv.IsValid() {
		t.Errorf("%s - is actually invalid ISO code", inv)
	}
}

func TestCountry(t *testing.T) {
	ae := locale.LocaleAE
	if ae.Country() != "United Arab Emirates" {
		t.Errorf("%s - represents 'United Arab Emirates'", ae)
	}

	gl := locale.LocaleWildcard
	if gl.Country() != "Global" {
		t.Errorf("%s - represents any country", ae)
	}

	xyz := locale.Locale("XYZ")
	if xyz.Country() != "" {
		t.Errorf("%s - does not represent any country", xyz)
	}
}

func TestMake(t *testing.T) {
	ca, err := locale.Make("   Ca      ")
	if err != nil {
		t.Error("Ca - should be trimmed and converted to LocaleCA")
	} else if ca != locale.LocaleCA {
		t.Error("Ca - should be trimmed and converted to LocaleCA")
	}

	_, err = locale.Make("  Invalid ")
	if err == nil {
		t.Error("Invalid ISO code should return an error")
	}

	us, err := locale.Make("en-US,en;q=0.5")
	if err != nil || us.ISOLowercase() != "us" {
		t.Error("`en-US,en;q=0.5` should be US")
	}

	ch, err := locale.Make("fr-CH, fr;q=0.9, en;q=0.8, de;q=0.7, *;q=0.5")
	if err != nil || ch.ISOLowercase() != "ch" {
		t.Error("`fr-CH, fr;q=0.9, en;q=0.8, de;q=0.7, *;q=0.5` should be CH")
	}
}

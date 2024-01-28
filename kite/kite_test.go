package kite

import (
	"testing"
)

func TestKiteWithAllProperties(t *testing.T) {
	expected := Kite{
		Brand:  "HighFly",
		Color:  "Yellow",
		Length: 50,
	}

	kite := NewKite(
		WithBrand("HighFly"),
		WithColor("Yellow"),
		WithLength(50),
	)

	assertKiteProperties(t, expected, *kite, "Kite with all properties")
}

func TestKiteWithSingleProperty(t *testing.T) {
	expected := Kite{
		Brand:  "DefaultBrand",
		Color:  "Green",
		Length: 999,
	}

	kite := NewKite(
		WithColor("Green"),
	)

	assertKiteProperties(t, expected, *kite, "Kite with single property")
}

func TestKiteWithDefaults(t *testing.T) {
	expected := Kite{
		Brand:  "DefaultBrand",
		Color:  "DefaultColor",
		Length: 999,
	}

	kite := NewKite()
	assertKiteProperties(t, expected, *kite, "Kite with defaults")
}

func assertKiteProperties(t *testing.T, expected, actual Kite, desc string) {
	t.Run(desc, func(t *testing.T) {
		if expected != actual {
			t.Errorf("Expected %v, Got %v", expected, actual)
		}
	})
}

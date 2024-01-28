package kite

import (
	"testing"
)

func TestKiteWithAllCustomProperties(t *testing.T) {
	expected := Kite{
		Brand:  "HighFly",
		Color:  "Yellow",
		Weight: 50,
	}

	kite := NewKite(
		WithBrand("HighFly"),
		WithColor("Yellow"),
		WithWeight(50),
	)

	assertKiteProperties(t, expected, *kite, "Kite with all properties")
}

func TestKiteWithSingleCustomProperty(t *testing.T) {
	expected := Kite{
		Brand:  DefaultBrand,
		Color:  "Green",
		Weight: DefaultWeight,
	}

	kite := NewKite(
		WithColor("Green"),
	)

	assertKiteProperties(t, expected, *kite, "Kite with single property")
}

func TestKiteWithNoCustomProperties(t *testing.T) {
	expected := Kite{
		Brand:  DefaultBrand,
		Color:  DefaultColor,
		Weight: DefaultWeight,
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

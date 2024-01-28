package sailboat

import (
	"testing"
)

func TestSailboatWithSingleProperty(t *testing.T) {
	expected := Sailboat{
		Brand:  "Hobie Cat",
		Color:  DefaultColor,
		Weight: DefaultWeight,
	}

	boat := NewSailboatBuilder().
		Brand("Hobie Cat").
		Build()

	assertSailboatProperties(t, expected, *boat, "Sailboat with single property")
}

func TestSailboatWithNoProperties(t *testing.T) {
	expected := Sailboat{
		Brand:  DefaultBrand,
		Color:  DefaultColor,
		Weight: DefaultWeight,
	}

	boat := NewSailboatBuilder().
		Build()

	assertSailboatProperties(t, expected, *boat, "Sailboat with no properties")
}

func TestSailboatWithAllProperties(t *testing.T) {
	expected := Sailboat{
		Brand:  "Catalina",
		Color:  "Blue",
		Weight: 16,
	}

	boat := NewSailboatBuilder().
		Color("Blue").
		Brand("Catalina").
		Length(16).
		Build()

	assertSailboatProperties(t, expected, *boat, "Sailboat with all properties")
}

func assertSailboatProperties(t *testing.T, expected, actual Sailboat, desc string) {
	t.Run(desc, func(t *testing.T) {
		if expected != actual {
			t.Errorf("Expected %v, Got %v", expected, actual)
		}
	})
}

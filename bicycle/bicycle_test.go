package bicycle

import (
	"testing"
)

func TestBicycleWithAllProperties(t *testing.T) {
	expected := Bicycle{
		Brand:  "Giant",
		Color:  "Blue",
		Weight: 10,
	}

	bike := NewBicycle(
		WithBrand("Giant"),
		WithColor("Blue"),
		WithWeight(10),
	)

	assertBicycleProperties(t, expected, *bike, "Bicycle with all properties explicitly specified")
}

func TestBicycleWithSingleProperty(t *testing.T) {
	expected := Bicycle{
		Brand:  DefaultBrand,
		Color:  "Green",
		Weight: DefaultWeight,
	}

	bike := NewBicycle(
		WithColor("Green"),
	)

	assertBicycleProperties(t, expected, *bike, "Bicycle with single property specified")
}

func TestBicycleWithCompositeLiteralAllProperties(t *testing.T) {
	expected := Bicycle{
		Brand:  "Specialized",
		Color:  "Red",
		Weight: 12,
	}

	bike := &Bicycle{
		Brand:  "Specialized",
		Color:  "Red",
		Weight: 12,
	}

	assertBicycleProperties(t, expected, *bike, "Bicycle with composite literal, all properties explicitly specified")
}

func TestBicycleWithCompositeLiteralSingleProperty(t *testing.T) {
	expected := Bicycle{
		Brand:  "SpeedX",
		Color:  "",
		Weight: 0,
	}

	bike := &Bicycle{
		Brand: "SpeedX",
	}

	assertBicycleProperties(t, expected, *bike, "Bicycle with composite literal, single property specified")
}

func TestBicycleWithNoPropertiesUsingDefaults(t *testing.T) {
	expected := Bicycle{
		Brand:  DefaultBrand,
		Color:  DefaultColor,
		Weight: DefaultWeight,
	}

	bike := NewBicycle()

	assertBicycleProperties(t, expected, *bike, "Bicycle with no properties specified")
}

func assertBicycleProperties(t *testing.T, expected, actual Bicycle, desc string) {
	t.Run(desc, func(t *testing.T) {
		if expected != actual {
			t.Errorf("Expected %v, Got %v", expected, actual)
		}
	})
}

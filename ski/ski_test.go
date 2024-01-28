package ski

import (
	"reflect"
	"testing"
)

func TestSkiWithAllPropertiesOverridden(t *testing.T) {
	expected := &Ski{
		Brand:  "Atomic",
		Color:  "Red",
		Length: 180,
	}

	ski := NewSkiPrototype()
	ski.Brand = "Atomic"
	ski.Color = "Red"
	ski.Length = 180

	assertSkiProperties(t, expected, ski, "Ski with all properties overridden")
}

func TestSkiWithSinglePropertyOverridden(t *testing.T) {
	expected := &Ski{
		Brand:  "DefaultBrand",
		Color:  "Blue",
		Length: 170,
	}

	ski := NewSkiPrototype()
	ski.Color = "Blue"

	assertSkiProperties(t, expected, ski, "Ski with single property overridden")
}

func TestSkiWithNoPropertiesOverridden(t *testing.T) {
	expected := &Ski{
		Brand:  "DefaultBrand",
		Color:  "DefaultColor",
		Length: 170,
	}

	ski := NewSkiPrototype()

	assertSkiProperties(t, expected, ski, "Ski with no properties overridden")
}

func assertSkiProperties(t *testing.T, expected, actual *Ski, desc string) {
	t.Run(desc, func(t *testing.T) {
		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("Expected %+v, Got %+v", expected, actual)
		}
	})
}
